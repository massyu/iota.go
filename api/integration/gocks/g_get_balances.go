package gocks

import (
	"strings"

	. "github.com/massyu/iota.go/api"
	. "github.com/massyu/iota.go/api/integration/samples"
	"gopkg.in/h2non/gock.v1"
)

func init() {
	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(GetBalancesCommand{
			Command:   Command{GetBalancesCmd},
			Addresses: SampleAddresses,
		}).
		Reply(200).
		JSON(GetBalancesResponse{
			Duration:       100,
			Balances:       []string{"99", "0", "1"},
			Milestone:      strings.Repeat("M", 81),
			MilestoneIndex: 1,
		})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(GetBalancesCommand{
			Command:   Command{GetBalancesCmd},
			Addresses: SampleAddresses[1:],
		}).
		Reply(200).
		JSON(GetBalancesResponse{
			Duration:       100,
			Balances:       []string{"0", "1"},
			Milestone:      strings.Repeat("M", 81),
			MilestoneIndex: 1,
		})
}
