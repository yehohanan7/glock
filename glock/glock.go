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

func retainOwnership(owner string, store LockStore) error {
	if _, err := store.RenewLock(owner); err != nil {
		return err
	}
	return nil
}

func Start(owner string, ticker *time.Ticker, store LockStore, master, slave, stop chan struct{}) error {
	for {
		select {
		case <-ticker.C:
			if lock, err := store.GetLock(); lock.Owner == owner && err == nil {
				retainOwnership(owner, store)
				continue
			}

			if lock, err := store.AcquireLock(owner); err != nil && lock.Owner == owner {
				master <- struct{}{}
				retainOwnership(owner, store)
				continue
			}

			slave <- struct{}{}
		case <-stop:
			glog.Info("stoping election...")
			return nil
		}
	}
}
