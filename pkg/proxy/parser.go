package proxy

import (
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type SummaryTable []PeerRow

type PeerRow struct {
	Name  string `json:"name"`
	Proto string `json:"proto"`
	Table string `json:"table"`
	State string `json:"state"`
	Since string `json:"since"`
	Info  string `json:"info"`
}

func (SummaryTable) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

type PeerDetails struct {
	Info    PeerRow `json:"info"`
	Details string  `json:"details"`
}

func (PeerDetails) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

type RouteDetails struct {
	Address string `json:"address"`
	Details string `json:"details"`
}

func (RouteDetails) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

func splitStringCoalesceSpaces(str string) []string {
	split := strings.Split(str, " ")

	i := 0
	for si := range split {
		if split[si] == "" {
			continue
		}

		split[i] = split[si]
		i++
	}

	return split[:i+1]
}

func parsePeerRow(line string) PeerRow {
	split := splitStringCoalesceSpaces(line)

	return PeerRow{
		Name:  split[0],
		Proto: split[1],
		Table: split[2],
		State: split[3],
		Since: split[4],
		Info:  split[5],
	}
}

func validateSummaryHeader(str string) bool {
	split := splitStringCoalesceSpaces(str)

	return split[0] == "Name" && split[1] == "Proto" && split[2] == "Table" && split[3] == "State" && split[4] == "Since" && split[5] == "Info"
}

func parseSummaryTable(str string) (SummaryTable, error) {
	rows := strings.Split(str, "\n")
	if !validateSummaryHeader(rows[0]) {
		return nil, errors.New("invalid protocol table")
	}

	var table SummaryTable
	for i := 0; i < len(rows)-1; i++ {
		row := strings.TrimSpace(rows[i+1])
		if row == "" {
			continue
		}
		table = append(table, parsePeerRow(row))
	}

	return table, nil
}

func parsePeerDetails(str string) (*PeerDetails, error) {
	details := strings.SplitN(str, "\n", 3)
	if len(details) != 3 || !validateSummaryHeader(details[0]) {
		return nil, errors.New("invalid protocol table")
	}

	details[1] = strings.TrimSpace(details[1])
	header := parsePeerRow(details[1])

	return &PeerDetails{
		Info:    header,
		Details: details[2],
	}, nil
}
