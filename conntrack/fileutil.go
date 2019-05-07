package conntrack

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// For reading a simple integer from a file
func GetIntFromFile(fname string) int {
	bs, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to read %s: %s", fname, err))
	}

	i, err := strconv.Atoi(strings.TrimSpace(string(bs)))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to parse integer: %s", err))
	}
	return i
}

func GetUint32FromFile(fname string) uint32 {
	return uint32(GetIntFromFile(fname))
}
