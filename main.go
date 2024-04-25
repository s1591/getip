package main

import tea "github.com/charmbracelet/bubbletea"

func main() {

	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		panic(err)
	}

}
