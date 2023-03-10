package menu

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	listTitleStyle    = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type MenuCode int

const (
	MAIN MenuCode = iota
	MIGRATIONS
	MAKE_MIGRATIONS
    SEED
)

var (
    MAIN_MENU = []list.Item{
        MenuItem{
            MIGRATIONS,
            "Migrations menu",
        },
    }
    MIGRATIONS_MENU = []list.Item{
        MenuItem{
            MAKE_MIGRATIONS,
            "Make migrations",
        },
        MenuItem{
            SEED,
            "Seed database",
        },
        MenuItem{
            MAIN,
            "Main menu",
        },
    }
)

type MenuItem struct {
	Code  MenuCode
	label string
}

func (m MenuItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (id itemDelegate) Height() int                               { return 1 }
func (id itemDelegate) Spacing() int                              { return 0 }
func (id itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (id itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(MenuItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.label)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(strs ...string) string {
			return selectedItemStyle.Render("> " + strs[0])
		}
	}

	fmt.Fprint(w, fn(str))
}

type Menu struct {
	Model list.Model
}

func NewMenu() *Menu {
	const defaultWidth = 20
	const listHeight = 14

	l := list.New(MAIN_MENU, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Database configuration"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = listTitleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return &Menu{
		l,
	}
}

func (m *Menu) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return cmd
}

func (m *Menu) ChangeMenu(menu []list.Item) tea.Cmd {
    return m.Model.SetItems(menu)
}
