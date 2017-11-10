package cassandra

import (
	"fmt"
	"testing"
	"time"

	"github.com/yehohanan7/glock/glock"
)

func TestGlock(t *testing.T) {
	store := NewStore(session)
	masterCh1, slaveCh1, stopCh1 := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go glock.Start("node1", time.NewTicker(1*time.Second), store, masterCh1, slaveCh1, stopCh1)
	shouldRecieveMessage(masterCh1, slaveCh1, "expected to become master", t)

	masterCh2, slaveCh2, stopCh2 := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go glock.Start("node2", time.NewTicker(1*time.Second), store, masterCh2, slaveCh2, stopCh2)
	shouldRecieveMessage(slaveCh2, masterCh1, "expected to become slave", t)

	fmt.Println("clearing...")
	stopCh1 <- struct{}{}
	stopCh2 <- struct{}{}

	if err := store.Clear(); err != nil {
		t.Error("error while clearing locks", err)
	}
}

func shouldRecieveMessage(ch, drain chan struct{}, message string, t *testing.T) {
	for {
		select {
		case <-ch:
			return
		case <-drain:
			continue
		case <-time.After(10 * time.Second):
			t.Error(message)
			return
		}
	}
}
