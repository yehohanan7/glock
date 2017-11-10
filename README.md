# Glock [![Build Status](https://travis-ci.org/yehohanan7/glock.svg)](https://travis-ci.org/yehohanan7/glock?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/yehohanan7/glock)](https://goreportcard.com/report/github.com/yehohanan7/glock)


# Example using cassandra

```go
	import "github.com/yehohanan7/glock/glock"
  session := //initialize a cassandra session using gocql
  notifyCh, stopCh := make(chan string), make(chan struct{})
  config := glock.Config {
    "node1",
    time.NewTicker(1*time.Second),
    glock.NewCassandraStore(session),
    notifyCh,
    stopCh,
  }
	store := glock.NewCassandraStore(session)
  go glock.Start()

```

