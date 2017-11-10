package glock

import (
	"errors"
	"time"

	"github.com/golang/glog"
)

var LockOwnershipLost = errors.New("lock is owned by other owner")

type Lock struct {
	Owner string
}

type LockStore interface {
	AcquireLock(owner string) (Lock, error)
	GetLock() (Lock, error)
	RenewLock(owner string) (Lock, error)
	Clear() error
}

func notifyChange(states chan string, notifyCh chan string) {
	var currentState = ""
	for state := range states {
		if state != currentState && state == "master" {
			notifyCh <- state
		}
		if state != currentState && state == "slave" {
			notifyCh <- state
		}
		currentState = state
	}
}

func Start(owner string, ticker *time.Ticker, store LockStore, notify chan string, stop chan struct{}) error {
	states := make(chan string)
	go notifyChange(states, notify)

	glog.Info("starting process...")
	for {
		select {
		case <-ticker.C:
			if _, err := store.AcquireLock(owner); err == nil {
				if _, err := store.RenewLock(owner); err == nil {
					states <- "master"
				} else {
					states <- "slave"
				}
			} else {
				states <- "slave"
			}
		case <-stop:
			close(states)
			glog.Info("stoping election...")
			return nil
		}
	}
}
