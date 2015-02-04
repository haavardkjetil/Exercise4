
package main


import(
"Network"
"time"
)


func main() {
	transmitChannel := make(chan Network.Packet,5)
	go Network.Send(transmitChannel)
	for i := 0; i < 5; i++ {
		transmitChannel <- Network.Packet{"This is struct number ", i, 0}
		time.Sleep(100*time.Millisecond)
	}
	transmitChannel <- Network.Packet{"Terminate", 0, 0}
	time.Sleep(100*time.Millisecond)
}