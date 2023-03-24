package test

import (
	"github.com/nats-io/nats.go"
	"testing"
)

func TestNags(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	//if err != nil {
	//	t.Error("CAN NOT CONNECT TO PUBLISHER ", err)
	//}
	_ = nc.Publish("dsm", []byte("QWQ"))
	_, _ = nc.Subscribe("dsm", func(m *nats.Msg) {
		t.Error("thank you for subscribe me,", m)
	})
}
