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

	"github.com/onflow/flow-nft/lib/go/templates"
)

func TestNFTDeployment(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	LeofyNFTAccountKey, _ := accountKeys.NewWithSigner()
	nftAddress, _, LeofyNFTAddress, _, _ := DeployContracts(b, t, []*flow.AccountKey{LeofyNFTAccountKey})

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		script := templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(0), supply)

		script = templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(0), length)
	})

}

func TestMintNFTs(t *testing.T) {
	b := newBlockchain()
	accountKeys := test.AccountKeyGenerator()

	LeofyNFTAccountKey, LeofyNFTSigner := accountKeys.NewWithSigner()
	nftAddress, metadataAddress, LeofyNFTAddress, _, _ := DeployContracts(b, t, []*flow.AccountKey{LeofyNFTAccountKey})

	// Create a new user account
	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	itemAuthor := cadence.String("John Doe")
	itemName := cadence.String("John Doe FirstArt")
	itemThumbnail := cadence.String("https://leofy.io/leofy-logo-y-2.svg")

	t.Run("Should be able to create a new Item", func(t *testing.T) {
		CreateItemTransaction(t, b, nftAddress, LeofyNFTAddress, LeofyNFTSigner, itemAuthor, itemName, itemThumbnail, CadenceUFix64("50.00"), false)

		script := templates.GenerateGetTotalItemsScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(1), supply)

		itemLength := executeScriptAndCheck(t, b, templates.GenerateGetItemsLengthScript(nftAddress, LeofyNFTAddress), nil)
		assert.Equal(t, cadence.NewInt(1), itemLength)
	})

	t.Run("Should be able to mint a token", func(t *testing.T) {
		MintNFTTransaction(t, b, nftAddress, LeofyNFTAddress, LeofyNFTSigner, 0, false)

		script := templates.GenerateBorrowNFTItemScript(nftAddress, LeofyNFTAddress)
		executeScriptAndCheck(
			t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(1), supply)

		script = templates.GenerateGetCollectionItemLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(
			t,
			b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			})
		assert.Equal(t, cadence.NewInt(1), length)
	})

	t.Run("Should be able to mint multiples token for a item", func(t *testing.T) {
		BulkMintNFTTransaction(t, b, nftAddress, LeofyNFTAddress, LeofyNFTSigner, 0, 20, false)

		script := templates.GenerateBorrowNFTItemScript(nftAddress, LeofyNFTAddress)
		executeScriptAndCheck(
			t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
				jsoncdc.MustEncode(cadence.NewUInt64(20)),
			},
		)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(21), supply)

		script = templates.GenerateGetCollectionItemLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewUInt64(0))})
		assert.Equal(t, cadence.NewInt(21), length)

	})

	t.Run("Should be able to withdraw an NFT from Item Collection and deposit to an account collection", func(t *testing.T) {
		script := templates.GenerateTransferNFTItemTransaction(nftAddress, LeofyNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

		tx.AddArgument(cadence.NewAddress(LeofyNFTAddress))
		tx.AddArgument(cadence.NewUInt64(0))
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

		script = templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(1), length)
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
		script := templates.GenerateBorrowNFTAccountScript(nftAddress, LeofyNFTAddress)

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
		CreateCollectionTransaction(t, b, nftAddress, LeofyNFTAddress, joshAddress, joshSigner, false)

		script := templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)
	})

	t.Run("Shouldn't be able to withdraw an NFT that doesn't exist in a collection", func(t *testing.T) {
		TransferNFTTransaction(t, b, nftAddress, LeofyNFTAddress, LeofyNFTAddress, LeofyNFTSigner, joshAddress, 100, true)

		script := templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(1), length)

	})

	// transfer an NFT
	t.Run("Should be able to withdraw an NFT and deposit to another accounts collection", func(t *testing.T) {
		TransferNFTTransaction(t, b, nftAddress, LeofyNFTAddress, LeofyNFTAddress, LeofyNFTSigner, joshAddress, 0, false)

		executeScriptAndCheck(
			t, b,
			templates.GenerateBorrowNFTAccountScript(nftAddress, LeofyNFTAddress),
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(joshAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		script := templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(1), length)

		script = templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(0), length)
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

		script = templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(LeofyNFTAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, LeofyNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(21), supply)
	})
}

