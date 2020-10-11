package proxy

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

const examplePeerRow string = "icez       BGP        ---        up     2020-10-10    Established"

var columnRegex = regexp.MustCompile(`(\w+)\s+(\w+)\s+([\w-]+)\s+(\w+)\s+([0-9\-.:]+)\s?(.*)`)

func TestParsePeerRow(t *testing.T) {
	reg := parsePeerRowRegex(examplePeerRow)
	impl := parsePeerRow(examplePeerRow)

	assert.Equal(t, reg, impl)
}

func parsePeerRowRegex(line string) PeerRow {
	split := columnRegex.FindStringSubmatch(line)

	// split[0] == whole string
	return PeerRow{
		Name:  strings.TrimSpace(split[1]),
		Proto: strings.TrimSpace(split[2]),
		Table: strings.TrimSpace(split[3]),
		State: strings.TrimSpace(split[4]),
		Since: strings.TrimSpace(split[5]),
		Info:  strings.TrimSpace(split[6]),
	}
}

func BenchmarkParsePeerRowRegex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRowRegex(examplePeerRow)
	}
}

func BenchmarkParsePeer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRow(examplePeerRow)
	}
}
