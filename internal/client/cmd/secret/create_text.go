package secret

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/usecase/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
	teax "github.com/evgenivanovi/gophkeeper/pkg/tea"
	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

type CreateTextModel struct {
	focus int
	area  textarea.Model

	done bool

	op      common.Options
	args    CreateSecretArg
	usecase secret.CreateDecodedSecretUsecase
	result  *CreateTextResultMsg
}

func ProvideCreateTextModel(
	op common.Options,
	args CreateSecretArg,
	usecase secret.CreateDecodedSecretUsecase,
) *CreateTextModel {

	area := textarea.New()
	area.Focus()
	area.ShowLineNumbers = false

	md := &CreateTextModel{
		area: area,

		op:      op,
		args:    args,
		usecase: usecase,
	}

	return md

}

func (m *CreateTextModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *CreateTextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case CreateTextResultMsg:
		m.done = true
		m.result = &msg
		return teax.Quit(m)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlX:
			return m, m.Execute()
		case tea.KeyEsc:
			if m.area.Focused() {
				m.area.Blur()
			}
		case tea.KeyCtrlC:
			return teax.Quit(m)
		default:
			if !m.area.Focused() {
				cmd = m.area.Focus()
				cmds = append(cmds, cmd)
			}
		}
	}

	m.area, cmd = m.area.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)

}

func (m *CreateTextModel) View() string {
	if m.done {
		return m.finish()
	}
	return m.resume()
}

func (m *CreateTextModel) Execute() tea.Cmd {
	ctx := context.Background()
	ctx = common.OptionsWithCtx(ctx, m.op)

	req := m.buildRequest(m.args)
	return ExecuteCreateText(m.usecase, ctx, req)
}

func (m *CreateTextModel) buildRequest(
	args CreateSecretArg,
) secretsharedmd.DecodedSecretDataModel {

	text := m.area.View()

	return secretsharedmd.DecodedSecretDataModel{
		Name: args.Name,
		Type: "TEXT",
		Content: &secretsharedmd.TextSecretContentModel{
			Text: text,
		},
	}

}

func (m *CreateTextModel) resume() string {
	var output strings.Builder
	output.WriteString(std.NL)
	output.WriteString(m.area.View())
	output.WriteString(std.NL)
	return output.String()
}

func (m *CreateTextModel) finish() string {

	if m.result.IsError() {
		output := strings.Builder{}
		output.WriteString(m.result.err.Error())
		output.WriteString(std.NL)
		return output.String()
	}

	if !m.result.IsError() {
		output := strings.Builder{}
		output.WriteString("Text created!")
		output.WriteString(std.NL)
		return output.String()
	}

	panic("unexpected")

}

/* __________________________________________________ */

type CreateTextResultMsg struct {
	err error
}

func (msg *CreateTextResultMsg) IsError() bool {
	return msg.err != nil
}

func ExecuteCreateText(
	uc secret.CreateDecodedSecretUsecase, ctx context.Context, data secretsharedmd.DecodedSecretDataModel,
) tea.Cmd {

	return func() tea.Msg {

		err := uc.Execute(ctx, data)
		if err != nil {
			return CreateTextResultMsg{
				err: err,
			}
		}

		return CreateTextResultMsg{}

	}

}

/* __________________________________________________ */
