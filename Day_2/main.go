package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: task-cli <command>")
		return
	}
	command := os.Args[1]
	if command != "github-activity" {
		fmt.Println(" We can only serve as we want")
		return
	}
	userName := os.Args[2]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", userName)
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("the Url was wrong ")
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		fmt.Println("the user was not found ", err)
		return
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("the body was not working ")
		return
	}
	
	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Println("the json was not working ")
		return
	}
	
	for _, e := range events {
		switch e.Type {
		case "PushEvent":
			fmt.Printf("- Pushed commits to %s\n", e.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Opened an issue in %s\n", e.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", e.Repo.Name)
		default:
			fmt.Printf("- %s in %s\n", e.Type, e.Repo.Name)
		}
	}
}
