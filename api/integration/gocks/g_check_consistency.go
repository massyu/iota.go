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
		JSON(CheckConsistencyCommand{Command: Command{CheckConsistencyCmd}, Tails: DefaultHashes()}).
		Reply(200).
		JSON(CheckConsistencyResponse{State: true, Info: ""})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(CheckConsistencyCommand{
			Command: Command{CheckConsistencyCmd},
			Tails:   append(DefaultHashes(), strings.Repeat("C", 81)),
		}).
		Reply(200).
		JSON(CheckConsistencyResponse{State: false, Info: "test response"})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(CheckConsistencyCommand{
			Command: Command{CheckConsistencyCmd},
			Tails:   Hashes{Bundle[0].Hash},
		}).
		Reply(200).
		JSON(CheckConsistencyResponse{State: true, Info: ""})
}
