package cassandra

import (
	"os"
	"testing"
	"time"

	"github.com/gocql/gocql"
)

var session *gocql.Session

func setupCassandra() {
	cluster := gocql.NewCluster("localhost")
	cluster.Timeout = 10 * time.Second
	sess, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	if err := sess.Query(`CREATE KEYSPACE test WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor': 1};`).Exec(); err != nil {
		panic(err)
	}
	if err := sess.Query(`CREATE TABLE test.lock (id text, owner text, PRIMARY KEY (id)) WITH default_time_to_live = 5;`).Exec(); err != nil {
		panic(err)
	}

	cluster.Keyspace = "test"
	session, _ = cluster.CreateSession()
}

func cleanupCassandra() {
	if err := session.Query(`drop keyspace test;`).Exec(); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	setupCassandra()
	retCode := m.Run()
	cleanupCassandra()
	os.Exit(retCode)
}
