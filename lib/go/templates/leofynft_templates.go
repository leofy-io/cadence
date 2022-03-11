package templates

import (
	"github.com/onflow/flow-go-sdk"

	_ "github.com/kevinburke/go-bindata"

	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	filenameSetupAccount = "LeofyNFT/setup_account.cdc"

	filenameGenerateItem    = "LeofyNFT/create_item.cdc"
	filenameChangeItemPrice = "LeofyNFT/change_item_price.cdc"

	filenameBatchMintNFT       = "LeofyNFT/batch_mint_nft.cdc"
	filenameMintNFT            = "LeofyNFT/mint_nft.cdc"
	filenameTransferNFTAccount = "LeofyNFT/transfer_nft_account.cdc"
	filenameTransferNFTItem    = "LeofyNFT/transfer_nft_item.cdc"
	filenameDestroyNFT         = "LeofyNFT/destroy_nft.cdc"

	filenamePurchaseItemNFT = "LeofyNFT/purchase_nft.cdc"

	//scripts

	filenameItemsTotal  = "LeofyNFT/scripts/get_items_supply.cdc"
	filenameItemsLength = "LeofyNFT/scripts/get_items_length.cdc"

	filenameItemPrice = "LeofyNFT/scripts/get_item_price.cdc"

	filenameBorrowNFTAccount = "LeofyNFT/scripts/borrow_nft_account.cdc"
	filenameBorrowNFTItem    = "LeofyNFT/scripts/borrow_nft_item.cdc"

	filenameGetCollectionItemLength    = "LeofyNFT/scripts/get_collection_item_length.cdc"
	filenameGetCollectionAccountLength = "LeofyNFT/scripts/get_collection_account_length.cdc"
	filenameGetTotalSupply             = "LeofyNFT/scripts/get_total_supply.cdc"

	filenameGetLeofyNFTMetadata = "LeofyNFT/scripts/get_nft_leofy_metadata.cdc"
	filenameGetNFTMetadata      = "LeofyNFT/scripts/get_nft_metadata.cdc"
)

// -----------------------------------------------------------------------
// Item Transactions
// -----------------------------------------------------------------------

// GenerateItemTransaction creates a script that instantiates a new item.
func GenerateItemTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGenerateItem)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateChangeItemPriceTransaction creates a script that instantiates a new item and change their price.
func GenerateChangeItemPriceTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameChangeItemPrice)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// -----------------------------------------------------------------------
// LeofyNFT Transactions
// -----------------------------------------------------------------------

// GenerateSetupAccountNFTTransaction returns a script that instantiates a new
// NFT collection instance, saves the collection in storage, then stores a
// reference to the collection.
func GenerateSetupAccountNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameSetupAccount)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateMintNFTTransaction returns script that uses the admin resource
// to mint a new NFT and deposit it into a user's collection.
func GenerateMintNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameMintNFT)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

func GenerateBatchMintNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameBatchMintNFT)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateTransferNFTTransaction returns a script that withdraws an NFT token
// from an Item Collection and deposits it into another collection.
func GenerateTransferNFTItemTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameTransferNFTItem)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateTransferNFTTransaction returns a script that withdraws an NFT token
// from an account collection and deposits it into another collection.
func GenerateTransferNFTAccountTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameTransferNFTAccount)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateDestroyNFTTransaction creates a script that withdraws an NFT token
// from a collection and destroys it.
func GenerateDestroyNFTTransaction(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameDestroyNFT)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateDestroyNFTTransaction creates a script that withdraws an NFT token
// from a collection and destroys it.
func GeneratePurchaseItemNFTTransaction(
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	ftAddress flow.Address,
	LeofyCoinAddress flow.Address,
) []byte {
	code := assets.MustAssetString(filenamePurchaseItemNFT)
	return replaceAddresses(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress, ftAddress, LeofyCoinAddress)
}

// -----------------------------------------------------------------------
// Item Scripts
// -----------------------------------------------------------------------
func GenerateGetTotalItemsScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameItemsTotal)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

func GenerateGetItemsLengthScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameItemsLength)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

func GenerateGetItemsPrice(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameItemPrice)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// -----------------------------------------------------------------------
// LeofyNFT Scripts
// -----------------------------------------------------------------------

// GenerateBorrowNFTAcountScript creates a script that retrieves an NFT collection
// from account storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateBorrowNFTAccountScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameBorrowNFTAccount)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateBorrowNFTScript creates a script that retrieves an NFT collection
// from ItemCollection storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateBorrowNFTItemScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameBorrowNFTItem)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateGetNFTMetadataScript creates a script that returns the metadata for an NFT.
func GenerateGetNFTMetadataScript(nftAddress, LeofyNFTAddress, metadataAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetNFTMetadata)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, metadataAddress)
}

func GenerateGetLeofyNFTMetadataScript(nftAddress, LeofyNFTAddress, metadataAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetLeofyNFTMetadata)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, metadataAddress)
}

// GenerateGetCollectionItemLengthScript creates a script that retrieves an NFT collection
// from Item storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateGetCollectionItemLengthScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetCollectionItemLength)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateGetCollectionAccountLengthScript creates a script that retrieves an NFT collection
// from Account storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetCollectionAccountLength)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}

// GenerateGetTotalSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameGetTotalSupply)
	return replaceAddressesLeofyNFT(code, nftAddress, LeofyNFTAddress, flow.EmptyAddress)
}
