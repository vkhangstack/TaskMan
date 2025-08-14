package cmd

import (
	"fmt"
	"github.com/vkhangstack/taskman/internal/task"
	"github.com/vkhangstack/taskman/internal/ui"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Long:  `Add a new task to your task list. You can provide the task description as arguments.`,
	Args:  cobra.MinimumNArgs(1),
	Example: `  taskman add "Buy groceries"
  taskman add "Call dentist" --priority high
  taskman add "Review code" -p medium --tags work,urgent`,
	RunE: addTask,
}

var (
	priority string
	tags     []string
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&priority, "priority", "p", "medium", "Task priority (low, medium, high)")
	addCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "Tags for the task")
}

func addTask(cmd *cobra.Command, args []string) error {
	description := strings.Join(args, " ")

	// Validate priority
	if !isValidPriority(priority) {
		return fmt.Errorf("invalid priority: %s. Valid priorities are: low, medium, high", priority)
	}

	store, err := task.NewFileStore()
	if err != nil {
		return fmt.Errorf("failed to initialize task store: %w", err)
	}

	if len(tags) > 0 {
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
			if tag == "" {
				return fmt.Errorf("tag cannot be empty")
			}
		}
	}

	newTask := &task.Task{
		Description: strings.TrimSpace(description),
		Priority:    priority,
		Tags:        tags,
		Status:      task.StatusTodo,
	}

	id, err := store.Add(newTask)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	ui.PrintSuccess(fmt.Sprintf("Task added successfully! (ID: %d)", id))

	// Print the added task
	fmt.Printf("  %s %s\n", ui.FormatPriority(priority), description)
	if len(tags) > 0 {
		fmt.Printf("  Tags: %s\n", ui.FormatTags(tags))
	}

	return nil
}

func isValidPriority(p string) bool {
	validPriorities := []string{"low", "medium", "high"}
	for _, valid := range validPriorities {
		if p == valid {
			return true
		}
	}
	return false
}
