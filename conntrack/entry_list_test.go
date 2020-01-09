package conntrack

import (
	"reflect"
	"testing"
)

func TestGetConnections(t *testing.T) {
	e, err := GetConnections("./test-data/conntrack-entry")
	if err != nil {
		t.Error(err)
	}
	if len(e.Items) != 22 {
		t.Errorf("invalid length")
	}
	//ipv4     2 tcp      6 86378 ESTABLISHED src=172.25.111.120 dst=172.25.111.120 sport=34770 dport=2379 src=172.25.111.120 dst=172.25.111.120 sport=2379 dport=34770 [ASSURED] mark=0 zone=0 use=2
	exportFirstItem := Entry{
		IpProto:    "ipv4",
		IpProtoNum: 2,
		TxProto:    "tcp",
		TxProtoNum: 6,
		TTL:        86378,
		State:      "ESTABLISHED",
		Outbound: &SocketPair{
			Src: &Socket{
				Addr: "172.25.111.120",
				Port: "34770",
			},
			Dst: &Socket{
				Addr: "172.25.111.120",
				Port: "2379",
			},
		},
		Inbound: &SocketPair{
			Src: &Socket{
				Addr: "172.25.111.120",
				Port: "2379",
			},
			Dst: &Socket{
				Addr: "172.25.111.120",
				Port: "34770",
			},
		},
		IsAssured: true,
		Mark:      0,
		Zone:      0,
		Use:       2,
	}

	if !reflect.DeepEqual(exportFirstItem, *e.Items[0]) {
		t.Errorf("first item not expected result")
	}
}
