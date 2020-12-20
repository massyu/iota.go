package gocks

import (
	"strings"

	. "github.com/massyu/iota.go/api"
	. "github.com/massyu/iota.go/api/integration/samples"
	"github.com/massyu/iota.go/consts"
	. "github.com/massyu/iota.go/trinary"
	"gopkg.in/h2non/gock.v1"
)

func init() {

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(GetTrytesCommand{Command: Command{GetTrytesCmd}, Hashes: DefaultHashes()}).
		Reply(200).
		JSON(GetTrytesResponse{Trytes: []Trytes{
			strings.Repeat("9", consts.TransactionTrytesSize),
			strings.Repeat("9", consts.TransactionTrytesSize),
		}})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(GetTrytesCommand{
			Command: Command{GetTrytesCmd},
			Hashes:  Hashes{BundleWithZeroValue[0].Hash},
		}).
		Reply(200).
		JSON(GetTrytesResponse{Trytes: BundleWithZeroValueTrytes})

	for i := range Bundle {
		gock.New(DefaultLocalIRIURI).
			Persist().
			Post("/").
			MatchType("json").
			JSON(GetTrytesCommand{
				Command: Command{GetTrytesCmd},
				Hashes:  Hashes{Bundle[i].Hash},
			}).
			Reply(200).
			JSON(GetTrytesResponse{Trytes: []Trytes{BundleTrytes[i]}})
	}

}
