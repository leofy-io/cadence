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
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-nft/lib/go/contracts"
	"github.com/onflow/flow-nft/lib/go/templates"
)

func TestNFTDeployment(t *testing.T) {
	b := newBlockchain()

	nftAddress := deploy(t, b, "NonFungibleToken", contracts.NonFungibleToken())
	metadataAddress := deploy(t, b, "MetadataViews", contracts.MetadataViews())

	_ = deploy(
		t, b,
		"LeofyNFT",
		contracts.LeofyNFT(nftAddress, metadataAddress),
	)
}

func TestMintNFTs(t *testing.T) {
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

	script := templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
	supply := executeScriptAndCheck(t, b, script, nil)
	assert.Equal(t, cadence.NewUInt64(0), supply)

	script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
	length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
	assert.Equal(t, cadence.NewInt(0), length)

	// Create a new user account
	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	itemAuthor := cadence.String("John Doe")
	itemName := cadence.String("John Doe FirstArt")
	itemThumbnail := cadence.String("https://leofy.io/leofy-logo-y-2.svg")

	t.Run("Should be able to create a new Item", func(t *testing.T) {
		script := templates.GenerateGetTotalItemsScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(0), supply)

		script = templates.GenerateItemTransaction(nftAddress, LeofyNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

		metadata := []cadence.KeyValuePair{
			{Key: cadence.String("author"), Value: itemAuthor},
			{Key: cadence.String("name"), Value: itemName},
			{Key: cadence.String("thumbnail"), Value: itemThumbnail},
		}
		play := cadence.NewDictionary(metadata)
		_ = tx.AddArgument(play)

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

		script = templates.GenerateGetTotalItemsScript(nftAddress, LeofyNFTAddress)
		supply = executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(1), supply)

		itemLength := executeScriptAndCheck(t, b, templates.GenerateGetItemsLengthScript(nftAddress, LeofyNFTAddress), nil)
		assert.Equal(t, cadence.NewInt(1), itemLength)
	})

	t.Run("Should be able to mint a token", func(t *testing.T) {

		script := templates.GenerateMintNFTTransaction(nftAddress, LeofyNFTAddress)

		tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

		tx.AddArgument(cadence.NewAddress(LeofyNFTAddress))
		tx.AddArgument(cadence.NewUInt64(0))

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

		script = templates.GenerateBorrowNFTScript(nftAddress, LeofyNFTAddress)
		executeScriptAndCheck(
			t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(1), supply)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(1), length)
	})

	t.Run("Should be able to mint multiples token for a item", func(t *testing.T) {
		tx := createTxWithTemplateAndAuthorizer(
			b,
			templates.GenerateBatchMintNFTTransaction(nftAddress, LeofyNFTAddress),
			LeofyNFTAddress,
		)

		tx.AddArgument(cadence.NewAddress(LeofyNFTAddress))
		tx.AddArgument(cadence.NewUInt64(0))
		tx.AddArgument(cadence.NewUInt64(20))

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

		script = templates.GenerateBorrowNFTScript(nftAddress, LeofyNFTAddress)
		executeScriptAndCheck(
			t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(20)),
			},
		)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(21), supply)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(21), length)

	})

	t.Run("Should be able to get moments metadata", func(t *testing.T) {
		resultNFT := executeScriptAndCheck(
			t, b,
			templates.GenerateGetNFTMetadataScript(nftAddress, LeofyNFTAddress, metadataAddress),
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		metadataViewNFT := resultNFT.(cadence.Struct)
		nftType := fmt.Sprintf("A.%s.LeofyNFT.NFT", LeofyNFTAddress)

		assert.Equal(t, itemName, metadataViewNFT.Fields[0])
		assert.Equal(t, cadence.String("NFT: 'John Doe FirstArt' from Author: 'John Doe' with serial number 1"), metadataViewNFT.Fields[1])
		assert.Equal(t, itemThumbnail, metadataViewNFT.Fields[2])
		assert.Equal(t, cadence.NewAddress(LeofyNFTAddress), metadataViewNFT.Fields[3])
		assert.Equal(t, cadence.String(nftType), metadataViewNFT.Fields[4])

		resultLeofyNFT := executeScriptAndCheck(
			t, b,
			templates.GenerateGetLeofyNFTMetadataScript(nftAddress, LeofyNFTAddress, metadataAddress),
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		metadataViewLeofyNFT := resultLeofyNFT.(cadence.Struct)

		assert.Equal(t, itemAuthor, metadataViewLeofyNFT.Fields[0])
		assert.Equal(t, itemName, metadataViewLeofyNFT.Fields[1])
		assert.Equal(t, cadence.String("NFT: 'John Doe FirstArt' from Author: 'John Doe' with serial number 1"), metadataViewLeofyNFT.Fields[2])
		assert.Equal(t, itemThumbnail, metadataViewLeofyNFT.Fields[3])
		assert.Equal(t, cadence.NewUInt64(0), metadataViewLeofyNFT.Fields[4])
		assert.Equal(t, cadence.NewUInt32(1), metadataViewLeofyNFT.Fields[5])
	})

	t.Run("Shouldn't be able to borrow a reference to an NFT that doesn't exist", func(t *testing.T) {
		script := templates.GenerateBorrowNFTScript(nftAddress, LeofyNFTAddress)

		result, err := b.ExecuteScript(
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(100)),
			},
		)
		require.NoError(t, err)

		assert.True(t, result.Reverted())
	})

	// create a new Collection
	t.Run("Should be able to create a new empty NFT Collection", func(t *testing.T) {

		script := templates.GenerateSetupAccountTransaction(nftAddress, LeofyNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				joshSigner,
			},
			false,
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)
	})

	t.Run("Shouldn't be able to withdraw an NFT that doesn't exist in a collection", func(t *testing.T) {

		script := templates.GenerateTransferNFTTransaction(nftAddress, LeofyNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

		tx.AddArgument(cadence.NewAddress(joshAddress))
		tx.AddArgument(cadence.NewUInt64(100))

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
			true,
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(21), length)
	})

	// transfer an NFT
	t.Run("Should be able to withdraw an NFT and deposit to another accounts collection", func(t *testing.T) {
		script := templates.GenerateTransferNFTTransaction(nftAddress, LeofyNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

		tx.AddArgument(cadence.NewAddress(joshAddress))
		tx.AddArgument(cadence.NewUInt64(0))

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

		executeScriptAndCheck(
			t, b,
			templates.GenerateBorrowNFTScript(nftAddress, LeofyNFTAddress),
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(joshAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(1), length)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(20), length)
	})

	// transfer an NFT
	t.Run("Should be able to withdraw an NFT and destroy it, not reducing the supply", func(t *testing.T) {

		script := templates.GenerateDestroyNFTTransaction(nftAddress, LeofyNFTAddress)

		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		tx.AddArgument(cadence.NewUInt64(0))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				joshSigner,
			},
			false,
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(20), length)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(21), supply)
	})
}
