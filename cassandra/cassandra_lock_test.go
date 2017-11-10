package cassandra

import (
	"fmt"
	"testing"
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
	owner := "some-owner"
	store := NewStore(session)
	_, err := store.AcquireLock(owner)
	if err != nil {
		t.Error("should acquire a new lock", err)
	}

	if _, err := store.AcquireLock(owner); err != nil {
		t.Error("should acquire new lock when the lock held by the same owner", err)
	}

	if _, err := store.AcquireLock("different-owner"); err == nil {
		t.Error("should not acquire lock held by a different owner", err)
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
