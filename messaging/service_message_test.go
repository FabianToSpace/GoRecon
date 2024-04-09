package messaging_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/FabianToSpace/GoRecon/messaging"
	"github.com/FabianToSpace/GoRecon/plugins"
)

func TestSerialize(t *testing.T) {
	message := &messaging.ServiceMessage{Target: "localhost",
		Service: plugins.Service{
			Target:   "Target",
			Protocol: "tcp",
			Port:     80,
		},
	}

	serialized, err := message.Serialize()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []byte("{\"Target\":\"localhost\",\"Service\":{\"Target\":\"Target\",\"Protocol\":\"tcp\",\"Port\":80,\"Name\":\"\",\"Secure\":false,\"Version\":\"\",\"Scheme\":\"\"}}\n")
	if !bytes.Equal(serialized, expected) {
		t.Errorf("Serialization incorrect. Got: %s, Expected: %s", serialized, expected)
	}
}

func TestDeserialize(t *testing.T) {
	jsonBytes := []byte("{\"Target\":\"localhost\",\"Service\":{\"Target\":\"Target\",\"Protocol\":\"tcp\",\"Port\":80,\"Name\":\"\",\"Secure\":false,\"Version\":\"\",\"Scheme\":\"\"}}\n")

	deserializedMessage := messaging.ServiceMessage{}
	expected := &messaging.ServiceMessage{Target: "localhost",
		Service: plugins.Service{
			Target:   "Target",
			Protocol: "tcp",
			Port:     80,
		},
	}

	deserialized, err := deserializedMessage.Deserialize(jsonBytes)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if reflect.DeepEqual(expected, deserialized) {
		t.Errorf("Serialization incorrect.")
	}
}
