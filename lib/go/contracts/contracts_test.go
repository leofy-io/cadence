package contracts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-go-sdk/test"

	"github.com/onflow/flow-nft/lib/go/contracts"
)

const addrA = "0x0A"

func TestNonFungibleTokenContract(t *testing.T) {
	contract := contracts.NonFungibleToken()
	assert.NotNil(t, contract)
}

func TestLeofyNFTContract(t *testing.T) {
	addresses := test.AddressGenerator()
	addressA := addresses.New()
	addressB := addresses.New()

	contract := contracts.LeofyNFT(addressA, addressB)
	assert.NotNil(t, contract)

	assert.Contains(t, string(contract), addressA.String())
	assert.Contains(t, string(contract), addressB.String())
}

func TestMetadataViewsContract(t *testing.T) {
	contract := contracts.MetadataViews()
	assert.NotNil(t, contract)
}
