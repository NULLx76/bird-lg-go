package proxy

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

const examplePeerRow string = "icez       BGP        ---        up     2020-10-10    Established"
const exampleProtocolHeader string = "Name       Proto      Table      State  Since         Info"
const exampleProtocolTable string = `Name       Proto      Table      State  Since         Info
device1    Device     ---        up     2020-10-10
static1    Static     dn42_roa   up     2020-10-10
static2    Static     dn42_roa_v6 up     2020-10-10
kernel1    Kernel     master6    up     2020-10-10
kernel2    Kernel     master4    up     2020-10-10
static3    Static     master4    up     2020-10-10
static4    Static     master6    up     2020-10-10
icez       BGP        ---        up     2020-10-10    Established
kioubit    BGP        ---        up     2020-10-10    Established
kioubit_v6 BGP        ---        up     2020-10-10    Established
tech9_v6   BGP        ---        up     2020-10-10    Established
tech9      BGP        ---        up     2020-10-10    Established
`

var columnRegex = regexp.MustCompile(`(\w+)\s+(\w+)\s+([\w-]+)\s+(\w+)\s+([0-9\-.:]+)\s?(.*)`)

func TestParsePeerRow(t *testing.T) {
	regx := parsePeerRowRegex(examplePeerRow)
	impl := parsePeerRow(examplePeerRow)

	assert.Equal(t, regx, impl)
}

func TestValidateHeader(t *testing.T) {
	regx := validateSummaryHeaderRegex(exampleProtocolHeader)
	impl := validateSummaryHeader(exampleProtocolHeader)

	assert.Equal(t, regx, impl)
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

var headerRegex = regexp.MustCompile(`Name\s+Proto\s+Table\s+State\s+Since\s+Info`)

func validateSummaryHeaderRegex(str string) bool {
	return headerRegex.MatchString(str)
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

func BenchmarkParseTable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = parseSummaryTable(exampleProtocolTable)
	}
}

func BenchmarkValidateHeaderRegex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		validateSummaryHeaderRegex(exampleProtocolHeader)
	}
}

func BenchmarkValidateHeader(b *testing.B) {
	for n := 0; n < b.N; n++ {
		validateSummaryHeader(exampleProtocolHeader)
	}
}
