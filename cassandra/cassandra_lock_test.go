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
	session, _ := cluster.CreateSession()
	store = NewStore(session)
}

func TestAcquireLock(t *testing.T) {
	_, err := store.AcquireLock()

	if err != nil {
		t.Error("error while acquiring lock", err)
	}
}
