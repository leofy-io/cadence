package test

import (
	"testing"

	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-nft/lib/go/contracts"
)

func DeployContracts(
	b *emulator.Blockchain,
	t *testing.T,
	key []*flow.AccountKey,
) (
	nftAddress flow.Address,
	metadataAddress flow.Address,
	LeofyNFTAddress flow.Address,
	ftAddress flow.Address,
	LeofyCoinAddress flow.Address,
) {
	var err error

	// Should be able to deploy a contract as a new account with no keys.
	nftCode, _ := DownloadFile(NonFungibleTokenContractsBaseFile)
	nftAddress, err = b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "NonFungibleToken",
			Source: string(nftCode),
		},
	})

	assert.NoError(t, err)

	// Should be able to deploy the MetadataViews contract with no keys.
	metadataCode, _ := DownloadFile(MetadataViewsContractsBaseFile)
	metadataAddress, err = b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "MetadataViews",
			Source: string(metadataCode),
		},
	})

	assert.NoError(t, err)

	// Should be able to deploy the Fungible contract with no keys.
	ftCode, _ := DownloadFile(FungibleTokenContractsBaseFile)
	ftAddress, err = b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "FungibleToken",
			Source: string(ftCode),
		},
	})

	assert.NoError(t, err)

	LeofyCoinCode := contracts.LeofyCoin(ftAddress)

	LeofyCoinAddress, err = b.CreateAccount(
		key,
		[]sdktemplates.Contract{
			{
				Name:   "LeofyCoin",
				Source: string(LeofyCoinCode),
			},
		},
	)
	assert.NoError(t, err)

	LeofyNFTCode := contracts.LeofyNFT(nftAddress, metadataAddress, LeofyCoinAddress, ftAddress)

	LeofyNFTAddress, err = b.CreateAccount(
		key,
		[]sdktemplates.Contract{
			{
				Name:   "LeofyNFT",
				Source: string(LeofyNFTCode),
			},
		},
	)
	assert.NoError(t, err)

	return nftAddress, metadataAddress, LeofyNFTAddress, ftAddress, LeofyCoinAddress
}
