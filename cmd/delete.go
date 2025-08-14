package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vkhangstack/taskman/internal/task"
	"github.com/vkhangstack/taskman/internal/ui"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Long:  `Delete a task by providing its ID. This will remove the task from your task list permanently.`,
	Args:  cobra.MinimumNArgs(1),
	Example: `  taskman delete 1
  taskman del 2
  taskman remove 3`,
	RunE: deleteTasks,
}
var forceDelete bool

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Force delete the task without confirmation")
}

func deleteTasks(cmd *cobra.Command, args []string) error {
	store, err := task.NewFileStore()
	if err != nil {
		return err
	}
	var toDelete []int
	var errors []error

	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", arg)
			continue
		}

		// Check if the task exists
		_, err = store.GetByID(id)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to retrieve task with ID %d: %w", id, err))
			continue
		}
		toDelete = append(toDelete, id)
	}
	if len(toDelete) == 0 {
		for _, err := range errors {
			ui.PrintError(err.Error())
		}
		if len(errors) <= 0 {
			return nil
		}
	}

	if !forceDelete && len(toDelete) > 0 {
		ui.PrintWarning("You are about to delete the following tasks:")
		for _, id := range toDelete {
			ui.PrintInfo(fmt.Sprintf("Task ID: %d", id))
		}
		ui.PrintWarning("This action cannot be undone. Are you sure you want to proceed? (yes/no)")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
			ui.PrintInfo("Deletion cancelled.")
			return nil
		}
	}
	// delete tasks
	var deleted []int
	var deleteErrors []string

	for _, id := range toDelete {
		if err := store.Delete(id); err != nil {
			deleteErrors = append(deleteErrors, fmt.Sprintf("failed to delete task with ID %d: %v", id, err))
		} else {
			deleted = append(deleted, id)
		}
	}
	// Print results
	if len(deleted) > 0 {
		if len(deleted) == 1 {
			ui.PrintSuccess(fmt.Sprintf("Task with ID %d deleted successfully!", deleted[0]))
		} else {
			ui.PrintSuccess(fmt.Sprintf("%d tasks deleted successfully!", len(deleted)))
			for _, id := range deleted {
				ui.PrintInfo(fmt.Sprintf("Task ID: %d", id))
			}
		}
	}
	if len(deleteErrors) > 0 {
		ui.PrintError("Some tasks could not be deleted:")
		for _, err := range deleteErrors {
			ui.PrintError(err)
		}
	}
	return nil
}
