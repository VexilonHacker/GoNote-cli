# To-Do List App in Go

A simple command-line To-Do list application written in Go, which uses a CSV file to store tasks and their statuses.

---

## Features

- Add multiple tasks at once, separated by commas
- View all tasks in a formatted table with status and timestamps
- Mark tasks as completed with a timestamp
- Clear all tasks from the database
- Clear the terminal screen
- Persistent task storage in `notes.csv`

---

## How It Works

- Tasks are stored in a CSV file with columns: ID, Task, Status, Timestamp, Completed At
- Status shows a green check mark (✅) for completed tasks and a red cross (❌) for incomplete tasks
- Tasks have unique incremental IDs for easy reference
- Timestamps track when tasks are added and completed

---

## Usage

Run the program and select options from the menu:

1. Add notes  
2. Show notes  
3. Complete a task  
4. Clear terminal  
5. Clear notes database  
6. Exit  

---

## Dependencies

- [Go](https://golang.org/dl/)  
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) Go package for rendering tables  

Install dependencies using:

```bash
go get github.com/olekukonko/tablewriter

