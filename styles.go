package yinlog

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	timestampStyle = lipgloss.NewStyle().Foreground(muted).Render
	iconStyle      = func(color lipgloss.AdaptiveColor, msg string) string {
		return lipgloss.NewStyle().Foreground(color).Render(msg)
	}

	levelStyle = func(color lipgloss.AdaptiveColor, msg string) string {
		return lipgloss.NewStyle().Padding(0).Foreground(color).Render(msg)
	}

	textStyle = func(color lipgloss.AdaptiveColor, msg string) string {
		return lipgloss.NewStyle().Foreground(color).Render(msg)
	}

	// List Styles
	listTitleStyle = lipgloss.NewStyle().Margin(0, 0, 1, 0).
			Padding(0, 1).
			Italic(true).Underline(true).Render

	listItemStyle = func(msg string, indentSize int) string {
		return lipgloss.NewStyle().PaddingLeft(indentSize).Render(msg)
	}
)
