package config

import (
	"path"
	"time"
)

// You can override these in env if you really want to
// Final path is "${ProcfsRoot}/${ENV[key]}"
var defaultFileMap = map[string]string{
	"NF_CONNTRACK_MAX":   "sys/net/netfilter/nf_conntrack_max",
	"NF_CONNTRACK_COUNT": "sys/net/netfilter/nf_conntrack_count",
	"NF_CONNTRACK_STAT":  "net/stat/nf_conntrack",
	"NF_CONNTRACK_LIST":  "net/nf_conntrack",
}

type Config struct {
	ProcfsRoot string            // FS Root for path generation
	Files      map[string]string // List of filenames -> abspath
	Interval   int               // Polling interval, in seconds
	LogFlags   int
	LogPrefix  string
}

// New returns a pointer to a new Config object
func New() *Config {
	cfg := &Config{
		ProcfsRoot: GetEnv("PROCFS_ROOT", "/proc"),
		Files:      make(map[string]string),
		Interval:   GetEnvInt("INTERVAL", 5),
		LogPrefix:  "",
		LogFlags:   0,
	}
	cfg.initFileMap()
	return cfg
}

// Allows customization of the filepaths by using env vars
func (c *Config) initFileMap() {
	for k, v := range defaultFileMap {
		c.Files[k] = path.Join(c.ProcfsRoot, GetEnv(k, v))
	}
}

// Returns a time.Duration
func (c *Config) GetInterval() time.Duration {
	return time.Duration(c.Interval) * time.Second
}
