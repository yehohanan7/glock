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
	DeleteLock() error
	Clear() error
}

func Start(host string, ticker *time.Ticker, lockStore LockStore, master, slave, stop chan struct{}) error {
	for {
		select {
		case <-ticker.C:
			glog.Info("running election")
		case <-stop:
			glog.Info("stoping election...")
			return nil
		}
	}
}
