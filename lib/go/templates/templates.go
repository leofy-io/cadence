package templates

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

var (
	placeholderNonFungibleToken = regexp.MustCompile(`"[^"\s].*/NonFungibleToken.cdc"`)
	placeholderLeofyNFT         = regexp.MustCompile(`"[^"\s].*/LeofyNFT.cdc"`)
	placeholderMetadataViews    = regexp.MustCompile(`"[^"\s].*/MetadataViews.cdc"`)
	placeholderFungibleToken    = regexp.MustCompile(`"[^"\s].*/FungibleToken.cdc"`)
	placeholderLeofyCoin        = regexp.MustCompile(`"[^"\s].*/LeofyCoin.cdc"`)
)

func replaceAddressesLeofyNFT(code string, nftAddress, LeofyNFTAddress flow.Address, metadataAddress flow.Address) []byte {
	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())
	code = placeholderLeofyNFT.ReplaceAllString(code, "0x"+LeofyNFTAddress.String())
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataAddress.String())
	return []byte(code)
}

func replaceAddressesLeofyCoin(code string, ftAddress flow.Address, LeofyCoinAddress flow.Address) []byte {
	code = placeholderLeofyCoin.ReplaceAllString(code, "0x"+LeofyCoinAddress.String())
	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+ftAddress.String())
	return []byte(code)
}

func replaceAddresses(
	code string,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	metadataAddress flow.Address,
	ftAddress flow.Address,
	LeofyCoinAddress flow.Address,
) []byte {
	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())
	code = placeholderLeofyNFT.ReplaceAllString(code, "0x"+LeofyNFTAddress.String())
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataAddress.String())
	code = placeholderLeofyCoin.ReplaceAllString(code, "0x"+LeofyCoinAddress.String())
	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+ftAddress.String())
	return []byte(code)
}
