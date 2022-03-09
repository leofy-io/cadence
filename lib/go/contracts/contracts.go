package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-nft/lib/go/contracts/internal/assets"
)

var (
	placeholderNonFungibleToken = regexp.MustCompile(`"[^"\s].*/NonFungibleToken.cdc"`)
	placeholderFungibleToken    = regexp.MustCompile(`"[^"\s].*/FungibleToken.cdc"`)
	placeholderMetadataViews    = regexp.MustCompile(`"[^"\s].*/MetadataViews.cdc"`)
	placeholderLeofyCoin        = regexp.MustCompile(`"[^"\s].*/LeofyCoin.cdc"`)
)

const (
	filenameLeofyNFT  = "LeofyNFT.cdc"
	filenameLeofyCoin = "LeofyCoin.cdc"
)

// LeofyNFT returns the LeofyNFT contract.
//
// The returned contract will import the NonFungibleToken contract from the specified address.
func LeofyNFT(nftAddress, metadataAddress flow.Address, leofyCoinAddress flow.Address, ftAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameLeofyNFT)

	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataAddress.String())
	code = placeholderLeofyCoin.ReplaceAllString(code, "0x"+leofyCoinAddress.String())
	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+ftAddress.String())
	//code = placeholderLeofyCoin.ReplaceAllString(code, "0x"+leofyCoinAddress.String())

	return []byte(code)
}

func LeofyCoin(ftAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameLeofyCoin)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+ftAddress.String())

	return []byte(code)
}
