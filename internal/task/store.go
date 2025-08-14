package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// FileStore implements the Store interface using file-based storage
type FileStore struct {
	filePath string
}

// TaskData represents the structure stored in the JSON file
type TaskData struct {
	Tasks    []*Task   `json:"tasks"`
	NextID   int       `json:"next_id"`
	Modified time.Time `json:"modified"`
}

// NewFileStore creates a new FileStore instance
func NewFileStore() (*FileStore, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	dataDir := filepath.Join(home, ".taskman")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	filePath := filepath.Join(dataDir, "tasks.json")

	store := &FileStore{
		filePath: filePath,
	}

	// Initialize file if it doesn't exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := store.save(&TaskData{NextID: 1, Modified: time.Now()}); err != nil {
			return nil, fmt.Errorf("failed to initialize tasks file: %w", err)
		}
	}

	return store, nil
}

// Add adds a new task and returns its ID
func (fs *FileStore) Add(task *Task) (int, error) {
	data, err := fs.load()
	if err != nil {
		return 0, err
	}

	task.ID = data.NextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	data.Tasks = append(data.Tasks, task)
	data.NextID++
	data.Modified = time.Now()

	if err := fs.save(data); err != nil {
		return 0, err
	}

	return task.ID, nil
}

// GetAll returns all tasks
func (fs *FileStore) GetAll() ([]*Task, error) {
	data, err := fs.load()
	if err != nil {
		return nil, err
	}

	// Sort tasks by ID (most recent first)
	sort.Slice(data.Tasks, func(i, j int) bool {
		return data.Tasks[i].ID > data.Tasks[j].ID
	})

	return data.Tasks, nil
}

// GetByID returns a task by its ID
func (fs *FileStore) GetByID(id int) (*Task, error) {
	data, err := fs.load()
	if err != nil {
		return nil, err
	}

	for _, task := range data.Tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", id)
}

// Update updates an existing task
func (fs *FileStore) Update(updatedTask *Task) error {
	data, err := fs.load()
	if err != nil {
		return err
	}

	for i, task := range data.Tasks {
		if task.ID == updatedTask.ID {
			updatedTask.UpdatedAt = time.Now()
			data.Tasks[i] = updatedTask
			data.Modified = time.Now()
			return fs.save(data)
		}
	}

	return fmt.Errorf("task with ID %d not found", updatedTask.ID)
}

// Delete removes a task by its ID
func (fs *FileStore) Delete(id int) error {
	data, err := fs.load()
	if err != nil {
		return err
	}

	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			data.Modified = time.Now()
			return fs.save(data)
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

// Complete marks a task as completed
func (fs *FileStore) Complete(id int) error {
	task, err := fs.GetByID(id)
	if err != nil {
		return err
	}

	task.MarkCompleted()
	return fs.Update(task)
}

// load reads the task data from file
func (fs *FileStore) load() (*TaskData, error) {
	file, err := os.ReadFile(fs.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks file: %w", err)
	}

	var data TaskData
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to parse tasks file: %w", err)
	}

	return &data, nil
}

// save writes the task data to file
func (fs *FileStore) save(data *TaskData) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal task data: %w", err)
	}

	if err := os.WriteFile(fs.filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write tasks file: %w", err)
	}

	return nil
}
