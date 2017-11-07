package cassandra

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/yehohanan7/glock/glock"
)

var store glock.LockStore

func init() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "test"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	session.Query(`CREATE KEYSPACE test WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'CHX' : 3 };`).Exec()
	session.Query(`CREATE TABLE test.lock (id text, owner text, PRIMARY KEY (id)) WITH default_time_to_live = 10;`).Exec()
	store = NewStore(session)
}

func TestAcquireLock(t *testing.T) {
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

}
