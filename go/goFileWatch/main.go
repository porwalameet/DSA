package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Create a new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path to watch.
	err = watcher.Add("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Block the main goroutine forever.
	<-make(chan bool)
}