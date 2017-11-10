# Glock [![Build Status](https://travis-ci.org/yehohanan7/glock.svg)](https://travis-ci.org/yehohanan7/glock?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/yehohanan7/glock)](https://goreportcard.com/report/github.com/yehohanan7/glock)


## Introduction
Glock is used to achieve consensus between different hosts using various persistent stores.


### Cassandra
The cassandra implementation is based on the article here: https://www.datastax.com/dev/blog/consensus-on-cassandra

#### Setup
```bash
CREATE KEYSPACE test WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor': 1};
CREATE TABLE test.lock (id text, owner text, PRIMARY KEY (id)) WITH default_time_to_live = 5;
```

#### Glock example
```go
	import "github.com/yehohanan7/glock/glock"
  session := //initialize a cassandra session using gocql

	config := glock.Config{
		"host1",
		time.NewTicker(2 * time.Second),
		glock.NewCassandraStore(session),
		notifyCh,
		stopCh,
	}

	go glock.Start(config)

	for {
		select {
		case state := <-notifyCh:
			if state == "master" {
				fmt.Println("became master")
			}

			if state == "slave" {
				fmt.Println("became slave")
			}
		}
	}
```

