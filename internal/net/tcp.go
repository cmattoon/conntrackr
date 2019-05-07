package net

var statusList = [...]string{
	"CLOSED",
	"LISTEN",
	"SYN_SENT",
	"SYN_RCVD",
	"SYN_SENT",
	"ESTABLISHED",
	"CLOSE_WAIT",
	"LAST_ACK",
	"FIN_WAIT_1",
	"FIN_WAIT_2",
	"CLOSING",
	"TIME_WAIT",
}

func IsTCPState(s string) bool {
	for _, st := range statusList {
		if s == st {
			return true
		}
	}
	return false
}
