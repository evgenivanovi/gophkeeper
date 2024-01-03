package secret

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/usecase/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
	teax "github.com/evgenivanovi/gophkeeper/pkg/tea"
	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

type CreateCredentialsViewSettings struct {
	NoStyle lipgloss.Style

	FocusedStyle  lipgloss.Style
	FocusedButton string

	BlurStyle  lipgloss.Style
	BlurButton string
}

func ProvideCreateCredentialsViewSettings() CreateCredentialsViewSettings {

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

	return CreateCredentialsViewSettings{
		NoStyle: NoStyle,

		FocusedStyle: FocusedStyle,
		BlurStyle:    BlurStyle,

		FocusedButton: FocusedButton,
		BlurButton:    BlurButton,
	}

}

type CreateCredentialsModel struct {
	focus  int
	inputs []textinput.Model
	view   CreateCredentialsViewSettings

	done bool

	op      common.Options
	args    CreateSecretArg
	usecase secret.CreateDecodedSecretUsecase
	result  *CreateCredentialsResultMsg
}

func ProvideCreateCredentialsModel(
	op common.Options,
	args CreateSecretArg,
	usecase secret.CreateDecodedSecretUsecase,
) *CreateCredentialsModel {

	md := &CreateCredentialsModel{
		inputs: make([]textinput.Model, 2),
		view:   ProvideCreateCredentialsViewSettings(),

		op:      op,
		args:    args,
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

func (m *CreateCredentialsModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *CreateCredentialsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case CreateCredentialsResultMsg:
		m.done = true
		m.result = &msg
		return teax.Quit(m)
	case tea.KeyMsg:
		switch msg.Type {

		default:

		// End
		case tea.KeyCtrlC, tea.KeyEsc:
			return teax.Quit(m)

		// Set focus to next input
		case tea.KeyEnter, tea.KeyTab, tea.KeyShiftTab, tea.KeyUp, tea.KeyDown:
			key := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if key == tea.KeyEnter.String() && m.focus == len(m.inputs) {
				return m, m.Execute()
			}

			// Cycle indexes
			if key == tea.KeyUp.String() || key == tea.KeyShiftTab.String() {
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

func (m *CreateCredentialsModel) updateInputs(msg tea.Msg) tea.Cmd {
	commands := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for ind := range m.inputs {
		m.inputs[ind], commands[ind] = m.inputs[ind].Update(msg)
	}

	return tea.Batch(commands...)
}

func (m *CreateCredentialsModel) View() string {
	if m.done {
		return m.finish()
	}
	return m.resume()
}

func (m *CreateCredentialsModel) Execute() tea.Cmd {
	ctx := context.Background()
	ctx = common.OptionsWithCtx(ctx, m.op)

	req := m.buildRequest(m.args)
	return ExecuteCreateCredentials(m.usecase, ctx, req)
}

func (m *CreateCredentialsModel) buildRequest(
	args CreateSecretArg,
) secretsharedmd.DecodedSecretDataModel {

	username := m.inputs[0].Value()
	password := m.inputs[1].Value()

	return secretsharedmd.DecodedSecretDataModel{
		Name: args.Name,
		Type: "CREDENTIALS",
		Content: &secretsharedmd.CredentialsSecretContentModel{
			Username: username,
			Password: password,
		},
	}

}

func (m *CreateCredentialsModel) resume() string {

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

func (m *CreateCredentialsModel) finish() string {

	if m.result.IsError() {
		output := strings.Builder{}
		output.WriteString(m.result.err.Error())
		output.WriteString(std.NL)
		return output.String()
	}

	if !m.result.IsError() {
		output := strings.Builder{}
		output.WriteString("Credentials created!")
		output.WriteString(std.NL)
		return output.String()
	}

	panic("unexpected")

}

/* __________________________________________________ */

type CreateCredentialsResultMsg struct {
	err error
}

func (msg *CreateCredentialsResultMsg) IsError() bool {
	return msg.err != nil
}

func ExecuteCreateCredentials(
	uc secret.CreateDecodedSecretUsecase, ctx context.Context, data secretsharedmd.DecodedSecretDataModel,
) tea.Cmd {

	return func() tea.Msg {

		err := uc.Execute(ctx, data)
		if err != nil {
			return CreateCredentialsResultMsg{
				err: err,
			}
		}

		return CreateCredentialsResultMsg{}

	}

}

/* __________________________________________________ */
