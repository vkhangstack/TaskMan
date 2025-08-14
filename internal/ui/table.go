package ui

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/vkhangstack/taskman/internal/task"
	"os"
)

// DisplayTasksTable displays tasks in a formatted table
func DisplayTasksTable(tasks []*task.Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Priority", "Description", "Tags", "Created"})

	// Configure table appearance
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	// Add rows
	for _, t := range tasks {
		row := []string{
			FormatID(t.ID),
			FormatStatus(t.Status),
			FormatPriority(t.Priority),
			FormatDescription(t.Description, t.Status),
			FormatTags(t.Tags),
			t.CreatedAt.Format("02/01/2006 15:04"),
		}
		table.Append(row)
	}

	table.Render()
}

// DisplayTaskDetails displays detailed information about a single task
func DisplayTaskDetails(t *task.Task) {
	fmt.Printf("\n")
	fmt.Printf("Task %s\n", FormatID(t.ID))
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	fmt.Printf("Description: %s\n", FormatDescription(t.Description, t.Status))
	fmt.Printf("Status:      %s\n", FormatStatus(t.Status))
	fmt.Printf("Priority:    %s\n", FormatPriority(t.Priority))

	if len(t.Tags) > 0 {
		fmt.Printf("Tags:        %s\n", FormatTags(t.Tags))
	}

	fmt.Printf("Created:     %s\n", t.CreatedAt.Format("02/01/2006 15:04"))
	fmt.Printf("Updated:     %s\n", t.UpdatedAt.Format("02/01/2006 15:04"))

	if t.CompletedAt != nil {
		fmt.Printf("Completed:   %s\n", t.CompletedAt.Format("02/01/2006 15:04"))
	}

	fmt.Printf("\n")
}

// DisplayTasksSummary displays a summary of tasks by status and priority
func DisplayTasksSummary(tasks []*task.Task) {
	pending := 0
	completed := 0
	high := 0
	medium := 0
	low := 0

	for _, t := range tasks {
		switch t.Status {
		case task.StatusPending:
		case task.StatusTodo:
		case task.StatusInProgress:
			pending++
		case task.StatusCompleted:
			completed++
		}

		switch t.Priority {
		case task.PriorityHigh:
			high++
		case task.PriorityMedium:
			medium++
		case task.PriorityLow:
			low++
		}
	}

	fmt.Printf("\n")
	fmt.Printf("Task Summary\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Total tasks:     %d\n", len(tasks))
	fmt.Printf("Pending:         %s (%d)\n", YellowText.Sprint("⏳"), pending)
	fmt.Printf("Completed:       %s (%d)\n", GreenText.Sprint("✓"), completed)
	fmt.Printf("\n")
	fmt.Printf("By Priority:\n")
	fmt.Printf("High:            %s (%d)\n", RedText.Sprint("●"), high)
	fmt.Printf("Medium:          %s (%d)\n", YellowText.Sprint("●"), medium)
	fmt.Printf("Low:             %s (%d)\n", GreenText.Sprint("●"), low)
	fmt.Printf("\n")
}

// truncateString truncates a string to the specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// FormatDuration returns a human-readable duration string
func FormatDuration(t *task.Task) string {
	if t.CompletedAt == nil {
		return ""
	}

	duration := t.CompletedAt.Sub(t.CreatedAt)
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24

	if days > 0 {
		return fmt.Sprintf("(%dd %dh)", days, hours)
	} else if hours > 0 {
		return fmt.Sprintf("(%dh)", hours)
	}
	return "(< 1h)"
}
