package gocks

import (
	. "github.com/massyu/iota.go/api"
	. "github.com/massyu/iota.go/api/integration/samples"
	"gopkg.in/h2non/gock.v1"
)

func init() {
	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(BroadcastTransactionsCommand{Command: Command{BroadcastTransactionsCmd}, Trytes: BundleTrytes}).
		Reply(200)
}
