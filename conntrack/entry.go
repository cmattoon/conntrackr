package conntrack

import (
	"strconv"
	"strings"

	"github.com/cmattoon/conntrackr/internal/net"
)

const (
	PROTO_TCP string = "tcp"
	PROTO_UDP string = "udp"
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return int(i)
}

// Returns nil, nil unless the line is a TCP or UDP entry
func NewEntryFromFields(fields []string) (*Entry, error) {
	e := &Entry{}
	e.IpProto = fields[0]
	e.IpProtoNum = atoi(fields[1])
	e.TxProto = fields[2]
	e.TxProtoNum = atoi(fields[3])
	e.TTL = atoi(fields[4])

	e.extractAddrs(fields)

	return e, nil
}

func _xf(field string) string {
	return strings.SplitN(field, "=", 1)[1]
}

// Sets src/dst pairs
// src= dst= sport= dport=
// src= dst= sport= dport=
func (e *Entry) extractAddrs(fields []string) {
	for i, f := range fields {
		if strings.HasPrefix(f, "src=") {
			src := _xf(fields[i])
			dst := _xf(fields[i+1])
			sport := _xf(fields[i+2])
			dport := _xf(fields[i+3])

			if i < 7 { // First pair is outgoing connection (host -> internet)
				e.Outbound = NewSocketPair(NewSocket(src, sport), NewSocket(dst, dport))
			} else { // Second pair is incoming connection/reply (internet -> host)
				e.Inbound = NewSocketPair(NewSocket(src, sport), NewSocket(dst, dport))
			}

		} else if strings.HasPrefix(f, "mark=") {
			e.Mark = atoi(_xf(fields[i]))
		} else if strings.HasPrefix(f, "zone=") {
			e.Zone = atoi(_xf(fields[i]))
		} else if strings.HasPrefix(f, "use=") {
			e.Use = atoi(_xf(fields[i]))
		} else if strings.Contains(f, "[ASSURED]") {
			e.IsAssured = true
		} else if net.IsTCPState(f) {
			e.State = f
		}
	}
}
