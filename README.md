---

# Task Tracker CLI

A simple and lightweight command-line **Task Tracker** written in Go.
It allows you to create and manage tasks directly from your terminal.
All tasks are stored in a JSON file for persistence.

---

## Features

* Add new tasks
* Update existing tasks
* Delete tasks
* Mark tasks as **in progress** or **done**
* List **all** tasks
* List tasks filtered by status:

  *  Done tasks
  * In progress tasks
  * Not done tasks

---

##  Installation

Clone the repository:

```bash
git clone https://github.com/zxCroshka/Task-tracker.git
cd Task-tracker
```

Initialize Go module:

```bash
go mod init github.com/zxCroshka/Task-tracker

```

Build the binary:

```bash
go build .
```

Create an alias for easier usage:

```bash
alias task-cli='./Task-Tracker'
```

---

##  Usage

Add a new task:

```bash
task-cli add "Text of new task"
```

Change status of task:

```bash
task-cli mark new_status task_id
```



Update a task:

```bash
task-cli update task_id "New task description"
```

Delete a task:

```bash
task-cli delete task_id
```

List all tasks:

```bash
task-cli list
```

List tasks by status:

```basg
task-cli list done
task-cli list todo
task-cli list in-progress
```

---

##  Data Storage

All tasks are saved in a local JSON file (e.g., `tasks.json`).
This ensures your tasks remain available between runs.

---

##  Requirements

* Go 1.20+ installed
* Linux/macOS/Windows supported

---



