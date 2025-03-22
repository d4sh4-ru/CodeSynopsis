package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/d4sh4-ru/CodeSynopsis/internal/ui"
)

func main() {
	p := tea.NewProgram(ui.InitModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
