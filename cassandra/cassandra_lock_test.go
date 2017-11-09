package cassandra

import (
	"fmt"
	"testing"

	"github.com/yehohanan7/glock/glock"
)

func TestAcquireNewLock(t *testing.T) {
	store := NewStore(session)
	acquiredLock, err := store.AcquireLock("some-host")

	if err != nil {
		t.Error("error while acquiring lock", err)
	}

	lock, err := store.GetLock()

	if err != nil {
		t.Error("error while fetching lock", err)
	}

	if lock.Owner != acquiredLock.Owner {
		t.Errorf("expected owner %v, got %v", acquiredLock.Owner, lock.Owner)
	}

	if err := store.Clear(); err != nil {
		t.Error("error while clearing locks", err)
	}
}

func TestAcquireExistingLock(t *testing.T) {
	store := NewStore(session)
	_, err := store.AcquireLock("some-host")
	if err != nil {
		t.Error("should acquire a new lock", err)
	}

	if _, err := store.AcquireLock("some-host"); err != glock.LockOwnershipLost {
		t.Error("should not acquire new lock when the lock exists", err)
	}

	if err := store.Clear(); err != nil {
		t.Error("error while clearing locks", err)
	}

}

func TestRenewLock(t *testing.T) {
	owner := "ahost"
	store := NewStore(session)
	store.AcquireLock(owner)

	if lock, err := store.RenewLock(owner); err != nil || lock.Owner != owner {
		t.Errorf("should renew lock:%v, err: %v", lock, err)
	}

	if lock, err := store.RenewLock("different-host"); err == nil || lock.Owner != "" {
		fmt.Println(lock)
		t.Error("should not renew lock acquired by different owner")
	}

	if err := store.Clear(); err != nil {
		t.Error("error while clearing locks", err)
	}
}

func TestClearLocks(t *testing.T) {
	store := NewStore(session)
	if _, err := store.AcquireLock("host1"); err != nil {
		t.Error("lock not acquired")
	}

	if err := store.Clear(); err != nil {
		t.Error("locks not cleared", err)
	}

	if lock, _ := store.GetLock(); lock.Owner != "" {
		t.Error("locks not cleared: ", lock.Owner)
	}
}
