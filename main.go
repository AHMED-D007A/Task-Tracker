package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var file *os.File

	info, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		file, err = os.Create("tasks.json")
		if err != nil {
			log.Println(err)
		}
		file.Write([]byte("[\n]"))
	} else {
		file, err = os.OpenFile("tasks.json", os.O_RDWR, 0644)
		if err != nil {
			log.Println(err)
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
			file.Seek(-1, io.SeekEnd)
			id := rand.Intn(900000) + 10000
			stringTask := fmt.Sprintf("\t{\"id\": %v, \"description\": \"%s\", \"status\": \"%s\", \"createdAt\": \"%v\", \"updatedAt\": \"%v\"},\n]", id, string(os.Args[2]), "todo", time.DateTime, time.DateTime)
			file.WriteString(stringTask)
		case "update":
		case "delete":
		case "mark-in-progress":
		case "mark-done":
		case "list":
		}
	}
}
