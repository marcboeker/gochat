package protocol

import (
	"encoding/json"
	"time"
)

// Message contains the user, the message and a timestamp.
type Message struct {
	User      string    `json:user`
	Message   string    `json:message`
	Timestamp time.Time `json:time`
}

func (m *Message) Serialize() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func MessageFromBytes(b []byte) (*Message, error) {
	m := new(Message)
	err := json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
