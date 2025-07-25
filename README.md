# TCP Chat Application in Go

A simple real-time chat application using TCP sockets, built in Go (Golang). This application allows multiple clients to connect to a server and exchange messages with each other in real time via the terminal.

## Features

- Real-time message broadcasting
- Custom client name registration
- Graceful disconnect handling
- Console-based UI

## Architecture

This application consists of two components:

1. **Server (`server.go`)**  
   Listens on port `8080` and handles multiple client connections. Broadcasts incoming messages from any client to all other connected clients.

2. **Client (`client.go`)**  
   Connects to the server, registers with a name, sends and receives messages, and handles graceful shutdown on Ctrl+C.

## How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-tcp-chat.git
cd go-tcp-chat
```

### 2. Run the Server

```bash
go run server.go
```

Server will start on `localhost:8080`.

### 3. Run the Client (in another terminal)

```bash
go run client.go
```

You'll be prompted to enter your name. After that, you can start chatting.

> 💡 You can run multiple clients in different terminal windows to simulate multiple users.

## Example

```bash
$ go run client.go
Enter your name: Alice
[You]: Hello everyone!
```

```bash
$ go run client.go
Enter your name: Bob
[Alice]: Hello everyone!
[You]: Hi Alice!
```

## Dependencies

- Go 1.18 or newer
- No external packages required (pure Go standard library)

## License

This project is licensed under the MIT License.
