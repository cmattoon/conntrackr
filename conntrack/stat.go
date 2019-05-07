package conntrack

import (
	"bufio"
	"os"
	"strings"
)

// Gets info from /proc/net/stat/nf_conntrack
func Stat(fname string) (*StatResultList, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	result := NewStatResultList()

	i := 0
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		sr, err := NewStatResultFromFields(i, fields)
		if err != nil && (i > 0 && strings.Contains(err.Error(), "Probably a header")) {
			return nil, err
		}

		result.Append(sr)
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
