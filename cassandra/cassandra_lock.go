package cassandra

import (
	"errors"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/golang/glog"
	. "github.com/yehohanan7/glock/glock"
)

const LockId string = "glock"

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
	if err := store.session.Query(`insert into lock values (?, ?)`, LockId, owner).Exec(); err != nil {
		glog.Errorf("error while acquiring lock", err)
		return Lock{}, nil
	}
	return Lock{owner}, nil
}

func (store *CassandraLockStore) RenewLock() (Lock, error) {
	return Lock{}, nil
}

func (store *CassandraLockStore) DeleteLock() error {
	return nil
}

func NewStore(session *gocql.Session) LockStore {
	return &CassandraLockStore{session}
}
