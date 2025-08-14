# TaskMan CLI

A simple, powerful task management CLI application built with Go and Cobra.

## Features

- ✅ Add, list, complete, and delete tasks
- 🎯 Priority levels (low, medium, high)
- 🏷️ Tag system for task organization
- 🎨 Colorful, intuitive interface
- 💾 File-based storage (JSON)
- 🔍 Advanced filtering options
- 📊 Task summaries and statistics

## Installation

### From Source

```bash
git clone https://github.com/vkhangstack/taskman.git
cd taskman
make install
```

### Using Go Install

```bash
go install github.com/vkhangstack/taskman@latest
```

### Download Binary

Download the latest release from the [releases page](https://github.com/vkhangstack/taskman/releases).

## Quick Start

```bash
# Add a new task
taskman add "Buy groceries"

# Add a task with priority and tags
taskman add "Review code" --priority high --tags work,urgent

#  Update task in progress 
taskman progress 1

# List all tasks
taskman list

# List only pending tasks
taskman list --status pending

# Complete a task
taskman complete 1

# Delete a task
taskman delete 2
```

## Usage

### Adding Tasks

```bash
# Basic task
taskman add "Call dentist"

# Task with priority
taskman add "Finish project" --priority high

# Task with tags
taskman add "Read book" --tags personal,learning

# Task with both priority and tags
taskman add "Team meeting" --priority medium --tags work,meeting
```
### Progressing Tasks

```bash
# Mark a task as in progress
taskman progress 1
taskman progress 2 3 4  # Mark multiple tasks as in progress
``` 

### Undo Progressing Tasks

```bash
# Undo progress on a task
taskman undo 1
taskman undo 2 3  # Undo progress on multiple tasks
```

### Listing Tasks

```bash
# List all tasks
taskman list

# Filter by status
taskman list --status pending
taskman list --status completed

# Filter by priority
taskman list --priority high

# Filter by tag
taskman list --tag work

# Show only completed tasks
taskman list --completed
```

### Managing Tasks

```bash
# Complete tasks
taskman complete 1
taskman complete 1 2 3

# Delete tasks
taskman delete 1
taskman delete 1 2 3 --force  # Skip confirmation

# Show version
taskman version
```

## Configuration

TaskMan stores tasks in `~/.taskman/tasks.json` and looks for configuration in `~/.taskman.yaml`.

Example configuration:

```yaml
# ~/.taskman.yaml
verbose: true
default_priority: medium
date_format: "Jan 02, 2006"
```

## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for convenience)

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Development build with race detection
make dev

# Run tests
make test

# Run tests with coverage
make test-coverage

# Format and lint code
make fmt
make lint
```

### Project Structure

```
taskman/
├── main.go              # Entry point
├── cmd/                 # Command definitions
│   ├── root.go         # Root command and config
│   ├── add.go          # Add command
│   ├── list.go         # List command
│   ├── complete.go     # Complete command
│   └── delete.go       # Delete command
├── internal/           # Internal packages
│   ├── task/           # Task domain logic
│   │   ├── task.go     # Task struct and methods
│   │   └── store.go    # File storage implementation
│   └── ui/             # User interface utilities
│       ├── colors.go   # Color formatting
│       └── table.go    # Table display
├── pkg/                # Public packages
└── configs/            # Configuration files
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Examples

### Daily Workflow

```bash
# Start your day by checking tasks
taskman list

# Add tasks as they come up
taskman add "Email client about proposal" --priority high --tags work
taskman add "Buy milk" --tags personal
taskman add "Call mom" --priority medium --tags personal

# Work on high priority items first
taskman list --priority high

# Complete tasks as you finish them
taskman complete 1
taskman complete 3

# Review what's left
taskman list --status pending

# Clean up completed tasks periodically
taskman list --completed
taskman delete 5 6 7 --force
```

### Project Management

```bash
# Add project tasks
taskman add "Design database schema" --priority high --tags project,backend
taskman add "Create API endpoints" --priority high --tags project,backend  
taskman add "Write unit tests" --priority medium --tags project,testing
taskman add "Update documentation" --priority low --tags project,docs

# Focus on project work
taskman list --tag project

# Track backend tasks
taskman list --tag backend

# Complete implementation tasks
taskman complete 1 2

# Review testing tasks
taskman list --tag testing
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Color](https://github.com/fatih/color) - Terminal colors
- [TableWriter](https://github.com/olekukonko/tablewriter) - ASCII tables