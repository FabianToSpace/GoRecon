package plugins

import (
	"bytes"
	"encoding/json"
)

type ServiceMessage struct {
	Target  string
	Service Service
}

func (m *ServiceMessage) Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(m)
	return b.Bytes(), err
}

func (m *ServiceMessage) Deserialize(data []byte) (ServiceMessage, error) {
	var msg ServiceMessage
	decoder := json.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&msg)
	return msg, err
}
