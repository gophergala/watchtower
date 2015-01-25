package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"os"
	"time"

	"github.com/gophergala/watchtower/config"
	whttp "github.com/gophergala/watchtower/layers/http"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func handleUDPPacket(packet []byte) error {
	return nil
}

func handleTCPConnection(c *net.Conn) error {
	return nil
}

func main() {
	// Flags
	var cores int
	var httpPort int
	var tcpPort int
	var udpPort int
	var secretKey string

	// Parse flags
	flag.StringVar(&secretKey, "secret", "supersecret", "The secret key clients must use to register")
	flag.IntVar(&cores, "cores", 1, "Number of cores (GOMAXPROCS) Watchtower can utilize")
	flag.IntVar(&httpPort, "http", -1, "Port to host the HTTP server on (defaults to OFF)")
	flag.IntVar(&tcpPort, "tcp", -1, "Port to host the TCP server on (defaults to OFF)")
	flag.IntVar(&udpPort, "udp", -1, "Port to host the UDP server on (defaults to OFF)")
	flag.Parse()

	config.SetSecret(secretKey)

	if httpPort != -1 {
		router := httprouter.New()
		// Register as a new user
		router.Handle("POST", "/register", whttp.RegisterHandler)

		// List or join Channels
		router.Handle("GET", "/channels", whttp.ListChannelsHandler)
		router.Handle("GET", "/channels/join", whttp.JoinChannelsStreamHandler)
		router.Handle("POST", "/channels/join/async", whttp.JoinChannelsAsyncHandler)

		// Send messages
		router.Handle("POST", "/broadcast", whttp.BroadcastHandler)
		router.Handle("POST", "/send", whttp.SendMessageHandler)

		// Add Watchtower's default headers
		wrappedRouter := NewDefaultHeadersHandler(router)
		// Wrap the router with some logging middleware
		wrappedRouter = handlers.CombinedLoggingHandler(os.Stderr, wrappedRouter)

		// TODO: clean shutdown of HTTP server
		go func() {
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), wrappedRouter))
		}()
	}

	if tcpPort != -1 {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpPort))
		if err != nil {
			panic(err) // If we can't start the server, might as well panic out
		}

		go func() {
			// TODO: Clean TCP server shutdown
			for {
				conn, err := ln.Accept()
				if err != nil {
					// handle error
				}
				go handleTCPConnection(&conn)
			}
		}()
	}

	if udpPort != -1 {
		addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", udpPort))
		sock, _ := net.ListenUDP("udp", addr)

		// TODO: Clean UDP server shutdown
		go func() {
			for {
				buffer := make([]byte, 1024)
				plen, _, err := sock.ReadFromUDP(buffer)
				if err != nil {
					// Error reading what the client send - TODO: Handle this better
					panic(err)
				}

				go handleUDPPacket(buffer[0:plen])
			}
		}()
	}

	if httpPort == -1 && tcpPort == -1 && udpPort == -1 {
		fmt.Println("you must active at least one endpoint (--http [PORT], --tcp [PORT] or --udp [PORT])")
		os.Exit(-1)
	}

	// Sleep while waiting for shutdown signals
	// TODO: Handle shutdown signals cleanly...
	for {
		time.Sleep(time.Second)
	}
}
