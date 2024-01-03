package tea

import tea "github.com/charmbracelet/bubbletea"

/* __________________________________________________ */

func Quit(m tea.Model) (tea.Model, func() tea.Msg) {
	return m, tea.Quit
}

/* __________________________________________________ */
