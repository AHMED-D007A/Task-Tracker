package main

import (
	"log"
	"os"
)

func main() {
	var file *os.File

	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		file, err = os.Create("tasks.json")
		if err != nil {
			log.Println(err)
		}
		file.Write([]byte("[\n]"))
		os.Setenv("taskID", "1")
	} else {
		file, err = os.OpenFile("tasks.json", os.O_RDWR, 0644)
		if err != nil {
			log.Println(err)
		}
	}

	switch os.Args[1] {
	case "add":
	case "update":
	case "delete":
	case "mark-in-progress":
	case "mark-done":
	case "list":
	}

}
