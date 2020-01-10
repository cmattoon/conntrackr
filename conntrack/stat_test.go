package conntrack

import (
	"testing"
)

func TestStat(t *testing.T) {
	e, err := Stat("./test-data/conntrack-stat")
	if err != nil {
		t.Error(err)
	}
	if len(e.Items) != 16 {
		t.Errorf("invalid length")
	}
	if e.Items[0].Delete != 20 {
		t.Errorf("invalid number")
	}

}
