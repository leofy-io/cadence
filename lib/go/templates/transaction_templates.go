package templates

import (
	"github.com/onflow/flow-go-sdk"

	_ "github.com/kevinburke/go-bindata"

	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	filenameSetupAccount = "LeofyNFT/setup_account.cdc"

	filenameGenerateItem = "LeofyNFT/create_item.cdc"

	filenameGenerateSet  = "LeofyNFT/create_set.cdc"
	filenameBatchMintNFT = "LeofyNFT/batch_mint_nft.cdc"
	filenameMintNFT      = "LeofyNFT/mint_nft.cdc"
	filenameTransferNFT  = "LeofyNFT/transfer_nft.cdc"
	filenameDestroyNFT   = "LeofyNFT/destroy_nft.cdc"
)

// -----------------------------------------------------------------------
// Item Transactions
// -----------------------------------------------------------------------

// GenerateItemTransaction creates a script that instantiates a new item.
func GenerateItemTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGenerateItem)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// -----------------------------------------------------------------------
// Set Transactions
// -----------------------------------------------------------------------

// GenerateMintSetTransaction creates a script that instantiates a new set.
func GenerateMintSetTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGenerateSet)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// -----------------------------------------------------------------------
// LeofyNFT Transactions
// -----------------------------------------------------------------------

// GenerateSetupAccountTransaction returns a script that instantiates a new
// NFT collection instance, saves the collection in storage, then stores a
// reference to the collection.
func GenerateSetupAccountTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameSetupAccount)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateMintNFTTransaction returns script that uses the admin resource
// to mint a new NFT and deposit it into a user's collection.
func GenerateMintNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameMintNFT)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

func GenerateBatchMintNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameBatchMintNFT)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateTransferNFTTransaction returns a script that withdraws an NFT token
// from a collection and deposits it into another collection.
func GenerateTransferNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameTransferNFT)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateDestroyNFTTransaction creates a script that withdraws an NFT token
// from a collection and destroys it.
func GenerateDestroyNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameDestroyNFT)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}
