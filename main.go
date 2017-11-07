package main

import (
	"flag"
	"time"

	"github.com/gocql/gocql"
	"github.com/golang/glog"
	"github.com/yehohanan7/glock/cassandra"
	"github.com/yehohanan7/glock/glock"
)

func init() {
	flag.Parse()
}

func main() {
	//cluster := gocql.NewCluster("localhost")
	//session, err := cluster.CreateSession()

	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "test"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	store := cassandra.NewStore(session)

	masterCh, slaveCh, stopCh := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go glock.Start("host1", time.NewTicker(2*time.Second), store, masterCh, slaveCh, stopCh)
	for {
		select {
		case <-masterCh:
		case <-slaveCh:
			glog.Info("received event")
		case <-time.After(5 * time.Second):
			stopCh <- struct{}{}
		}
	}
}
