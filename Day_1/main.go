package main

import (
"fmt"
"os"
"strconv"
"time"
)

func main() {

if len(os.Args) < 2 {
fmt.Println("Usage: task-cli <command>")
return
}

command := os.Args[1]
tasks, _ := loadTasks()

switch command {

case "add":
desc := os.Args[2]
id := len(tasks) + 1

task := Task{
ID:          id,
Description: desc,
Status:      "todo",
CreatedAt:   time.Now(),
UpdatedAt:   time.Now(),
}

tasks = append(tasks, task)
saveTasks(tasks)
fmt.Println("Task added successfully (ID:", id, ")")

case "list":
for _, t := range tasks {
fmt.Println(t.ID, t.Description, t.Status)
}

case "delete":
id, _ := strconv.Atoi(os.Args[2])
newTasks := []Task{}

for _, t := range tasks {
if t.ID != id {
newTasks = append(newTasks, t)
}
}

saveTasks(newTasks)
fmt.Println("Task deleted")

case "mark-done":
id, _ := strconv.Atoi(os.Args[2])

for i := range tasks {
if tasks[i].ID == id {
tasks[i].Status = "done"
tasks[i].UpdatedAt = time.Now()
}
}

saveTasks(tasks)
fmt.Println("Task marked as done")

case "mark-in-progress":
id, _ := strconv.Atoi(os.Args[2])

for i := range tasks {
if tasks[i].ID == id {
tasks[i].Status = "in-progress"
tasks[i].UpdatedAt = time.Now()
}
}

saveTasks(tasks)
fmt.Println("Task marked as in progress")

default:
fmt.Println("Invalid command")
}
}
