package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/marcboeker/gochat/protocol"
)

var (
	broker  *Broker
	clients Clients = Clients{}
)

// Clients manages a list of all connected clients.
type Clients map[string]net.Conn

// Add adds a client's net.Conn to the list of connected clients.
// It uses the client's remote addr as identifier.
func (c *Clients) Add(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	fmt.Println("Adding client %s", addr)
	clients[addr] = conn
}

// Remove removes a client from the list of connected clients.
func (c *Clients) Remove(addr string) {
	fmt.Println("Removing client %s", addr)
	delete(*c, addr)
}

// Broker takes care of dispatching messages to connected clients.
type Broker struct{}

// Dispatch takes care of delivering incoming messages to each client.
// If a client does not respond within a given time, it is removed from the list.
func (b *Broker) Dispatch(m *protocol.Message) {
	m.Timestamp = time.Now()
	msg, err := m.Serialize()
	if err != nil {
		fmt.Errorf("Could not serialze message", err)
	}

	for addr, c := range clients {
		fmt.Println("Dispatching to", addr)
		_, err := fmt.Fprintf(c, "%s\n", msg)
		if err != nil {
			c.Close()
			clients.Remove(addr)
		}
	}
}

func serve(conn net.Conn) {
	connbuf := bufio.NewReader(conn)
	for {
		recv, err := connbuf.ReadBytes('\n')
		if err != nil {
			break
		}

		if len(recv) > 0 {
			m, err := protocol.MessageFromBytes(recv)
			if err != nil {
				fmt.Errorf("Could not parse message: ", err)
				return
			}

			fmt.Println("Received message:", m)

			broker.Dispatch(m)
		}
	}
}

// Start starts a new server and accepts connections from clients.
func Start() {
	broker = new(Broker)

	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		clients.Add(conn)
		go serve(conn)
	}
}
