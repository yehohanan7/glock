package cassandra

import (
	"github.com/gocql/gocql"
	. "github.com/yehohanan7/glock/glock"
)

type CassandraLockStore struct {
	session *gocql.Session
}

func (store *CassandraLockStore) AcquireLock() (Lock, error) {
	return Lock{}, nil
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
