package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var file *os.File

	// checks for file status.
	info, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		file, err = os.Create("tasks.json")

		if err != nil {
			log.Println(err)
			return
		}

		file.Write([]byte("[\n]"))
	} else {
		file, err = os.OpenFile("tasks.json", os.O_RDWR, 0644)

		if err != nil {
			log.Println(err)
			return
		}

		if info.Size() == 0 {
			file.Write([]byte("[\n]"))
		}
	}

	if len(os.Args) == 1 {
		fmt.Println("Expected [command] [...argument].")
	} else {
		switch os.Args[1] {
		case "add":
			addition(file)
		case "update":
			update(file)
		case "delete":
			deleteTask(file)
		case "mark-in-progress":
			inProgress(file)
		case "mark-done":
			done(file)
		case "list":
			list(file)
		}
	}
}
