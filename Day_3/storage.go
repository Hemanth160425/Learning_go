package main

import (
"encoding/json"
"os"
)

const fileName = "expenses.json"

func loadExpenses() ([]Expense, error) {

file, err := os.ReadFile(fileName)

if err != nil {

if os.IsNotExist(err) {
return []Expense{}, nil
}

return nil, err
}

var expenses []Expense

err = json.Unmarshal(file, &expenses)

return expenses, err
}

func saveExpenses(expenses []Expense) error {

data, err := json.MarshalIndent(expenses, "", " ")

if err != nil {
return err
}

return os.WriteFile(fileName, data, 0644)
}
