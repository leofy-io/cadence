package test

import (
	"fmt"
	"testing"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-nft/lib/go/contracts"
	"github.com/onflow/flow-nft/lib/go/templates"
)

func TestGetNFTMetadata(t *testing.T) {
	b := newBlockchain()

	nftAddress := deploy(t, b, "NonFungibleToken", contracts.NonFungibleToken())
	metadataAddress := deploy(t, b, "MetadataViews", contracts.MetadataViews())

	accountKeys := test.AccountKeyGenerator()

	LeofyNFTAccountKey, LeofyNFTSigner := accountKeys.NewWithSigner()
	LeofyNFTAddress := deploy(
		t, b,
		"LeofyNFT",
		contracts.LeofyNFT(nftAddress, metadataAddress),
		LeofyNFTAccountKey,
	)

	script := templates.GenerateMintNFTTransaction(nftAddress, LeofyNFTAddress)

	tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

	const (
		itemId = 0
	)

	tx.AddArgument(cadence.NewAddress(LeofyNFTAddress))
	tx.AddArgument(cadence.NewUInt64(itemId))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			LeofyNFTAddress,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			LeofyNFTSigner,
		},
		false,
	)

	script = templates.GenerateGetNFTMetadataScript(nftAddress, LeofyNFTAddress, metadataAddress)
	result := executeScriptAndCheck(
		t, b,
		script,
		[][]byte{
			jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress)),
			jsoncdc.MustEncode(cadence.NewUInt64(0)),
		},
	)

	nftResult := result.(cadence.Struct)

	nftType := fmt.Sprintf("A.%s.LeofyNFT.NFT", LeofyNFTAddress)

	assert.Equal(t, cadence.NewAddress(LeofyNFTAddress), nftResult.Fields[3])
	assert.Equal(t, cadence.String(nftType), nftResult.Fields[4])
}
