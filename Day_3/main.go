package main

import(
    "fmt"
    "os"
    "time"
"strconv"
)

func main() {

if len(os.Args) < 2 {
fmt.Println("Usage: expense-tracker <command>")
return
}

command := os.Args[1]

expenses, err := loadExpenses()

if err != nil {
fmt.Println("Error loading expenses:", err)
return
}

switch command {

case "add":

if len(os.Args) < 4 {
fmt.Println("Usage: add <description> <amount>")
return
}

description := os.Args[2]

amount, err := strconv.Atoi(os.Args[3])

if err != nil {
fmt.Println("Invalid amount")
return
}

id := len(expenses) + 1

date := time.Now().Format("2006-01-02")

e := Expense{
ID:          id,
Description: description,
Amount:      amount,
Date:        date,
}

expenses = append(expenses, e)

saveExpenses(expenses)

fmt.Println("Expense added successfully (ID:", id, ")")

case "list":

fmt.Println("ID  Date       Description  Amount")

for _, e := range expenses {
fmt.Printf("%d   %s  %s       $%d\n",
e.ID, e.Date, e.Description, e.Amount)
}

case "summary":

total := 0

if len(os.Args) == 4 && os.Args[2] == "month" {

month, err := strconv.Atoi(os.Args[3])

if err != nil {
fmt.Println("Invalid month")
return
}

for _, e := range expenses {

t, _ := time.Parse("2006-01-02", e.Date)

if int(t.Month()) == month {
total += e.Amount
}
}

fmt.Println("Total expenses for month:", total)

} else {

for _, e := range expenses {
total += e.Amount
}

fmt.Println("Total expenses:", total)
}

case "delete":

if len(os.Args) < 3 {
fmt.Println("Usage: delete <id>")
return
}

id, err := strconv.Atoi(os.Args[2])

if err != nil {
fmt.Println("Invalid ID")
return
}

for i, e := range expenses {

if e.ID == id {

expenses = append(expenses[:i], expenses[i+1:]...)

saveExpenses(expenses)

fmt.Println("Expense deleted successfully")

return
}
}

fmt.Println("Expense not found")

default:
fmt.Println("Unknown command")
}
}
