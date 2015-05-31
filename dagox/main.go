package main

import (
	"fmt"
	"encoding/json"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
	"log"
	"flag"
)

// the WPC message structure
type WpcMessage struct {
	ObjId        string         `json:"objId"`
	Id           string         `json:"id"`
	Timestamp    string         `json:"ts"`
	ObjType      string         `json:"type"`
	Company      string         `json:"comp"`
	Message      string         `json:"message"`
	Movement     string         `json:"mov"`
}

type WpcMessageContent struct {
	IdExt        string         `json:"id_Ext"`
	Id           string         `json:"id"`
	ContainerId  string         `json:"entryContainerID"`
	Movement     string         `json:"movement"`
	CoreAttribs  []Attrib       `json:"coreAttribs"`
}

type Attrib struct {
	Id           string         `json:"id"`
	Type         string         `json:"type"`
	Value        string         `json:"value"`
	Children     []Attrib 	    `json:"children"`
}

// a lookup structure
//    about JSON mapping: https://eager.io/blog/go-and-json/    OK!
//                        http://stackoverflow.com/questions/23798283/golang-unmarshal-json     OK!
//                        http://stackoverflow.com/questions/21268000/unmarshaling-nested-json-objects-in-golang          
type Lookup struct {
	IdExt        string            `json:"id_Ext"`
	ContainerId  string            `json:"entryContainerID"`
	Id           string            `json:"id"`
	LookupType   string            `json:"coreAttribs[0].id"`
//	Description  string            `json:"id_Ext"`
//	Descriptions map[string]string `json:"id_Ext"`
}

// the WPC message types
const typeLookup                string = "190"
const typeReferenciaSuperEspaña string = "101"

func main() {
	printTitle()

	// process command line arguments
    inputArg := flag.String("input", "/var/tmp", "WPC input files directory. Default is /var/tmp")
    helpArg  := flag.Bool("help", false, "show help")
	flag.Parse()
	if *helpArg {
		printHelp();
		return
	}

	//fmt.Printf("Hello %s!\n", Fullname("Gordon", "the Gopher"))
	log.Println("Starting....")
	startWatcher(*inputArg /*"/var/dabox/wpc"*/)
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

func process(wpcMessage WpcMessage) {
	switch  wpcMessage.ObjType {
	case typeLookup:
		log.Println("Es una [Lookup] :" + wpcMessage.Message)
		messageContent := wpcMessage.Message
		wpcMessageContent, err := unmarshalMessageContent(messageContent)
		if err == nil {
			log.Printf("processed messageContent.idExt: " + wpcMessageContent.IdExt)
		}
	case typeReferenciaSuperEspaña:
		log.Println("Es una [Referencia Super España]")
	default:
		log.Println("Es un tipo no manejado por la plataforma")
	}
}

func unmarshal(filename string) (ret WpcMessage, err error) {
	wpcMsg   := &WpcMessage{}
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("error reading file", err)
		return *wpcMsg, err
	}
	json.Unmarshal(dat, &wpcMsg)
	return *wpcMsg, nil
}

//func unmarshalLookup(wpcMessage WpcMessage) (Lookup, error) {
//	lookup := &Lookup{}
//	json.Unmarshal([]byte(wpcMessage.Message), &lookup)
//	return *lookup,nil
//}

func unmarshalMessageContent(messageContent string) {
	content := &WpcMessageContent{}
	json.Unmarshal([]byte(messageContent), &content)
	log.Println("msg content " + content.IdExt)
}

func printTitle() {
	fmt.Println("                         ,---,                                   ");
	fmt.Println("                      ,`--.' |                                   ");
	fmt.Println("    ,---,             |   :  :    ,---,.                         ");
	fmt.Println("  .'  .' `\\           |   |  '  ,'  .'  \\                      ");
	fmt.Println(",---.'     \\          '   :  |,---.' .' |   ,---.               ");
	fmt.Println("|   |  .`\\  |         ;   |.' |   |  |: |  '   ,'\\ ,--,  ,--,  ");                                                               
	fmt.Println(":   : |  '  |  ,--.--.'---'   :   :  :  / /   /   ||'. \\/ .`|   ");
	fmt.Println("|   ' '  ;  : /       \\       :   |    ; .   ; ,. :'  \\/  / ;  ");
	fmt.Println("'   | ;  .  |.--.  .-. |      |   :     '    | |: : \\  \\.' /   ");
	fmt.Println("|   | :  |  ' \\__\\/: . .      |   |   . |'   | .; :  \\  ;  ;  ");
	fmt.Println("'   : | /  ;  ,\" .--.; |      '   :  '; ||   :    | / \\  \\  \\");
	fmt.Println("|   | '` ,/  /  /  ,.  |      |   |  | ;  \\   \\  /./__;   ;  \\");
	fmt.Println(";   :  .'   ;  :   .'   \\     |   :   /    `----' |   :/\\  \\ ;");
	fmt.Println("|   ,.'     |  ,     .-./     |   | ,'            `---'  `--`    ");
	fmt.Println("'---'        `--`---'         `----'                        \n\n");
	fmt.Println("Da'box: The WPC - New CatalogBrowsing API integration tool.\n")
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("    dagox [-input=<path_to_repo>] [-help]\n")
	fmt.Println("where:")
	fmt.Println("    -input=<path_to_repo>: Set the input WPC files directory. Default value is /var/tmp.")
	fmt.Println("                           The application mus have read permissions on the directory.")
	fmt.Println("    -help: Show this message.")
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


