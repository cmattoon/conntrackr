package conntrack

const ( // Column numbers
	ENTRIES = iota
	SEARCHED
	FOUND
	NEW
	INVALID
	IGNORE
	DELETE
	DELETE_LIST
	INSERT
	INSERT_FAILED
	DROP
	EARLY_DROP
	ICMP_ERROR
	EXPECT_NEW
	EXPECT_CREATE
	EXPECT_DELETE
	SEARCH_RESTART
)

// StatResult contains info from /proc/net/stat/nf_conntrack
type StatResult struct {
	Id            int    `json:"id"`             // The row/CPU number. Not part of the kernel data structure
	Entries       uint32 `json:"entries"`        // # of entries in conntrack table.
	Searched      uint32 `json:"searched"`       // # of conntrack table lookups performed.
	Found         uint32 `json:"found"`          // # of searched entries which were successful.
	New           uint32 `json:"new"`            // # of conntrack entries added which were not expected before.
	Invalid       uint32 `json:"invalid"`        // # of packets seen which can not be tracked.
	Ignore        uint32 `json:"ignore"`         // # of packets seen which are already connected to a conntrack entry
	Delete        uint32 `json:"delete"`         // # of conntrack entries which were removed.
	DeleteList    uint32 `json:"delete_list"`    // # of conntrack entries which were put to dying list.
	Insert        uint32 `json:"insert"`         // # of entries inserted into the list.
	InsertFailed  uint32 `json:"insert_failed"`  // # of entries where list insertion was attempted but failed.
	Drop          uint32 `json:"drop"`           // # of packets dropped due to conntrack failure
	EarlyDrop     uint32 `json:"early_drop"`     // # of dropped conntrack entries to make room for new ones (maxsize reached)
	IcmpError     uint32 `json:"icmp_error"`     // # of packets which couldn't be tracked due to error. Subset of invalid
	ExpectNew     uint32 `json:"expect_new"`     // # of entries added after an expectation was already present
	ExpectCreate  uint32 `json:"expect_create"`  // # of expectations added
	ExpectDelete  uint32 `json:"expect_delete"`  // # of expectations deleted
	SearchRestart uint32 `json:"search_restart"` // # of table lookups restarted due to hashtable resizes
}

type StatResultList struct {
	Items []*StatResult `json:"items"` // An array of StatResults
}

type EntryList struct {
	Items []*Entry `json:"items"` // An array of conntrack entries
}

// A line item from /proc/net/nf_conntrack
type Entry struct {
	IpProto    string      `json:"ip_proto"`
	IpProtoNum int         `json:"ip_proto_num"`
	TxProto    string      `json:"tx_proto"`
	TxProtoNum int         `json:"tx_proto_num"`
	TTL        int         `json:"ttl"`
	State      string      `json:"state"`
	Outbound   *SocketPair `json:"outbound"` // The outgoing connection (host:port, dst:port)
	Inbound    *SocketPair `json:"inbound"`  // The incoming connection/reply (dst:port, host:port)
	IsAssured  bool        `json:"is_assured"`
	Mark       int         `json:"mark"`
	Zone       int         `json:"zone"`
	Use        int         `json:"use"`
}

type Socket struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type SocketPair struct {
	Src *Socket `json:"src"`
	Dst *Socket `json:"dst"`
}
