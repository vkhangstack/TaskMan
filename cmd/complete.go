package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vkhangstack/taskman/internal/task"
	"github.com/vkhangstack/taskman/internal/ui"
	"strconv"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Complete a task",
	Long:  `Mark a task as completed by providing its ID. This will update the task status to 'completed'.`,
	Args:  cobra.ExactArgs(1),
	Example: `taskman complete 1
  	taskman complete 1 2 3
	taskman done 5 --force`,
	RunE: completeTask,
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
func completeTask(cmd *cobra.Command, args []string) error {
	// Convert args to integers
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", arg)
		}
		ids = append(ids, id)
	}

	store, err := task.NewFileStore()
	if err != nil {
		return fmt.Errorf("failed to initialize task store: %w", err)
	}

	for _, id := range ids {
		if err := store.Complete(id); err != nil {
			return fmt.Errorf("failed to complete task with ID %d: %w", id, err)
		}
	}

	ui.PrintSuccess("Tasks completed successfully!")
	return nil
}
