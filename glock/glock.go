package glock

import (
	"time"

	"github.com/golang/glog"
)

type Lock struct {
	Owner string
}

type LockStore interface {
	AcquireLock(host string) (Lock, error)
	GetLock() (Lock, error)
	RenewLock() (Lock, error)
	DeleteLock() error
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
