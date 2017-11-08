package cassandra

import (
	"testing"
	"time"

	"github.com/yehohanan7/glock/glock"
)

func TestGlock(t *testing.T) {
	masterCh, slaveCh, stopCh := make(chan struct{}), make(chan struct{}), make(chan struct{})

	go glock.Start("node1", time.NewTicker(2*time.Second), NewStore(session), masterCh, slaveCh, stopCh)

	//<-masterCh
}
