package main

import (
	"encoding/json"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
	"log"
)

// the WPC message structure
type WpcMessage struct {
	ObjId     string `json:"objId"`
	Id        string `json:"id"`
	Timestamp string `json:"ts"`
	ObjType   string `json:"type"`
	Company   string `json:"comp"`
	Message   string `json:"message"`
	Movement  string `json:"mov"`
}

// the WPC message types
const typeLookup                string = "190"
const typeReferenciaSuperEspaña string = "101"

func main() {
	//fmt.Printf("Hello %s!\n", Fullname("Gordon", "the Gopher"))
	log.Println("Starting....")
	startWatcher("/var/dabox/wpc")
	log.Println("Starded")
}

func startWatcher(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for { // buble infinito
			log.Printf("watching on %s ...", path)
			select {
			case event := <-watcher.Events:
				log.Println("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file: " + event.Name)

					if wpcMsg, err := unmarshal(event.Name); err == nil {
						process(wpcMsg)
					}
				}
			case err := <-watcher.Errors:
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func unmarshal(filename string) (ret WpcMessage, err error) {
	dat, err := ioutil.ReadFile(filename)
	wpcMsg := &WpcMessage{}
	if err != nil {
		log.Println("error reading file", err)
		return *wpcMsg, err
	}
	json.Unmarshal(dat, &wpcMsg)
	return *wpcMsg, nil
}

func process(wpcMessage WpcMessage) {
	switch  wpcMessage.ObjType {
	case typeLookup:
		log.Println("Es una [Lookup]")
	case typeReferenciaSuperEspaña:
		log.Println("Es una [Referencia Super España]")
	default:
		log.Println("Es un tipo no manejado por la plataforma")
	}
}

func print(wpcMessage WpcMessage) {
	log.Println(" __ WPC Message _______________________________________")
	//log.Printf("    wpcMessage=",          wpcMsg)
	log.Printf("    wpcMessage.ObjId=",     wpcMessage.ObjId)
	log.Printf("    wpcMessage.Id=",        wpcMessage.Id)
	log.Printf("    wpcMessage.Timestamp=", wpcMessage.Timestamp)
	log.Printf("    wpcMessage.ObjType=",   wpcMessage.ObjType)
	log.Printf("    wpcMessage.Company=",   wpcMessage.Company)
	log.Printf("    wpcMessage.Movement=",  wpcMessage.Movement)
	//log.Printf("    wpcMessage.Message=",   wpcMessage.Message)
	log.Println(" ______________________________________________________")
}
