package ui

import (
	"fmt"
	"github.com/vkhangstack/taskman/internal/task"
	"strings"

	"github.com/fatih/color"
)

var (
	// Text colors
	RedText    = color.New(color.FgRed)
	GreenText  = color.New(color.FgGreen)
	YellowText = color.New(color.FgYellow)
	BlueText   = color.New(color.FgBlue)
	CyanText   = color.New(color.FgCyan)
	WhiteText  = color.New(color.FgWhite)

	// Bold colors
	RedBold    = color.New(color.FgRed, color.Bold)
	GreenBold  = color.New(color.FgGreen, color.Bold)
	YellowBold = color.New(color.FgYellow, color.Bold)
	BlueBold   = color.New(color.FgBlue, color.Bold)
	CyanBold   = color.New(color.FgCyan, color.Bold)

	// Background colors
	RedBg    = color.New(color.BgRed, color.FgWhite)
	GreenBg  = color.New(color.BgGreen, color.FgWhite)
	YellowBg = color.New(color.BgYellow, color.FgBlack)
)

// PrintSuccess prints a success message in green
func PrintSuccess(message string) {
	fmt.Printf("%s %s\n", GreenBold.Sprint("✓"), GreenText.Sprint(message))
}

// PrintError prints an error message in red
func PrintError(message string) {
	fmt.Printf("%s %s\n", RedBold.Sprint("✗"), RedText.Sprint(message))
}

// PrintWarning prints a warning message in yellow
func PrintWarning(message string) {
	fmt.Printf("%s %s\n", YellowBold.Sprint("⚠"), YellowText.Sprint(message))
}

// PrintInfo prints an info message in blue
func PrintInfo(message string) {
	fmt.Printf("%s %s\n", BlueBold.Sprint("ℹ"), BlueText.Sprint(message))
}

// FormatPriority returns a colored priority indicator
func FormatPriority(priority string) string {
	switch strings.ToLower(priority) {
	case task.PriorityHigh:
		return RedBold.Sprint("HIGH  ")
	case task.PriorityMedium:
		return YellowBold.Sprint("MEDIUM")
	case task.PriorityLow:
		return GreenText.Sprint("LOW   ")
	default:
		return WhiteText.Sprint("    ")
	}
}

// FormatStatus returns a colored status indicator
func FormatStatus(status string) string {
	switch strings.ToLower(status) {
	case task.StatusCompleted:
		return GreenBold.Sprint("DONE")
	case task.StatusTodo:
		return YellowText.Sprint("TODO")
	case task.StatusDeleted:
		return RedText.Sprint("DELETE")
	case task.StatusArchived:
		return CyanText.Sprint("ARCHIVE")
	case task.StatusInProgress:
		return BlueText.Sprint("PROGRESS")
	default:
		return WhiteText.Sprint("UNKNOWN")
	}
}

// FormatTags returns formatted tags with colors
func FormatTags(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	var formatted []string
	for _, tag := range tags {
		formatted = append(formatted, CyanText.Sprintf("#%s", tag))
	}

	return strings.Join(formatted, " ")
}

// FormatID returns a formatted task ID
func FormatID(id int) string {
	return BlueBold.Sprintf("#%d", id)
}

// FormatDescription returns formatted description based on status
func FormatDescription(description, status string) string {
	if status == task.StatusCompleted {
		return color.New(color.CrossedOut).Sprint(description)
	}
	return description
}
