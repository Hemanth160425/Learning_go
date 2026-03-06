# GitHub Activity CLI (Day 2)

A simple Command Line Interface (CLI) application built in Go that fetches and displays the recent public activity of any GitHub user.

This project is part of a 100 Days of Code challenge to learn Go, focusing on HTTP requests and JSON parsing.

## 🚀 Features

- Fetches a user's recent GitHub activity using the public GitHub API.
- Parses the JSON response into Go structs.
- Prints a clean, readable list of their last 10 actions (PushEvent, IssuesEvent, WatchEvent, etc.).

## 📦 What I Learned
Through building this project, I learned the following core Go concepts:
- Making HTTP GET requests using `net/http`
- Reading HTTP response bodies using `io.ReadAll`
- Structuring and mapping complex JSON data using `encoding/json`
- Using `defer` to ensure network connections are closed (`defer resp.Body.Close()`)
- Formatted string building with `fmt.Sprintf`
- Command-line argument parsing and validation (`os.Args`)

## 🛠️ Usage

To run the application, navigate to the `Day_2` directory and use the `go run` command followed by `github-activity` and a GitHub username.

```bash
# General syntax
go run main.go github-activity <username>

# Example usage
go run main.go github-activity Hemanth160425
```

### Example Output
```text
Total arguments: 3
- Pushed commits to Hemanth160425/Learning_go
- CreateEvent in Hemanth160425/Learning_go
- Pushed commits to Hemanth160425/Goredis
- Starred SomeUser/SomeAwesomeRepo
```

## 📁 Project Structure

- `main.go`: The central entry point that handles terminal input, HTTP requests, and console output.
- `go.mod`: Project module definition files.
