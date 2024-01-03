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

type SigninViewSettings struct {
	NoStyle lipgloss.Style

	FocusedStyle  lipgloss.Style
	FocusedButton string

	BlurStyle  lipgloss.Style
	BlurButton string
}

func ProvideSigninViewSettings() SigninViewSettings {

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

	return SigninViewSettings{
		NoStyle: NoStyle,

		FocusedStyle: FocusedStyle,
		BlurStyle:    BlurStyle,

		FocusedButton: FocusedButton,
		BlurButton:    BlurButton,
	}

}

type SigninModel struct {
	focus    int
	inputs   []textinput.Model
	settings SigninViewSettings

	done bool

	op      common.Options
	usecase auth.SigninUsecase
	result  *SigninResultMsg
}

func ProvideSigninModel(
	op common.Options,
	usecase auth.SigninUsecase,
) *SigninModel {

	md := &SigninModel{
		inputs:   make([]textinput.Model, 2),
		settings: ProvideSigninViewSettings(),

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

func (m *SigninModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *SigninModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case SigninResultMsg:
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
					m.inputs[ind].PromptStyle = m.settings.FocusedStyle
					m.inputs[ind].TextStyle = m.settings.FocusedStyle
					continue
				}

				if ind != m.focus {
					// Remove focused state
					m.inputs[ind].Blur()
					m.inputs[ind].PromptStyle = m.settings.NoStyle
					m.inputs[ind].TextStyle = m.settings.NoStyle
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

func (m *SigninModel) updateInputs(msg tea.Msg) tea.Cmd {
	commands := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for ind := range m.inputs {
		m.inputs[ind], commands[ind] = m.inputs[ind].Update(msg)
	}

	return tea.Batch(commands...)
}

func (m *SigninModel) Execute() tea.Cmd {
	ctx := context.Background()
	ctx = common.OptionsWithCtx(context.Background(), m.op)

	req := m.buildRequest()
	return ExecuteSignin(m.usecase, ctx, req)
}

func (m *SigninModel) buildRequest() auth.SignInRequest {

	username := m.inputs[0].Value()
	password := m.inputs[1].Value()

	return auth.SignInRequest{
		Payload: auth.SignInRequestPayload{
			Credentials: authshared.CredentialsModel{
				Username: username,
				Password: password,
			},
		},
	}

}

func (m *SigninModel) View() string {
	if m.done {
		return m.finish()
	}
	return m.resume()
}

func (m *SigninModel) resume() string {

	var output strings.Builder

	for in := range m.inputs {
		output.WriteString(m.inputs[in].View())
		if in <= len(m.inputs) {
			output.WriteString(std.NL)
		}
	}

	button := &m.settings.BlurButton
	if m.focus == len(m.inputs) {
		button = &m.settings.FocusedButton
	}

	output.WriteString(std.NL)
	output.WriteString(*button)
	output.WriteString(std.NL)
	return output.String()

}

func (m *SigninModel) finish() string {

	if m.result.IsError() {
		output := strings.Builder{}
		output.WriteString(m.result.err.Error())
		output.WriteString(std.NL)
		return output.String()
	}

	if !m.result.IsError() {
		output := strings.Builder{}
		output.WriteString("Signin succeed!")
		output.WriteString(std.NL)
		return output.String()
	}

	panic("unexpected")

}

/* __________________________________________________ */

type SigninResultMsg struct {
	result auth.SignInResponse
	err    error
}

func (msg *SigninResultMsg) IsError() bool {
	return msg.err != nil
}

func ExecuteSignin(
	uc auth.SigninUsecase, ctx context.Context, request auth.SignInRequest,
) tea.Cmd {

	return func() tea.Msg {

		response, err := uc.Execute(ctx, request)
		if err != nil {
			return SigninResultMsg{
				err: err,
			}
		}

		return SigninResultMsg{
			result: response,
		}

	}

}

/* __________________________________________________ */
