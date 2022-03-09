package test

import (
	"testing"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-nft/lib/go/templates"

	"github.com/stretchr/testify/assert"
)

func CreateTokenTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	fungibleAddr flow.Address,
	leofyCoinAddr flow.Address,
	joshAddress flow.Address,
	transactionSigner crypto.Signer,
) {
	script := templates.GenerateCreateTokenScript(fungibleAddr, leofyCoinAddr)
	tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			joshAddress,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		false,
	)

	script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
	result := executeScriptAndCheck(t, b,
		script,
		[][]byte{
			jsoncdc.MustEncode(cadence.Address(joshAddress)),
		},
	)

	assert.Equal(t, CadenceUFix64("0.0"), result)

}

func TransferTokenTransaction(
	t *testing.T,
	b *emulator.Blockchain,
	fungibleAddr flow.Address,
	leofyCoinAddr flow.Address,
	transactionSigner crypto.Signer,
	receiverAddress flow.Address,
	amount cadence.Value,
	shouldThrowError bool,
) {
	script := templates.GenerateTransferVaultScript(fungibleAddr, leofyCoinAddr)
	tx := createTxWithTemplateAndAuthorizer(b, script, leofyCoinAddr)

	tx.AddArgument(amount)
	tx.AddArgument(cadence.NewAddress(receiverAddress))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			leofyCoinAddr,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			transactionSigner,
		},
		shouldThrowError,
	)
}

func getVaultScript(
	t *testing.T,
	b *emulator.Blockchain,
	fungibleAddr flow.Address,
	leofyCoinAddr flow.Address,
	addressVault flow.Address,
	amount cadence.Value,
) {
	script := templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
	result := executeScriptAndCheck(t, b,
		script,
		[][]byte{
			jsoncdc.MustEncode(cadence.Address(addressVault)),
		},
	)

	assert.Equal(t, amount, result)
}
