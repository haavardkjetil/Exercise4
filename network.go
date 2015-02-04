package network

import(
"net"
"fmt"
"bytes"
"encoding/gob"
"log"
)

const bcast = "129.241.187.255"
const udpPort = "2878" 

type Message struct {
	Type string
	Postition int
	Order int
}

func receiveMessage(transmitChannel chan Message) {
	//Initializing
	localAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort( "",udpPort))
	if err != nil {
		log.Fatal( "Failed to resolve addr for :" + udpPort, err );
	}

	recieveConnection, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatal("UDP recv connection error on " + localAddr.String(), err)
	}
	
	defer recieveConnection.Close()

	receiveBufferRaw := make( []byte, 1600 ) // standard MTU size -- no packet should be bigger
	var receiveBuffer bytes.Buffer
	messageDecoder := gob.NewDecoder(&receiveBuffer)
	//Initialization done


		for i := 0; i<100; i++ {
			_, from, err := recieveConnection.ReadFromUDP( receiveBufferRaw )
			if from.String() == recieveConnection.LocalAddr().String() {
				continue
			}
			if err != nil {
				log.Fatal("Error receiving UDP packet: " + err.Error(),err )
			}

			receiveBuffer.Write(receiveBufferRaw)
			var mssg Message
			err = messageDecoder.Decode(&mssg)
			if err != nil {
				log.Fatal("Could not decode message: ", err)
			}
			transmitChannel <- mssg 
			receiveBuffer.Reset()

		}



}