package integration_test

import (
	"strings"

	. "github.com/massyu/iota.go/api"
	. "github.com/massyu/iota.go/api/integration/samples"
	"github.com/massyu/iota.go/bundle"
	"github.com/massyu/iota.go/checksum"
	. "github.com/massyu/iota.go/consts"
	. "github.com/massyu/iota.go/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("PrepareTransfers()", func() {

	api, err := ComposeAPI(HTTPClientSettings{}, nil)
	if err != nil {
		panic(err)
	}

	inputs := []Input{
		{
			Address:  SampleAddressesWithChecksum[0],
			KeyIndex: 0,
			Security: 2,
			Balance:  3,
		},
		{
			Address:  SampleAddressesWithChecksum[1],
			KeyIndex: 1,
			Security: 2,
			Balance:  4,
		},
	}

	a, err := checksum.AddChecksum(strings.Repeat("A", 81), true, AddressChecksumTrytesSize)
	if err != nil {
		panic(err)
	}
	b, err := checksum.AddChecksum(strings.Repeat("B", 81), true, AddressChecksumTrytesSize)
	if err != nil {
		panic(err)
	}

	transfers := bundle.Transfers{
		{
			Address: a,
			Value:   3,
			Tag:     "TAG",
			Message: "9",
		},
		{
			Address: b,
			Value:   3,
			Tag:     "TAG",
		},
	}

	targetAddr, err := checksum.AddChecksum("OHXRRYM9XAOOXBLWIFWSMMDUYSRVRK9RWHPMNRFDTKUYZWENMPGHPHKBECU9HRJMAYSQM9JRAS9CTGWBW", true, AddressChecksumTrytesSize)
	if err != nil {
		panic(err)
	}
	zeroValueTransfer := bundle.Transfers{
		{
			Address: targetAddr,
			Value:   0,
			Tag:     "DJBETBPXOIKY",
			Message: "K9X",
		},
	}

	expectedZeroValueTrytes := []Trytes{
		"K9X999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999OHXRRYM9XAOOXBLWIFWSMMDUYSRVRK9RWHPMNRFDTKUYZWENMPGHPHKBECU9HRJMAYSQM9JRAS9CTGWBW999999999999999999999999999NNCETBPXOIKY999999999999999T9XRIZD99999999999999999999VRBLQHKIUGAFFPBTZROLCDHHCOVPXCJNFBRSNZIOJCCHVZNBSKD99LUOMV9AKLWIKCGBY9UZWCBNPLWYC999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999DJBETBPXOIKY999999999999999999999999999999999999999999999999999999999999999999999",
	}

	remainderAddress := SampleAddressesWithChecksum[2]

	// necessary so that the output is going to be the same no matter when it is run
	timestamp := uint64(1522219924)

	Context("call", func() {

		It("returns correctly prepared trytes for transfer", func() {
			trnfs, err := api.PrepareTransfers(Seed, transfers, PrepareTransfersOptions{Inputs: inputs, Timestamp: &timestamp, RemainderAddress: &remainderAddress})
			Expect(err).ToNot(HaveOccurred())
			Expect(trnfs).To(Equal(ExpectedPrepareTransfersTrytes))
		})

		It("resolves to correct account data with zero value transfers", func() {
			ts := uint64(1539936677)
			trnfs, err := api.PrepareTransfers(Seed, zeroValueTransfer, PrepareTransfersOptions{Timestamp: &ts})
			Expect(err).ToNot(HaveOccurred())
			Expect(trnfs).To(Equal(expectedZeroValueTrytes))
		})
	})

	Context("invalid input", func() {
		It("returns an error for invalid seed", func() {
			_, err := api.PrepareTransfers("asdf", zeroValueTransfer, PrepareTransfersOptions{})
			Expect(errors.Cause(err)).To(Equal(ErrInvalidSeed))
		})
	})

})
