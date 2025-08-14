package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vkhangstack/taskman/internal/task"
	"github.com/vkhangstack/taskman/internal/ui"
	"strconv"
	"time"
)

var undoCmd = &cobra.Command{
	Use:   "undo [task ID]",
	Short: "Undo the last action on a task",
	Long:  `Undo the last action performed on a task, such as marking it as completed or deleting it. This will restore the task to its previous state.`,
	Args:  cobra.MinimumNArgs(1),
	Example: `  taskman undo 1
	  taskman undo 2 3 
	  taskman undo 3 4 5`,
	RunE: undoTask,
}

func init() {
	rootCmd.AddCommand(undoCmd)
}

func undoTask(cmd *cobra.Command, args []string) error {
	store, err := task.NewFileStore()
	if err != nil {
		return err
	}

	var errors []error
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			errors = append(errors, fmt.Errorf("invalid task ID: %s", arg))
			continue
		}
		// Check if the task exists
		taskRecord, err := store.GetByID(id)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to retrieve task with ID %d: %w", id, err))
			continue
		}

		if taskRecord.Status == task.StatusCompleted {
			errors = append(errors, fmt.Errorf("task with ID %d is already completed and cannot be undone", id))
			continue
		}
		if taskRecord.Status == task.StatusDeleted {
			errors = append(errors, fmt.Errorf("task with ID %d is already deleted and cannot be undone", id))
			continue
		}
		taskRecord.Status = task.StatusPending // Reset status to pending
		taskRecord.CompletedAt = nil           // Clear completed timestamp
		taskRecord.UpdatedAt = time.Now()

		if err := store.Update(taskRecord); err != nil {
			errors = append(errors, fmt.Errorf("failed to undo action for task with ID %d: %w", id, err))
		}
	}

	if len(errors) > 0 {
		for _, err := range errors {
			ui.PrintError(err.Error())
		}
		return nil
	}

	ui.PrintSuccess("Tasks undone successfully!")
	return nil

}
