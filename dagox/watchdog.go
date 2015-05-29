// Copyright 2009 José Moreno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Dabox GO implementation. 
// Now contains filesystem watchdog.
//  
// https://github.com/go-fsnotify/fsnotify/blob/master/example_test.go
// https://gopkg.in/fsnotify.v1
package main

 import (
	"log"
 	"gopkg.in/fsnotify.v1"
)

func filesystemListener() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)	
	}
	defer  watcher.Close()

	done := make(chan bool)
	go func() {
		for {	// buble infinito
			select {
			case event := <- watcher.Events:
				log.Println("event: ", event)
				if event.Op & fsnotify.Write == fsnotify.Write {
					log.Println("Modified file: " + event.Name)
				}
			case err := <- watcher.Errors:
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.Add("/var/dabox/wpc")
	if err != nil {
		log.Fatal(err)
	}
	<- done
}
