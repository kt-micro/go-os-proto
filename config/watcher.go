package config

import (
	"time"

	proto "github.com/kt-micro/config-srv-proto/proto/config"
)

type sourceWatcher struct {
	w proto.Config_WatchService
}

func (w *sourceWatcher) Next() (*ChangeSet, error) {
	c, err := w.w.Recv()
	if err != nil {
		return nil, err
	}
	return &ChangeSet{
		Timestamp: time.Unix(c.ChangeSet.Timestamp, 0),
		Data:      []byte(c.ChangeSet.Data),
		Checksum:  c.ChangeSet.Checksum,
		Source:    c.ChangeSet.Source,
	}, nil
}

func (w *sourceWatcher) Stop() error {
	return w.w.Close()
}
