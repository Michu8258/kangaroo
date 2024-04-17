package models

import "github.com/charmbracelet/lipgloss"

type styles struct {
	DefaultStyle lipgloss.Style
	PrimaryStyle lipgloss.Style
	SuccessStyle lipgloss.Style
	ErrorStyle   lipgloss.Style
	BorderStyle  lipgloss.Style
	DebugStyle   lipgloss.Style
}

var TerminalStyles = styles{
	DefaultStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#eeeeee")),
	PrimaryStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#7777ff")).Bold(true),
	SuccessStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#33ff33")),
	ErrorStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#ff3333")),
	BorderStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")),
	DebugStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#ffd700")),
}
