Task Tracker CLI

A simple command-line task tracking application built in Go.

This CLI allows you to manage tasks directly from the terminal. Tasks are stored in a local tasks.json file in the current directory.

Features

Add tasks

Update tasks

Delete tasks

Mark tasks as:

todo

in-progress

done

List all tasks

Filter tasks by status

Automatically creates tasks.json if it does not exist

No external libraries used

Installation

Clone the repository:

git clone <your-repo-url>
cd task-tracker

Initialize Go module (if not already initialized):

go mod init task-tracker

Build the project:

go build -o task-tracker
Usage
Show Help
./task-tracker --help
Commands
Add a Task
./task-tracker add "Buy groceries"

Output:

Task added successfully (ID: 1)
Update a Task
./task-tracker update 1 "Buy groceries and cook dinner"
Delete a Task
./task-tracker delete 1
Mark Task Status

Mark as in progress:

./task-tracker mark-in-progress 1

Mark as done:

./task-tracker mark-done 1

Mark as todo:

./task-tracker mark-todo 1
List Tasks

List all tasks:

./task-tracker list

List only completed tasks:

./task-tracker list done

List only pending tasks:

./task-tracker list todo

List tasks in progress:

./task-tracker list in-progress
Task Structure

Each task contains the following fields:

id – Unique identifier

description – Task description

status – Current status (todo, in-progress, done)

createdAt – Creation timestamp

updatedAt – Last updated timestamp

Data Storage

All tasks are stored in:

tasks.json

The file is automatically created in the project directory if it does not exist.

Error Handling

The application handles:

Invalid commands

Missing arguments

Invalid task IDs

Non-existent tasks

Invalid status values

Design Decisions

Uses Go standard library only

JSON file-based persistence

Simple positional argument parsing

Stateless CLI execution (no background service)

Future Improvements

Monotonic ID generator

Concurrency-safe writes

HTTP REST version

In-memory caching layer

Rate limiting

Docker support
