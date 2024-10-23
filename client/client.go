package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	// Ambil nama pengguna saat program dijalankan
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	clientName, _ := reader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)

	// Hubungkan ke server TCP
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Kirim nama pengguna ke server
	conn.Write([]byte(clientName + "\n"))

	// Kirim pesan bahwa pengguna telah bergabung
	conn.Write([]byte(fmt.Sprintf("[%s] has joined the chat\n", clientName)))

	// Jalankan goroutine untuk membaca pesan dari server
	go readMessages(conn, clientName)

	// Fungsi untuk menangani penutupan terminal
	go handleClose(conn, clientName)

	// Kirim pesan ke server
	writeMessages(conn, clientName)
}

// Fungsi untuk membaca pesan dari server
func readMessages(conn net.Conn, clientName string) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Pisahkan antara pengirim dan isi pesan
		messageParts := strings.SplitN(message, ":", 2)
		if len(messageParts) == 2 {
			sender := strings.TrimSpace(messageParts[0])
			content := strings.TrimSpace(messageParts[1])

			// Cek apakah pengirimnya adalah pengguna itu sendiri atau orang lain
			if sender != clientName {
				fmt.Printf("[%s]: %s\n", sender, content) // Tampilkan [Name] untuk pengguna lain
			}
		} else {
			// Untuk pesan bergabung atau keluar
			fmt.Print(message)
		}
	}
}

// Fungsi untuk menulis pesan ke server
func writeMessages(conn net.Conn, clientName string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		// Tampilkan pesan yang dikirimkan oleh diri sendiri
		fmt.Printf("[You]: %s\n", message)

		// Format pesan dengan nama pengguna dan kirim ke server
		formattedMessage := fmt.Sprintf("%s: %s\n", clientName, message)
		conn.Write([]byte(formattedMessage))
	}
}

// Fungsi untuk menangani penutupan terminal
func handleClose(conn net.Conn, clientName string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Tunggu sampai sinyal diterima
	<-c
	// Kirim pesan bahwa pengguna telah meninggalkan chat
	conn.Write([]byte(fmt.Sprintf("[%s] has left the chat\n", clientName)))
	conn.Close() // Menutup koneksi
	os.Exit(0)    // Keluar dari program
}
