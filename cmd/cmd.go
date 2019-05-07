package cmd

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/cmattoon/conntrackr/config"
	CT "github.com/cmattoon/conntrackr/conntrack"
)

type Command interface {
	Run() (int, error)
}

type command struct {
	config  *config.Config
	results []*ScanResult
	logger  *log.Logger
}

func NewFromConfig(cfg *config.Config) *command {
	logger := log.New(os.Stdout, cfg.LogPrefix, cfg.LogFlags)
	return &command{
		config:  cfg,
		results: make([]*ScanResult, 0),
		logger:  logger,
	}
}

// Main loop
func (c *command) Run() {
	interval := c.config.GetInterval()

	for _ = range time.NewTicker(interval).C {
		err := c.runOnce()
		if err != nil {
			log.Println(err)
		}
	}
}

// Run once per c.config.Interval
func (c *command) runOnce() error {
	tc := CT.GetUint32FromFile(c.config.Files["NF_CONNTRACK_COUNT"])
	mc := CT.GetUint32FromFile(c.config.Files["NF_CONNTRACK_MAX"])

	sr, err := CT.Stat(c.config.Files["NF_CONNTRACK_STAT"])
	if err != nil {
		return err
	}

	el, err := CT.GetConnections(c.config.Files["NF_CONNTRACK_LIST"])
	if err != nil {
		return err
	}

	scan := NewScanResult(tc, mc, sr, el)

	strval, _ := json.Marshal(scan)
	c.logger.Println(string(strval))

	return nil
}
