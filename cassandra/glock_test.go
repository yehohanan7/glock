package cassandra

import (
	"testing"
	"time"

	"github.com/golang/glog"
	"github.com/yehohanan7/glock/glock"
)

func TestGlock(t *testing.T) {
	store := NewStore(session)
	masterCh, slaveCh, stopCh := make(chan struct{}), make(chan struct{}), make(chan struct{})

	go glock.Start("node1", time.NewTicker(1*time.Second), store, masterCh, slaveCh, stopCh)

	select {
	case <-masterCh:
		glog.Info("master!")
	case <-slaveCh:
		t.Error("expected to become master")
	case <-time.After(10 * time.Second):
		stopCh <- struct{}{}
		t.Error("didnt become master!")
	}

	if err := store.Clear(); err != nil {
		t.Error("error while clearing locks", err)
	}
}
