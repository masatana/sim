package main

import (
	"flag"
	"log"
	"net"
	"time"
)

func doServer() {
	log.Println("Server is running at 0.0.0.0:8888")
	conn, err := net.ListenPacket("udp", "0.0.0.0:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))

	}
}

func doClient(numProcs int) {
	ch := make(chan int, numProcs)
	// wg := sync.WaitGroup{}
	// wg.Add(10)
	for {
		ch <- 1
		go func() {
			laddr, err := net.ResolveUDPAddr("udp4", "localhost:60000")
			raddr := net.UDPAddr{IP: net.ParseIP("localhost"), Port: 8888}
			conn, err := net.DialUDP("udp4", laddr, &raddr)
			// conn, err := net.Dial("udp4", "localhost:8888")
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			log.Println("Sending to server")
			_, err = conn.Write([]byte("Hello from Client"))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(1 * time.Second)
			<-ch
			// wg.Done()
		}()
	}
	//wg.Wait()
}

func main() {
	var serverMode = flag.Bool("serverMode", false, "run in server mode")
	var numProcs = flag.Int("numProcs", 2, "number of goroutines")
	flag.Parse()
	if *serverMode {
		doServer()
	} else {
		doClient(*numProcs)
	}
}
