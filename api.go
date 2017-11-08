package glock

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/yehohanan7/glock/cassandra"
	"github.com/yehohanan7/glock/glock"
)

type GlockConfig struct {
	Owner    string
	Ticker   *time.Ticker
	Store    glock.LockStore
	MasterCh chan struct{}
	SlaveCh  chan struct{}
	StopCh   chan struct{}
}

func NewCassandraStore(session *gocql.Session) glock.LockStore {
	return cassandra.NewStore(session)
}

func StartGlock(config GlockConfig) {
	glock.Start(config.Owner, config.Ticker, config.Store, config.MasterCh, config.SlaveCh, config.StopCh)
}
