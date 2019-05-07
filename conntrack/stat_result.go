package conntrack

import (
	"errors"
	"fmt"
	"reflect"

	tc "github.com/cmattoon/conntrackr/internal/typeconv"
)

func NewStatResultFromFields(id int, fields []string) (*StatResult, error) {
	if len(fields) != 17 { // Quick sanity check
		return nil, errors.New(fmt.Sprintf("Don't know what to do with %d fields", len(fields)))
	}
	if fields[0] == "entries" { // It's a header field
		return nil, errors.New("Probably a header field, but definitely not a value field")
	}
	return &StatResult{
		Id:            id,
		Entries:       tc.Hex2uint32(fields[ENTRIES]),
		Searched:      tc.Hex2uint32(fields[SEARCHED]),
		Found:         tc.Hex2uint32(fields[FOUND]),
		New:           tc.Hex2uint32(fields[NEW]),
		Invalid:       tc.Hex2uint32(fields[INVALID]),
		Ignore:        tc.Hex2uint32(fields[IGNORE]),
		InsertFailed:  tc.Hex2uint32(fields[INSERT_FAILED]),
		Drop:          tc.Hex2uint32(fields[DROP]),
		EarlyDrop:     tc.Hex2uint32(fields[EARLY_DROP]),
		IcmpError:     tc.Hex2uint32(fields[ICMP_ERROR]),
		ExpectNew:     tc.Hex2uint32(fields[EXPECT_NEW]),
		ExpectCreate:  tc.Hex2uint32(fields[EXPECT_CREATE]),
		ExpectDelete:  tc.Hex2uint32(fields[EXPECT_DELETE]),
		SearchRestart: tc.Hex2uint32(fields[SEARCH_RESTART]),
	}, nil
}

func (c *StatResult) PPrint() {
	for _, txt := range c.pprint() {
		fmt.Println(txt)
	}
}

func (c *StatResult) pprint() (lines []string) {
	val := reflect.ValueOf(c).Elem()

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		t := val.Type().Field(i)
		txt := fmt.Sprintf("%s: %d", t.Name, f.Interface())
		if t.Name == "Timestamp" {
			txt = fmt.Sprintf("%s: %v", t.Name, f.Interface())
		}
		lines = append(lines, txt)
	}
	return
}
