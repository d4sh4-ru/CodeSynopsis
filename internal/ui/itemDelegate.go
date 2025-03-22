package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/d4sh4-ru/CodeSynopsis/internal/utils"
)

type itemDelegate struct {
	SelectedItems  map[string]bool
	DirectoryStyle lipgloss.Style
}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

// Render отрисовывает страницу в терминале пользователя.
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i := listItem.(item)
	checked := d.SelectedItems[i.Path]

	// Определяем символ для элемента
	var symbol string
	if i.IsDir {
		symbol = "📁" // Символ для директорий
	} else {
		symbol = utils.GetFileSymbol(i.Name) // Определяем символ для файла
	}

	// Определяем стиль для элемента
	var displayName string
	if i.IsDir {
		displayName = d.DirectoryStyle.Render(i.Name)
	} else {
		displayName = i.Name
	}

	// Формируем строку с чекбоксом
	str := fmt.Sprintf("%s %s", symbol, displayName)
	if checked {
		str = "• " + str
	} else {
		str = "  " + str
	}

	// Добавляем маркер текущего элемента
	cursor := " "
	if m.Index() == index {
		cursor = ">"
	}

	fmt.Fprintf(w, "%s %s", cursor, str)
}
