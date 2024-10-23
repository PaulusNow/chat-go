package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var clients []net.Conn
var clientNames = make(map[net.Conn]string)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server running on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		clients = append(clients, conn)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Ambil nama client pertama kali saat mereka terhubung
	clientName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading client name:", err)
		return
	}
	clientName = strings.TrimSpace(clientName)
	clientNames[conn] = clientName

	// Tampilkan pesan "[Name] has joined the chat" hanya di server
	fmt.Printf("[%s] has joined the chat\n", clientName)

	// Baca pesan dari klien dan broadcast ke semua klien lain
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("[%s] has left the chat.\n", clientName)
			return
		}
		broadcastMessage(message, conn)
	}
}

func broadcastMessage(message string, sender net.Conn) {
	for _, client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message))
			if err != nil {
				fmt.Println("Error broadcasting message:", err)
				return
			}
		}
	}
}
