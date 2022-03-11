package test

import (
	"testing"

	"github.com/onflow/cadence"
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-nft/lib/go/templates"
)

func CreateItemTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	transactionSigner crypto.Signer,
	itemAuthor cadence.String,
	itemName cadence.String,
	itemThumbnail cadence.String,
	price cadence.Value,
	shouldThrowError bool,
) {
	script := templates.GenerateItemTransaction(nftAddress, LeofyNFTAddress)
	tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

	metadata := []cadence.KeyValuePair{
		{Key: cadence.String("author"), Value: cadence.String("John Doe")},
		{Key: cadence.String("name"), Value: itemName},
		{Key: cadence.String("thumbnail"), Value: itemThumbnail},
	}

	tx.AddArgument(cadence.NewDictionary(metadata))
	tx.AddArgument(price)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			LeofyNFTAddress,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		shouldThrowError,
	)
}

func changeItemPriceTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	transactionSigner crypto.Signer,
	itemID uint64,
	price cadence.Value,
	shouldThrowError bool,
) {
	script := templates.GenerateChangeItemPriceTransaction(nftAddress, LeofyNFTAddress)
	tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

	tx.AddArgument(cadence.NewUInt64(itemID))
	tx.AddArgument(price)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			LeofyNFTAddress,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		shouldThrowError,
	)
}

func MintNFTTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	transactionSigner crypto.Signer,
	itemID uint64,
	shouldThrowError bool,
) {
	script := templates.GenerateMintNFTTransaction(nftAddress, LeofyNFTAddress)

	tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)
	tx.AddArgument(cadence.NewUInt64(itemID))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			LeofyNFTAddress,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		false,
	)
}

func BulkMintNFTTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	transactionSigner crypto.Signer,
	itemID uint64,
	quantity uint64,
	shouldThrowError bool,
) {
	tx := createTxWithTemplateAndAuthorizer(
		b,
		templates.GenerateBatchMintNFTTransaction(nftAddress, LeofyNFTAddress),
		LeofyNFTAddress,
	)

	tx.AddArgument(cadence.NewUInt64(itemID))
	tx.AddArgument(cadence.NewUInt64(quantity))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			LeofyNFTAddress,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		false,
	)
}

func CreateCollectionTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	transactionAuthorizer flow.Address,
	transactionSigner crypto.Signer,
	shouldThrowError bool,
) {

	script := templates.GenerateSetupAccountNFTTransaction(nftAddress, LeofyNFTAddress)
	tx := createTxWithTemplateAndAuthorizer(b, script, transactionAuthorizer)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			transactionAuthorizer,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		false,
	)

}

func TransferNFTTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	transactionAuthorizer flow.Address,
	transactionSigner crypto.Signer,
	nftReceiverAddr flow.Address,
	nftId uint64,
	shouldThrowError bool,
) {

	script := templates.GenerateTransferNFTAccountTransaction(nftAddress, LeofyNFTAddress)
	tx := createTxWithTemplateAndAuthorizer(b, script, LeofyNFTAddress)

	tx.AddArgument(cadence.NewAddress(nftReceiverAddr))
	tx.AddArgument(cadence.NewUInt64(nftId))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			transactionAuthorizer,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		shouldThrowError,
	)
}

func PurchaseItemNFTTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	nftAddress flow.Address,
	LeofyNFTAddress flow.Address,
	ftAddress flow.Address,
	LeofyCoinAddress flow.Address,
	transactionAuthorizer flow.Address,
	transactionSigner crypto.Signer,
	itemId uint64,
	shouldThrowError bool,
) {

	script := templates.GeneratePurchaseItemNFTTransaction(nftAddress, LeofyNFTAddress, ftAddress, LeofyCoinAddress)
	tx := createTxWithTemplateAndAuthorizer(b, script, transactionAuthorizer)

	tx.AddArgument(cadence.NewUInt64(itemId))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			transactionAuthorizer,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		shouldThrowError,
	)

}
