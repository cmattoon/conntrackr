package cmd

import (
	"time"

	CT "github.com/cmattoon/conntrackr/conntrack"
)

type ScanResult struct {
	Timestamp time.Time          `json:"timestamp"` // Time of sample
	Totalconn uint32             `json:"totalconn"` // Total connections per /proc/sys/net/nf_conntrack_count
	Maxconn   uint32             `json:"maxconn"`   // Max connections per /proc/sys/netfilteer/nf_conntrack_max
	Stat      *CT.StatResultList `json:"stat"`      // Per-CPU StatResults
	Entries   *CT.EntryList      `json:"entries"`   // Info about connections from /proc/net/nf_conntrack
}

func NewScanResult(tc uint32, mc uint32, sr *CT.StatResultList, el *CT.EntryList) *ScanResult {
	return &ScanResult{
		Timestamp: time.Now(),
		Totalconn: tc,
		Maxconn:   mc,
		Stat:      sr, //NewStatResultList(),
		Entries:   el, //NewEntryList(),
	}
}
