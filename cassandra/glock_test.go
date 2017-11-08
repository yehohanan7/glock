package cassandra

import (
	"testing"
	"time"

	"github.com/golang/glog"
	"github.com/yehohanan7/glock/glock"
)

func TestGlock(t *testing.T) {
	masterCh, slaveCh, stopCh := make(chan struct{}), make(chan struct{}), make(chan struct{})

	go glock.Start("node1", time.NewTicker(2*time.Second), NewStore(session), masterCh, slaveCh, stopCh)

	select {
	case <-masterCh:
		glog.Info("master!")
	case <-slaveCh:
		t.Error("expected to become master")
	case <-time.After(10 * time.Second):
		t.Error("didnt become master!")
	}
}
