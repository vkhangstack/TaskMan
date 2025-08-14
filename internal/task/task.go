package task

import (
	"time"
)

var (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDeleted    = "deleted"
	StatusArchived   = "archived"
	StatusPending    = "pending"
	StatusCompleted  = "completed"
	PriorityLow      = "low"
	PriorityMedium   = "medium"
	PriorityHigh     = "high"
)

// Task represents a task item
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      string     `json:"status"`   // pending, completed
	Priority    string     `json:"priority"` // low, medium, high
	Tags        []string   `json:"tags"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// Store defines the interface for task storage
type Store interface {
	Add(task *Task) (int, error)
	GetAll() ([]*Task, error)
	GetByID(id int) (*Task, error)
	Update(task *Task) error
	Delete(id int) error
	Complete(id int) error
	MarkPending(id int) error
	MarkTodo(id int) error
	MarkInProgress(id int) error
	MarkDeleted(id int) error
	MarkArchived(id int) error
	MarkCompleted(id int) error
	MarkStatus(id int, status string) error
	IsCompleted(id int) bool
	IsPending(id int) bool
	IsHighPriority(id int) bool
	IsLowPriority(id int) bool
	IsMediumPriority(id int) bool
	HasTag(id int, tag string) bool
	AddTag(id int, tag string) error
	RemoveTag(id int, tag string) error
}

// IsCompleted returns true if the task is completed
func (t *Task) IsCompleted() bool {
	return t.Status == StatusCompleted
}

// IsPending returns true if the task is pending
func (t *Task) IsPending() bool {
	return t.Status == StatusPending
}

// IsHighPriority returns true if the task has high priority
func (t *Task) IsHighPriority() bool {
	return t.Priority == PriorityHigh
}

// IsLowPriority returns true if the task has low priority
func (t *Task) IsLowPriority() bool {
	return t.Priority == PriorityLow
}

// IsMediumPriority returns true if the task has medium priority
func (t *Task) IsMediumPriority() bool {
	return t.Priority == PriorityMedium
}

// HasTag returns true if the task has the specified tag
func (t *Task) HasTag(tag string) bool {
	for _, t := range t.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// AddTag adds a tag to the task if it doesn't already exist
func (t *Task) AddTag(tag string) {
	if !t.HasTag(tag) {
		t.Tags = append(t.Tags, tag)
		t.UpdatedAt = time.Now()
	}
}

// RemoveTag removes a tag from the task
func (t *Task) RemoveTag(tag string) {
	for i, tagName := range t.Tags {
		if tagName == tag {
			t.Tags = append(t.Tags[:i], t.Tags[i+1:]...)
			t.UpdatedAt = time.Now()
			break
		}
	}
}

// MarkCompleted marks the task as completed
func (t *Task) MarkCompleted() {
	t.Status = StatusCompleted
	now := time.Now()
	t.CompletedAt = &now
	t.UpdatedAt = now
}

// MarkPending marks the task as pending
func (t *Task) MarkPending() {
	t.Status = StatusPending
	t.CompletedAt = nil
	t.UpdatedAt = time.Now()
}
func (t *Task) MarkTodo() {
	t.Status = StatusTodo
	t.CompletedAt = nil
	t.UpdatedAt = time.Now()
}
func (t *Task) MarkInProgress() {
	t.Status = StatusInProgress
	t.CompletedAt = nil
	t.UpdatedAt = time.Now()
}
func (t *Task) MarkDeleted() {
	t.Status = StatusDeleted
	t.CompletedAt = nil
	t.UpdatedAt = time.Now()
}
func (t *Task) MarkArchived() {
	t.Status = StatusArchived
	t.CompletedAt = nil
	t.UpdatedAt = time.Now()
}
func (t *Task) MarkStatus(status string) {
	switch status {
	case StatusTodo:
		t.MarkTodo()
	case StatusInProgress:
		t.MarkInProgress()
	case StatusDeleted:
		t.MarkDeleted()
	case StatusArchived:
		t.MarkArchived()
	case StatusPending:
		t.MarkPending()
	case StatusCompleted:
		t.MarkCompleted()
	default:
		t.Status = status
	}
}
