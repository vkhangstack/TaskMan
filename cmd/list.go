package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vkhangstack/taskman/internal/task"
	"github.com/vkhangstack/taskman/internal/ui"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks in your task list. You can filter by status, priority, or tags.`,
	Example: `  taskman list
  taskman list --status completed
  taskman list --priority high
  taskman list --tags work,urgent`,
	RunE: listTasks,
}

var (
	statusFilter    string
	priorityFilter  string
	tagsFilter      []string
	completedFilter bool
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filter tasks by status (pending, completed)")
	listCmd.Flags().StringVarP(&priorityFilter, "priority", "p", "", "Filter tasks by priority (low, medium, high)")
	listCmd.Flags().StringSliceVarP(&tagsFilter, "tags", "t", []string{}, "Filter tasks by tags (comma-separated)")
	listCmd.Flags().BoolVar(&completedFilter, "completed", false, "Show only completed tasks")

}
func listTasks(cmd *cobra.Command, args []string) error {
	store, err := task.NewFileStore()
	if err != nil {
		return err
	}

	tasks, err := store.GetAll()
	if err != nil {
		return err
	}
	filteredTasks := filterTasks(tasks)
	if len(filteredTasks) == 0 {
		ui.PrintInfo("No tasks found matching the filters.")
		return nil
	}
	for _, t := range filteredTasks {
		fmt.Println(t.Description)
		fmt.Println(len(t.Description))
	}

	ui.DisplayTasksTable(filteredTasks)
	showSummary(filteredTasks)
	return nil
}
func filterTasks(tasks []*task.Task) []*task.Task {
	var filteredTasks []*task.Task

	for _, t := range tasks {
		if statusFilter != "" && t.Status != statusFilter {
			continue
		}
		if priorityFilter != "" && t.Priority != priorityFilter {
			continue
		}
		if len(tagsFilter) > 0 {
			matched := false
			for _, tag := range tagsFilter {
				if t.HasTag(tag) {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}
		if completedFilter && !t.IsCompleted() {
			continue
		}

		filteredTasks = append(filteredTasks, t)
	}

	return filteredTasks

}

func containsTag(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func showSummary(tasks []*task.Task) {
	if len(tasks) == 0 {
		ui.PrintInfo("No tasks found.")
		return
	}

	ui.DisplayTasksSummary(tasks)
}
