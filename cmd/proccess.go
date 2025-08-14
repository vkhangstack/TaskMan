package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vkhangstack/taskman/internal/task"
	"github.com/vkhangstack/taskman/internal/ui"
	"strconv"
	"time"
)

var processingCmd = &cobra.Command{
	Use:   "process",
	Short: "Process tasks",
	Long:  `Process tasks based on their status, priority, or tags. This command allows you to perform batch operations on tasks.`,
	Example: `  taskman process 1
  taskman process 2 --priority high
  taskman process 3 4 5 -p medium`,
	Args: cobra.MinimumNArgs(1),
	RunE: processingTask,
}

func init() {
	rootCmd.AddCommand(processingCmd)

	processingCmd.Flags().StringP("priority", "p", "", "Filter tasks by priority (low, medium, high)")
}

func processingTask(cmd *cobra.Command, args []string) error {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", arg)
		}
		ids = append(ids, id)
	}

	if !isValidPriority(priority) {
		return fmt.Errorf("invalid priority: %s. Valid priorities are: low, medium, high", priority)
	}

	store, err := task.NewFileStore()

	if err != nil {
		return fmt.Errorf("failed to initialize task store: %w", err)
	}
	for _, id := range ids {
		taskRecord, err := store.GetByID(id)
		if err != nil {
			return fmt.Errorf("failed to retrieve task with ID %d: %w", id, err)
		}

		switch taskRecord.Status {
		case task.StatusTodo:
			fmt.Printf("Processing pending task: %s\n", taskRecord.Description)
			// Add your processing logic here
			taskRecord.Status = task.StatusInProgress
			taskRecord.UpdatedAt = time.Now()
			if priority != "" {
				taskRecord.Priority = priority
			}
			if err := store.Update(taskRecord); err != nil {
				return fmt.Errorf("failed to update task with ID %d: %w", id, err)
			}
			fmt.Printf("Task %d is now in progress.\n", id)
		case task.StatusInProgress:
			fmt.Printf("Task %d is already in progress.\n", id)
		case task.StatusCompleted:
			fmt.Printf("Task %d is already completed.\n", id)
		case task.StatusDeleted:
			fmt.Printf("Task %d is deleted and cannot be processed.\n", id)
		default:
			fmt.Printf("Unknown status for task %d: %s\n", id, taskRecord.Status)
		}
	}
	ui.PrintSuccess("Tasks processed successfully!")
	return nil

}
