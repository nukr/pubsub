package pubsub

import (
	"fmt"
	"testing"
)

func TestPubSub(t *testing.T) {
	expected := "hihi"
	ps := New()
	ch1 := ps.Subscribe("ch1")
	ch2 := ps.Subscribe("ch1")
	ps.Publish("ch1", expected)
	select {
	case data := <-ch1:
		fmt.Println(data)
	default:
		t.Errorf("expected %s, but got nothing\n", expected)
	}
	select {
	case data := <-ch2:
		fmt.Println(data)
	default:
		t.Errorf("expected %s, but got nothing\n", expected)
	}
}
