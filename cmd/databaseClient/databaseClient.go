package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/goldencoderam/tr-evolvers-backend/pkg/database"
	"github.com/goldencoderam/tr-evolvers-backend/pkg/tui"
	"github.com/goldencoderam/tr-evolvers-backend/pkg/tui/menu"
	"github.com/jaswdr/faker"
)

var (
	titleStyle     = lipgloss.NewStyle().PaddingLeft(4).Foreground(lipgloss.Color("86"))
	quitTextStyle  = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	paragraphStyle = lipgloss.NewStyle().Padding(1, 0, 1, 4).Width(80)
)

type model struct {
	menu    *menu.Menu
	spinner *tui.Spinner

	quitting bool
}

func (m model) Init() tea.Cmd {
	return m.spinner.Model.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.menu.Model.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.menu.Model.SelectedItem().(menu.MenuItem)
			if ok {
				switch i.Code {
				case menu.MAIN:
					cmds = append(cmds, m.menu.ChangeMenu(menu.MAIN_MENU))
				case menu.MIGRATIONS:
					cmds = append(cmds, m.menu.ChangeMenu(menu.MIGRATIONS_MENU))
				case menu.MAKE_MIGRATIONS:
					cmd = m.spinner.DoWithSpinner(func() error {
						return database.AutoMigrate()
					}, func(err chan error) error {
						e := <-err
						m.spinner.IsLoading = false
						return e
					}, "Making migrations...")

					cmds = append(cmds, cmd)
				case menu.SEED:
					fake := faker.New()

					for range make([]int, 10) {
						error := database.CreateEnergyMeterBrand(fake.Company().Name())
						if error != nil {
							panic("Error seeding EnergyMeterBrand table.")
						}
					}
				}
			}
		}
	}

	cmds = append(cmds, m.menu.Update(msg))
	cmds = append(cmds, m.spinner.Update(msg))

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("Nothing")
	}

	if m.spinner.IsLoading {
		return m.spinner.GenerateMessage("Making migrations...")
	}

	return lipgloss.JoinVertical(lipgloss.Left, append(
		[]string{
			paragraphStyle.Render("Here are some commands for the database administration."),
		},
		m.menu.Model.View(),
	)...)
}

func main() {
	s := spinner.New()
	s.Spinner = spinner.Dot

	m := model{
		menu:    menu.NewMenu(),
		spinner: tui.NewSpinner(),
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		panic(err)
	}
}
