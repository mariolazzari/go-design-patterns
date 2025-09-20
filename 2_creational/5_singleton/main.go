package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Database struct {
}

var database *Database

func getInstance() *Database {
	if database == nil {
		lock.Lock()
		defer lock.Unlock()
		if database == nil {
			fmt.Println("Creating a single database")
			database = &Database{}
		} else {
			fmt.Println("Database has already been created")
		}
	} else {
		fmt.Println("Database has already been created")
	}
	return database

}

func main() {
	for range 5 {
		getInstance()
	}
}
