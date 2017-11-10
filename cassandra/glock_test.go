package cassandra

import (
	"testing"
	"time"

	"github.com/yehohanan7/glock/glock"
)

func TestGlock(t *testing.T) {
	store := NewStore(session)
	states1, stopCh1 := make(chan string), make(chan struct{})
	go glock.Start("node1", time.NewTicker(1*time.Second), store, states1, stopCh1)
	shouldRecieveMessage(states1, "master", "expected to become master", t)

	states2, stopCh2 := make(chan string), make(chan struct{})
	go glock.Start("node2", time.NewTicker(1*time.Second), store, states2, stopCh2)
	shouldRecieveMessage(states2, "slave", "expected to become slave", t)
	stopCh1 <- struct{}{}

	shouldRecieveMessage(states2, "master", "expected to become slave", t)

	stopCh2 <- struct{}{}
	if err := store.Clear(); err != nil {
		t.Error("error while clearing locks", err)
	}
}

func shouldRecieveMessage(ch chan string, expected, errorMsg string, t *testing.T) {
	for {
		select {
		case msg := <-ch:
			if msg != expected {
				continue
			}
			return
		case <-time.After(10 * time.Second):
			t.Error(errorMsg)
			return
		}
	}
}
