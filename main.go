
package main

import(
"net"
"fmt"
//"bytes"
//"encoding/gob"
"log"
)

const bcast = "129.241.187.255"
const udpPort = "2877" 

type Message struct {
	Type string
	Postition int
	Order int
}

func receiveMessage(transmitChannel chan string) {
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

	receiveBuffer := make( []byte, 1600 ) // standard MTU size -- no packet should be bigger

//	unpacker := gob.NewDecoder(&receiveBuffer)
	//Initialization done


		for i := 0; i<100; i++ {
			len, from, err := recieveConnection.ReadFromUDP( receiveBuffer )
			if from.String() == recieveConnection.LocalAddr().String() {
				continue
			}
			if err != nil {
				log.Fatal("Error receiving UDP packet: " + err.Error(),err )
			}
			transmitChannel <- string( receiveBuffer[ :len ] )
			//		err = unpacker.Decode(&message)
		}



}


func send(){
	//
}


func main(){
	transmitChannel := make(chan string,5)
	go receiveMessage(transmitChannel)
	for {
		message := <- transmitChannel
		fmt.Println(message)
	}
}
