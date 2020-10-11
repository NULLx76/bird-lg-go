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

func parsePeerRowNewArray(line string) PeerRow {
	split := strings.Split(line, " ")

	b := make([]string, 6)
	i := 0
	for _, x := range split {
		if x != "" {
			b[i] = x
			i++
		}
	}

	return PeerRow{
		Name:  split[0],
		Proto: split[1],
		Table: split[2],
		State: split[3],
		Since: split[4],
		Info:  split[5],
	}
}

func parsePeerRowNoAllocA(line string) PeerRow {
	split := strings.Split(line, " ")

	i := 0
	for index := range split {
		if split[index] == "" {
			continue
		}

		split[i] = split[index]
		i++
	}
	split = split[:6]

	return PeerRow{
		Name:  split[0],
		Proto: split[1],
		Table: split[2],
		State: split[3],
		Since: split[4],
		Info:  split[5],
	}
}

func parsePeerRowNoAllocANoResize(line string) PeerRow {
	split := strings.Split(line, " ")

	i := 0
	for index := range split {
		if split[index] == "" {
			continue
		}

		split[i] = split[index]
		i++
	}

	return PeerRow{
		Name:  split[0],
		Proto: split[1],
		Table: split[2],
		State: split[3],
		Since: split[4],
		Info:  split[5],
	}
}

func parsePeerRowNoAllocB(line string) PeerRow {
	//split := columnRegex.FindStringSubmatch(line)
	split := strings.Split(line, " ")

	b := split[:0]
	for _, x := range split {
		if x != "" {
			b = append(b, x)
		}
	}

	return PeerRow{
		Name:  b[0],
		Proto: b[1],
		Table: b[2],
		State: b[3],
		Since: b[4],
		Info:  b[5],
	}
}

func parsePeerRowNoAllocC(line string) PeerRow {
	split := strings.Split(line, " ")

	b := split[:0]
	for i := range split {
		if split[i] != "" {
			b = append(b, split[i])
		}
	}

	return PeerRow{
		Name:  b[0],
		Proto: b[1],
		Table: b[2],
		State: b[3],
		Since: b[4],
		Info:  b[5],
	}
}

func BenchmarkParsePeerRowRegex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRowRegex(examplePeerRow)
	}
}

func BenchmarkParsePeerRowNoAllocA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRowNoAllocA(examplePeerRow)
	}
}

func BenchmarkParsePeerRowNoAllocANoResize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRowNoAllocANoResize(examplePeerRow)
	}
}

func BenchmarkParsePeerRowNoAllocB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRowNoAllocB(examplePeerRow)
	}
}

func BenchmarkParsePeerRowNoAllocC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRowNoAllocC(examplePeerRow)
	}
}

func BenchmarkParsePeerRowNewArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRowNewArray(examplePeerRow)
	}
}

func BenchmarkParsePeer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsePeerRow(examplePeerRow)
	}
}
