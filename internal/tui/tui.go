package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/veggiemonk/awesome-docker/internal/cache"
)

// Run launches the TUI browser. It blocks until the user quits.
func Run(entries []cache.HealthEntry) error {
	m := New(entries)
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}
