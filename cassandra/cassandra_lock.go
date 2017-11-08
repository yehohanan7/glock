package cassandra

import (
	"errors"
	"strconv"

	"github.com/gocql/gocql"
	. "github.com/yehohanan7/glock/glock"
)

const (
	LockId = "glock"
)

type CassandraLockStore struct {
	session *gocql.Session
}

func (store *CassandraLockStore) GetLock() (Lock, error) {
	var owner string
	iter := store.session.Query(`select owner from lock where id = ?`, LockId).Iter()
	if iter.NumRows() != 1 {
		return Lock{}, errors.New("error while getting the lock, rows returned: " + strconv.Itoa(iter.NumRows()))
	}
	iter.Scan(&owner)
	return Lock{owner}, nil
}

func (store *CassandraLockStore) AcquireLock(owner string) (Lock, error) {
	applied, err := store.session.Query(`insert into lock (id, owner) values(?, ?) if not exists`, LockId, owner).ScanCAS(nil, nil)
	if err != nil {
		return Lock{}, err
	}

	if !applied {
		return Lock{}, LockOwnershipLost
	}

	return Lock{owner}, nil
}

func (store *CassandraLockStore) RenewLock(owner string) (Lock, error) {
	applied, err := store.session.Query(`update lock set owner = ? where id = ? if owner = ?`, owner, LockId, owner).ScanCAS(nil)
	if err != nil {
		return Lock{}, err
	}

	if !applied {
		return Lock{}, LockOwnershipLost
	}

	return Lock{owner}, nil
}

func (store *CassandraLockStore) DeleteLock() error {
	return nil
}

func (store *CassandraLockStore) Clear() error {
	return store.session.Query(`truncate lock`).Exec()
}

func NewStore(session *gocql.Session) LockStore {
	return &CassandraLockStore{session}
}
