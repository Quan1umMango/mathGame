package main

import (
  "log"
  "QuantumMango/mathGame/cli"
  tea "github.com/charmbracelet/bubbletea"
)



func main() {

  p := tea.NewProgram(cli.NewInputModel())

  if _, err := p.Run(); err != nil {
    log.Fatal(err)
  }
}
