package wpc

import (
	"encoding/json"
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

func Unmarshal(wpcMessageContent []byte) (WpcMessage, error) {
	wpcMessage   := &WpcMessage{}
	//dat, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	log.Println("error reading file", err)
	//	return *wpcMsg, err
	//}
	json.Unmarshal(wpcMessageContent, &wpcMessage)
	return *wpcMessage, nil
}
