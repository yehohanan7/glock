package main

import (
	"github.com/gocql/gocql"
	"github.com/yehohanan7/glock/cassandra"
	. "github.com/yehohanan7/glock/glock"
)

func NewCassandraStore(session *gocql.Session) LockStore {
	return cassandra.NewStore(session)
}
