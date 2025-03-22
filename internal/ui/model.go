package ui

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/d4sh4-ru/CodeSynopsis/internal/fileutils"
	"github.com/d4sh4-ru/CodeSynopsis/internal/markdown"
)

type Model struct {
	list          list.Model
	selectedItems map[string]bool
	currentDir    string
	dirStack      []string
	width         int
	height        int
	showHelp      bool
	helpWidth     int
	helpMargin    int
	helpStyle     lipgloss.Style
}

// Init
func InitModel() Model {
	m := Model{
		selectedItems: make(map[string]bool),
		currentDir:    ".",
		dirStack:      []string{},
		showHelp:      true, // Показываем справку по умолчанию
		helpWidth:     30,   // Ширина панели
		helpMargin:    2,    // Отступ справки
		helpStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Margin(0, 1),
	}

	items := getFilteredDirectoryItems(m.currentDir)
	delegate := itemDelegate{
		SelectedItems:  m.selectedItems,
		DirectoryStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true),
	}

	l := list.New(items, delegate, 0, 0)
	l.Title = "Выберите файлы и директории"
	l.SetShowHelp(false)
	m.list = l

	return m
}

func (m Model) helpView() string {
	helpText := " Справка:\n" +
		"  F1 - Скрыть/показать справку\n" +
		"  ↑/↓ - Навигация\n" +
		"  Enter - Открыть директорию\n" +
		"  Backspace - Назад\n" +
		"  Пробел - Выбрать элемент\n" +
		"  Alt+E - Экспорт в Markdown\n" +
		"  q - Выход"

	return m.helpStyle.Render(helpText)
}

// Функция для получения отфильтрованных элементов в текущей директории
func getFilteredDirectoryItems(dir string) []list.Item {
	entries, err := fileutils.GetSortedDirEntries(dir)
	if err != nil {
		return []list.Item{item{Name: fmt.Sprintf("Error reading directory: %v", err)}}
	}

	var items []list.Item

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		items = append(items, item{
			Name:  entry.Name(),
			Path:  path,
			IsDir: entry.IsDir(),
		})
	}
	return items
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update обрабатывает действия пользователя и выполняет необходимые действия.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "f1":
			m.showHelp = !m.showHelp
			if m.showHelp {
				m.list.SetWidth(m.width - m.helpWidth - m.helpMargin*2 - 2)
			} else {
				m.list.SetWidth(m.width)
			}
			return m, nil

		case "enter":
			selectedItem := m.list.SelectedItem().(item)
			if selectedItem.IsDir {
				m.dirStack = append(m.dirStack, m.currentDir)
				m.currentDir = selectedItem.Path
				m.list.SetItems(getFilteredDirectoryItems(m.currentDir))
				m.list.ResetSelected()
			}
			return m, nil

		case "backspace":
			if len(m.dirStack) > 0 {
				m.currentDir = m.dirStack[len(m.dirStack)-1]
				m.dirStack = m.dirStack[:len(m.dirStack)-1]
				m.list.SetItems(getFilteredDirectoryItems(m.currentDir))
				m.list.ResetSelected()
			}
			return m, nil

		case " ":
			selectedItem := m.list.SelectedItem().(item)
			m.selectedItems[selectedItem.Path] = !m.selectedItems[selectedItem.Path]
			m.list.SetItems(getFilteredDirectoryItems(m.currentDir))
			return m, nil

		case "alt+e":
			markdown.GenerateMarkdown(m.selectedItems)
			return m, tea.Quit

		case "q":
			return m, tea.Quit

		default:
			// Передаем необработанные клавиши в список
			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		listWidth := msg.Width
		if m.showHelp {
			listWidth = msg.Width - m.helpWidth - m.helpMargin*2 - 2
		}

		m.list.SetSize(listWidth, msg.Height-2)
		m.helpStyle.Width(m.helpWidth)

		return m, nil

	default:
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
}

// View отображает список и, если необходимо, справку.
func (m Model) View() string {
	listView := m.list.View()

	// Добавляем условие отображения справки
	if m.showHelp {
		helpView := m.helpView()
		return lipgloss.JoinHorizontal(lipgloss.Top, listView, helpView)
	}
	return listView
}
