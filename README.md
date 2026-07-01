# Task Tracker CLI

A simple Command Line Interface (CLI) application built with Go to manage daily tasks. This project is based on the Task Tracker project from roadmap.sh and was developed as part of my Go backend learning journey.

## Features

* Add a new task
* Update an existing task
* Delete a task
* Mark a task as **In Progress**
* Mark a task as **Done**
* List all tasks
* Filter tasks by status
* Display task statistics
* Store data in a local JSON file

## Tech Stack

* Go
* Standard Library only
* JSON File Storage

## Project Structure

```text
task-tracker/
│── model/
│── service/
│── storage/
│── main.go
│── go.mod
│── README.md
```

## Getting Started

Clone the repository:

```bash
git clone https://github.com/yourusername/task-tracker.git
```

Move into the project directory:

```bash
cd task-tracker
```

Run the application:

```bash
go run .
```

## Available Commands

### Add a task

```bash
go run . add "Learn Go"
```

### Update a task

```bash
go run . update 1 "Learn Go Advanced"
```

### Delete a task

```bash
go run . delete 1
```

### Mark task as in progress

```bash
go run . mark-in-progress 1
```

### Mark task as done

```bash
go run . mark-done 1
```

### List all tasks

```bash
go run . list
```

### List completed tasks

```bash
go run . list done
```

### List tasks in progress

```bash
go run . list in-progress
```

### List todo tasks

```bash
go run . list todo
```

### Display task statistics

```bash
go run . stats
```

## Learning Objectives

This project helped me practice:

* Go project structure
* Working with JSON files
* File I/O
* CLI development
* Error handling
* Clean Code
* Separation of Concerns
* Basic software architecture

## Future Improvements

* Unit testing
* Logging
* Colored CLI output
* Persistent storage using SQLite
* Interactive CLI

## Acknowledgements

This project is based on the Task Tracker challenge from roadmap.sh.

## Project Source

https://roadmap.sh/projects/task-tracker