func TestPurchaseNFT(t *testing.T) {
	b := newBlockchain()
	accountKeys := test.AccountKeyGenerator()

	LeofyAccountKey, LeofySigner := accountKeys.NewWithSigner()
	nftAddress, _, LeofyNFTAddress, fungibleAddr, leofyCoinAddr := DeployContracts(b, t, []*flow.AccountKey{LeofyAccountKey})

	// Create a new user account
	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	CreateTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner)
	CreateTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, LeofyNFTAddress, LeofySigner)

	getVaultScript(t, b, fungibleAddr, leofyCoinAddr, leofyCoinAddr, CadenceUFix64("1000.00"))
	getVaultScript(t, b, fungibleAddr, leofyCoinAddr, joshAddress, CadenceUFix64("0.00"))

	TransferTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, LeofySigner, joshAddress, CadenceUFix64("40.0"), false)
	CreateCollectionTransaction(t, b, nftAddress, LeofyNFTAddress, joshAddress, joshSigner, false)
	CreateItemTransaction(t, b, nftAddress, LeofyNFTAddress, LeofySigner, cadence.String("Author"), cadence.String("NFT Name"), cadence.String("NFT Name"), CadenceUFix64("50.00"), false)
	MintNFTTransaction(t, b, nftAddress, LeofyNFTAddress, LeofySigner, 0, false)

	// transfer an NFT
	t.Run("Shouldn't purchase if not enough tokens on balance for price item", func(t *testing.T) {
		PurchaseItemNFTTransaction(t, b, nftAddress, LeofyNFTAddress, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner, 0, true)

		script := templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetCollectionItemLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewUInt64(0))})
		assert.Equal(t, cadence.NewInt(1), length)

		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, leofyCoinAddr, CadenceUFix64("960.00"))
		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, joshAddress, CadenceUFix64("40.00"))
		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, LeofyNFTAddress, CadenceUFix64("0.00"))
	})

	// transfer an NFT
	t.Run("Should purcharse if has balance for price item", func(t *testing.T) {
		TransferTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, LeofySigner, joshAddress, CadenceUFix64("60.0"), false)
		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, leofyCoinAddr, CadenceUFix64("900.00"))
		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, joshAddress, CadenceUFix64("100.00"))

		PurchaseItemNFTTransaction(t, b, nftAddress, LeofyNFTAddress, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner, 0, false)

		script := templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(1), length)

		script = templates.GenerateGetCollectionItemLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewUInt64(0))})
		assert.Equal(t, cadence.NewInt(0), length)

		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, leofyCoinAddr, CadenceUFix64("900.00"))
		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, LeofyNFTAddress, CadenceUFix64("50.00"))
		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, joshAddress, CadenceUFix64("50.00"))
	})

	// transfer an NFT
	t.Run("Shouldn't purchase if NFT's lefts on item", func(t *testing.T) {
		PurchaseItemNFTTransaction(t, b, nftAddress, LeofyNFTAddress, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner, 0, true)

		script := templates.GenerateGetCollectionAccountLengthScript(nftAddress, LeofyNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(1), length)

		script = templates.GenerateGetCollectionItemLengthScript(nftAddress, LeofyNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewUInt64(0))})
		assert.Equal(t, cadence.NewInt(0), length)

		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, LeofyNFTAddress, CadenceUFix64("50.00"))
		getVaultScript(t, b, fungibleAddr, leofyCoinAddr, joshAddress, CadenceUFix64("50.00"))

	})

}
