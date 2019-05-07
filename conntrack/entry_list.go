package conntrack

import (
	"bufio"
	"os"
	"strings"
)

// Return information about all connections in the table
func GetConnections(filename string) (*EntryList, error) {
	// Not working yet
	el := NewEntryList()
	return el, nil

	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		e, err := NewEntryFromFields(fields)
		if err != nil {
			return nil, err
		}
		if e != nil {
			el.Append(e)
		}
	}
	return el, nil
}

func NewEntryList() *EntryList {
	return &EntryList{
		Items: make([]*Entry, 0),
	}
}

func (l *EntryList) Append(e *Entry) {
	l.Items = append(l.Items, e)
}
