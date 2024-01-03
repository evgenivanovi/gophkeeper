package auth

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/usecase/auth"
	authshared "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	teax "github.com/evgenivanovi/gophkeeper/pkg/tea"
	"github.com/evgenivanovi/gpl/std"
)

type SignupViewSettings struct {
	NoStyle lipgloss.Style

	FocusedStyle  lipgloss.Style
	FocusedButton string

	BlurStyle  lipgloss.Style
	BlurButton string
}

func ProvideSignupViewSettings() SignupViewSettings {

	NoStyle := lipgloss.
		NewStyle()

	FocusedStyle := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("205"))

	BlurStyle := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("240"))

	FocusedButton := FocusedStyle.
		Copy().
		Render("OK")

	BlurButton := BlurStyle.
		Copy().
		Render("OK")

	return SignupViewSettings{
		NoStyle: NoStyle,

		FocusedStyle: FocusedStyle,
		BlurStyle:    BlurStyle,

		FocusedButton: FocusedButton,
		BlurButton:    BlurButton,
	}

}

type SignupModel struct {
	focus  int
	inputs []textinput.Model
	view   SignupViewSettings

	done bool

	op      common.Options
	usecase auth.SignupUsecase
	result  *SignupResultMsg
}

func ProvideSignupModel(
	op common.Options,
	usecase auth.SignupUsecase,
) *SignupModel {

	md := &SignupModel{
		inputs: make([]textinput.Model, 2),
		view:   ProvideSignupViewSettings(),

		op:      op,
		usecase: usecase,
	}

	for ind := range md.inputs {

		in := textinput.New()

		switch ind {
		case 0:
			in.Focus()
			in.Placeholder = "Username"
			in.EchoMode = textinput.EchoNormal
		case 1:
			in.Placeholder = "Password"
			in.EchoMode = textinput.EchoPassword
			in.EchoCharacter = '•'
		}

		md.inputs[ind] = in

	}

	return md

}

func (m *SignupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *SignupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case SignupResultMsg:
		m.done = true
		m.result = &msg
		return teax.Quit(m)
	case tea.KeyMsg:
		switch msg.String() {

		// End
		case "ctrl+c", "esc":
			return teax.Quit(m)

		// Set focus to next input
		case "enter", "tab", "shift+tab", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focus == len(m.inputs) {
				return m, m.Execute()
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focus--
			} else {
				m.focus++
			}

			if m.focus > len(m.inputs) {
				m.focus = 0
			} else if m.focus < 0 {
				m.focus = len(m.inputs)
			}

			commands := make([]tea.Cmd, len(m.inputs))
			for ind := 0; ind < len(m.inputs); ind++ {

				if ind == m.focus {
					// Set focused state
					commands[ind] = m.inputs[ind].Focus()
					m.inputs[ind].PromptStyle = m.view.FocusedStyle
					m.inputs[ind].TextStyle = m.view.FocusedStyle
					continue
				}

				if ind != m.focus {
					// Remove focused state
					m.inputs[ind].Blur()
					m.inputs[ind].PromptStyle = m.view.NoStyle
					m.inputs[ind].TextStyle = m.view.NoStyle
					continue
				}

			}

			return m, tea.Batch(commands...)

		}

	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	return m, cmd

}

func (m *SignupModel) updateInputs(msg tea.Msg) tea.Cmd {
	commands := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for ind := range m.inputs {
		m.inputs[ind], commands[ind] = m.inputs[ind].Update(msg)
	}

	return tea.Batch(commands...)
}

func (m *SignupModel) Execute() tea.Cmd {
	ctx := context.Background()
	ctx = common.OptionsWithCtx(ctx, m.op)

	req := m.buildRequest()
	return ExecuteSignup(m.usecase, ctx, req)
}

func (m *SignupModel) buildRequest() auth.SignUpRequest {

	username := m.inputs[0].Value()
	password := m.inputs[1].Value()

	return auth.SignUpRequest{
		Payload: auth.SignUpRequestPayload{
			Credentials: authshared.CredentialsModel{
				Username: username,
				Password: password,
			},
		},
	}

}

func (m *SignupModel) View() string {
	if m.done {
		return m.finish()
	}
	return m.resume()
}

func (m *SignupModel) resume() string {

	var output strings.Builder

	for in := range m.inputs {
		output.WriteString(m.inputs[in].View())
		if in <= len(m.inputs) {
			output.WriteString(std.NL)
		}
	}

	button := &m.view.BlurButton
	if m.focus == len(m.inputs) {
		button = &m.view.FocusedButton
	}

	output.WriteString(std.NL)
	output.WriteString(*button)
	output.WriteString(std.NL)
	return output.String()

}

func (m *SignupModel) finish() string {

	if m.result.IsError() {
		output := strings.Builder{}
		output.WriteString(m.result.err.Error())
		output.WriteString(std.NL)
		return output.String()
	}

	if !m.result.IsError() {
		output := strings.Builder{}
		output.WriteString("Signup succeed!")
		output.WriteString(std.NL)
		return output.String()
	}

	panic("unexpected")

}

/* __________________________________________________ */

type SignupResultMsg struct {
	result auth.SignUpResponse
	err    error
}

func (msg *SignupResultMsg) IsError() bool {
	return msg.err != nil
}

func ExecuteSignup(
	uc auth.SignupUsecase, ctx context.Context, request auth.SignUpRequest,
) tea.Cmd {

	return func() tea.Msg {

		response, err := uc.Execute(ctx, request)
		if err != nil {
			return SignupResultMsg{
				err: err,
			}
		}

		return SignupResultMsg{
			result: response,
		}

	}

}

/* __________________________________________________ */
