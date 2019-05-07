package conntrack

import (
	"fmt"
)

func NewStatResultList() *StatResultList {
	return &StatResultList{
		Items: make([]*StatResult, 0),
	}
}

// Creates and appends a StatResult from a row entry
func (l *StatResultList) Append(r *StatResult) {
	l.Items = append(l.Items, r)
}

func (l *StatResultList) PPrint() {
	for i, rs := range l.Items {
		fmt.Printf("Line %d\n", i)
		for _, line := range rs.pprint() {
			fmt.Println(fmt.Sprintf("    %s", line))
		}
	}
}
