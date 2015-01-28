
package networkCommunication

import(
"net"
"fmt"
"bytes"
"encoding/gob"
"log"
)

type Message struct {
	Type string
	Postition int
	Order int
}

func receive(message *Message) {
	var network bytes.Buffer
	unpacker := gob.NewDecoder(&network)

	err = unpacker.Decode(message)
	if err != nil {
        log.Fatal("decode error:", err)  //TODO: Fault acceptance
    }
}

func send(){
	//
}