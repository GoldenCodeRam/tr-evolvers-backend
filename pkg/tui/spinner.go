package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Spinner struct {
	Model     spinner.Model
	IsLoading bool

	channel chan error
}

func NewSpinner() *Spinner {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return &Spinner{
		s, false, make(chan error),
	}
}

func (s *Spinner) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	s.Model, cmd = s.Model.Update(msg)
	return cmd
}

func (s *Spinner) DoWithSpinner(
	call func() error,
	then func(chan error) error,
	message string,
) tea.Cmd {
	s.IsLoading = true

	return func() tea.Msg {
		go then(s.channel)

		error := call()
		if error != nil {
			s.channel <- nil
		}
		s.channel <- error

		return nil
	}
}

func (s *Spinner) GenerateMessage(message string) string {
	return lipgloss.JoinHorizontal(lipgloss.Left, append(
		[]string{
			s.Model.View(),
			message,
		},
	)...)
}
