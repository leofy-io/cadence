package templates

import (
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	filenameItemsTotal  = "LeofyNFT/scripts/get_items_supply.cdc"
	filenameItemsLength = "LeofyNFT/scripts/get_items_length.cdc"

	filenameGetSetName      = "LeofyNFT/scripts/get_set_name.cdc"
	filenameGetSetIdsByName = "LeofyNFT/scripts/get_set_ids_by_name.cdc"

	filenameBorrowNFT           = "LeofyNFT/scripts/borrow_nft.cdc"
	filenameGetCollectionLength = "LeofyNFT/scripts/get_collection_length.cdc"
	filenameGetTotalSupply      = "LeofyNFT/scripts/get_total_supply.cdc"

	filenameGetLeofyNFTMetadata = "LeofyNFT/scripts/get_nft_leofy_metadata.cdc"
	filenameGetNFTMetadata      = "LeofyNFT/scripts/get_nft_metadata.cdc"
)

// -----------------------------------------------------------------------
// Item Scripts
// -----------------------------------------------------------------------
func GenerateGetTotalItemsScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameItemsTotal)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

func GenerateGetItemsLengthScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameItemsLength)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// -----------------------------------------------------------------------
// Item Scripts
// -----------------------------------------------------------------------
func GenerateGetSetNameScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetSetName)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

func GenerateGetSetIDsByNameScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetSetIdsByName)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// -----------------------------------------------------------------------
// LeofyNFT Scripts
// -----------------------------------------------------------------------

// GenerateBorrowNFTScript creates a script that retrieves an NFT collection
// from storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateBorrowNFTScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameBorrowNFT)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateGetNFTMetadataScript creates a script that returns the metadata for an NFT.
func GenerateGetNFTMetadataScript(nftAddress, LeofyNFTAddress, metadataAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetNFTMetadata)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, metadataAddress)
}

func GenerateGetLeofyNFTMetadataScript(nftAddress, LeofyNFTAddress, metadataAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetLeofyNFTMetadata)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, metadataAddress)
}

// GenerateGetCollectionLengthScript creates a script that retrieves an NFT collection
// from storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetCollectionLength)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateGetTotalSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetTotalSupply)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}
