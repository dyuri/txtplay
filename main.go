package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/charmbracelet/bubbletea"
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

func (m model) Init() bubbletea.Cmd {
	return tick()
}

func (m model) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	switch msg := msg.(type) {
	case bubbletea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, bubbletea.Quit
		case "p":
			m.running = !m.running
		}
	case tickMsg:
		if m.running {
			m.current_frame++
			files, _ := ioutil.ReadDir(m.folder)
			if m.current_frame >= len(files) {
				m.current_frame = 0
			}
			content, _ := ioutil.ReadFile(fmt.Sprintf("%s/%d.txt", m.folder, m.current_frame))
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

func tick() bubbletea.Cmd {
	return func() bubbletea.Msg {
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
