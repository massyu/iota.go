package gocks

import (
	"strings"

	. "github.com/massyu/iota.go/api"
	. "github.com/massyu/iota.go/api/integration/samples"
	. "github.com/massyu/iota.go/trinary"
	"gopkg.in/h2non/gock.v1"
)

func init() {
	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(GetInclusionStatesCommand{
			Command:      Command{GetInclusionStatesCmd},
			Transactions: DefaultHashes(),
		}).
		Reply(200).
		JSON(GetInclusionStatesResponse{States: []bool{true, false}})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(GetInclusionStatesCommand{
			Command: Command{GetInclusionStatesCmd},
			Transactions: Hashes{
				strings.Repeat("9", 81),
			},
		}).
		Reply(200).
		JSON(GetInclusionStatesResponse{States: []bool{true, false}})
}
