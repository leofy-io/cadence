package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"

	"github.com/onflow/flow-nft/lib/go/templates"
)

func TestTokenDeployment(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	leofyCoinAccountKey, _ := accountKeys.NewWithSigner()
	_, _, _, fungibleAddr, leofyCoinAddr := DeployContracts(b, t, []*flow.AccountKey{leofyCoinAccountKey})

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		script := templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}

func TestCreateToken(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	leofyCoinAccountKey, _ := accountKeys.NewWithSigner()
	_, _, _, fungibleAddr, leofyCoinAddr := DeployContracts(b, t, []*flow.AccountKey{leofyCoinAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	t.Run("Should be able to create empty Vault that doesn't affect supply", func(t *testing.T) {
		CreateTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner)

		script := templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})
}

func TestExternalTransfers(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	leofyCoinAccountKey, leofyCoinSigner := accountKeys.NewWithSigner()
	_, _, _, fungibleAddr, leofyCoinAddr :=
		DeployContracts(b, t, []*flow.AccountKey{leofyCoinAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	CreateTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner)

	t.Run("Shouldn't be able to withdraw more than the balance of the Vault", func(t *testing.T) {
		// Transfer more than vault balance
		TransferTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, leofyCoinSigner, joshAddress, CadenceUFix64("30000.0"), true)

		// Assert that the vaults' balances are correct
		script := templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(leofyCoinAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("1000.0"), result)

		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("0.0"), result)
	})

	t.Run("Should be able to withdraw and deposit tokens from a vault", func(t *testing.T) {
		// Transfer 300 Tokens to JoshAddress
		TransferTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, leofyCoinSigner, joshAddress, CadenceUFix64("300.0"), false)

		// Assert that the vaults' balances are correct
		script := templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(leofyCoinAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("700.0"), result)

		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("300.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should be able to transfer to multiple accounts ", func(t *testing.T) {

		/*recipient1Address := cadence.Address(joshAddress)
		recipient1Amount := CadenceUFix64("300.0")

		pair := cadence.KeyValuePair{Key: recipient1Address, Value: recipient1Amount}
		recipientPairs := make([]cadence.KeyValuePair, 1)
		recipientPairs[0] = pair

		script := templates.GenerateTransferManyAccountsScript(fungibleAddr, leofyCoinAddr)

		tx := flow.NewTransaction().
			SetScript(script).
			SetGasLimit(100).
			SetProposalKey(
				b.ServiceKey().Address,
				b.ServiceKey().Index,
				b.ServiceKey().SequenceNumber,
			).
			SetPayer(b.ServiceKey().Address).
			AddAuthorizer(leofyCoinAddr)

		_ = tx.AddArgument(cadence.NewDictionary(recipientPairs))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				leofyCoinAddr,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				leofyCoinSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result, err := b.ExecuteScript(
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(leofyCoinAddr)),
			},
		)
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance := result.Value
		assert.Equal(t, CadenceUFix64("400.0"), balance)

		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result, err = b.ExecuteScript(
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		require.NoError(t, err)
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
		balance = result.Value
		assert.Equal(t, CadenceUFix64("600.0"), balance)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)*/
	})

}

func TestVaultDestroy(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	leofyCoinAccountKey, leofyCoinSigner := accountKeys.NewWithSigner()
	_, _, _, fungibleAddr, leofyCoinAddr := DeployContracts(b, t, []*flow.AccountKey{leofyCoinAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// Create Vault for JoshAddress
	CreateTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner)

	// Transfer 300 Tokens to JoshAddress
	TransferTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, leofyCoinSigner, joshAddress, CadenceUFix64("300.0"), false)

	t.Run("Should subtract tokens from supply when they are destroyed", func(t *testing.T) {
		script := templates.GenerateDestroyVaultScript(fungibleAddr, leofyCoinAddr)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, leofyCoinAddr)

		amount, _ := cadence.NewUFix64("100.00")

		tx.AddArgument(amount)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, leofyCoinAddr},
			[]crypto.Signer{b.ServiceKey().Signer(), leofyCoinSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(leofyCoinAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("600.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("900.0"), supply)
	})

	t.Run("Should subtract tokens from supply when they are destroyed by a different account", func(t *testing.T) {
		script := templates.GenerateDestroyVaultScript(fungibleAddr, leofyCoinAddr)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, joshAddress)

		amount, _ := cadence.NewUFix64("100.00")

		tx.AddArgument(amount)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress},
			[]crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)
		assert.Equal(t, CadenceUFix64("200.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("800.0"), supply)
	})

}

func TestMinting(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	leofyCoinAccountKey, leofyCoinSigner := accountKeys.NewWithSigner()
	_, _, _, fungibleAddr, leofyCoinAddr := DeployContracts(b, t, []*flow.AccountKey{leofyCoinAccountKey})

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	CreateTokenTransaction(t, b, fungibleAddr, leofyCoinAddr, joshAddress, joshSigner)
	// then deploy the tokens to an account

	t.Run("Shouldn't be able to mint zero tokens", func(t *testing.T) {
		script := templates.GenerateMintTokensScript(fungibleAddr, leofyCoinAddr)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, leofyCoinAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(CadenceUFix64("0.0"))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				leofyCoinAddr,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				leofyCoinSigner,
			},
			true,
		)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(leofyCoinAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("1000.0"), result)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("0.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1000.0"), supply)
	})

	t.Run("Should mint tokens, deposit, and update balance and total supply", func(t *testing.T) {
		script := templates.GenerateMintTokensScript(fungibleAddr, leofyCoinAddr)
		tx := createTxWithTemplateAndAuthorizer(
			b, script, leofyCoinAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(CadenceUFix64("50.0"))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				leofyCoinAddr,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				leofyCoinSigner,
			},
			false,
		)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result := executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(leofyCoinAddr)),
			},
		)

		assert.Equal(t, CadenceUFix64("1000.0"), result)

		// Assert that the vaults' balances are correct
		script = templates.GenerateInspectVaultScript(fungibleAddr, leofyCoinAddr)
		result = executeScriptAndCheck(t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.Address(joshAddress)),
			},
		)

		assert.Equal(t, CadenceUFix64("50.0"), result)

		script = templates.GenerateInspectSupplyScript(fungibleAddr, leofyCoinAddr)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, CadenceUFix64("1050.0"), supply)
	})
}
