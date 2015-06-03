package client

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/marcboeker/gochat/protocol"
)

var (
	conn net.Conn
	err  error
)

func listen() {
	connbuf := bufio.NewReader(conn)
	for {
		recv, err := connbuf.ReadBytes('\n')
		if err != nil {
			break
		}

		if len(recv) > 0 {
			m, err := protocol.MessageFromBytes(recv)
			if err != nil {
				fmt.Errorf("Could not parse incoming message", err)
				continue
			}

			ts := m.Timestamp.Format("2006-01-02/15:04:05")
			fmt.Printf("[%s - %s] %s\n", m.User, ts, m.Message)
		}
	}
}

func send(m *protocol.Message) error {
	b, err := m.Serialize()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(conn, "%s\n", b)
	return err
}

// Start starts a client session with the given username.
func Start(username string) {
	conn, err = net.Dial("tcp", ":1337")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Println("You are now connected! Go ahead, type something and hit enter.")

	go listen()

	for {
		var message string
		_, err := fmt.Scanln(&message)
		if err == nil {
			m := &protocol.Message{User: username, Message: message}
			err = send(m)
			if err != nil {
				fmt.Errorf("Could not send message", err)
			}
		}
	}
}
