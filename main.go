package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type model struct {
	width         int
	height        int
	content       string
	folder        string
	current_frame int
	running       bool
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "p":
			m.running = !m.running
		}
	case tickMsg:
		if m.running {
			m.current_frame++
			files, _ := os.ReadDir(m.folder)
			if m.current_frame >= len(files) {
				m.current_frame = 0
			}
			content, _ := os.ReadFile(fmt.Sprintf("%s/%d.txt", m.folder, m.current_frame))
			m.content = string(content)
		}
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	return m.content
}

type tickMsg struct{}

func tick() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(20 * time.Millisecond)
		return tickMsg{}
	}
}

var rootCmd = &cobra.Command{
	Use:   "txtplay",
	Short: "Play text art animations in the terminal",
}

var playCmd = &cobra.Command{
	Use:   "play [folder]",
	Short: "Play the animation from the specified folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		folder := args[0]
		p := tea.NewProgram(model{folder: folder, running: true})
		if err := p.Start(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func main() {
	rootCmd.AddCommand(playCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
