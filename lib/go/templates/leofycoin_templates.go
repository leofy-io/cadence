package templates

import (
	_ "github.com/kevinburke/go-bindata"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	transferTokensFilename       = "LeofyCoin/transfer_tokens.cdc"
	transferManyAccountsFilename = "LeofyCoin/transfer_many_accounts.cdc"
	setupAccountFilename         = "LeofyCoin/setup_account.cdc"
	mintTokensFilename           = "LeofyCoin/mint_tokens.cdc"
	destroyVaultFilename         = "LeofyCoin/destroy_vault.cdc"
	burnTokensFilename           = "LeofyCoin/burn_tokens.cdc"

	scriptsPath         = "LeofyCoin/scripts/"
	readBalanceFilename = "get_balance.cdc"
	readSupplyFilename  = "get_supply.cdc"
)

// -----------------------------------------------------------------------
// LeofyCoin Transactions
// -----------------------------------------------------------------------

// GenerateCreateTokenScript creates a script that instantiates
// a new Vault instance and stores it in storage.
// balance is an argument to the Vault constructor.
// The Vault must have been deployed already.
func GenerateCreateTokenScript(fungibleAddr, tokenAddr flow.Address) []byte {
	code := assets.MustAssetString(setupAccountFilename)
	return replaceAddressesLeofyCoin(code, fungibleAddr, tokenAddr)
}

// GenerateTransferVaultScript creates a script that withdraws an tokens from an account
// and deposits it to another account's vault
func GenerateTransferVaultScript(fungibleAddr, tokenAddr flow.Address) []byte {
	code := assets.MustAssetString(transferTokensFilename)
	return replaceAddressesLeofyCoin(code, fungibleAddr, tokenAddr)
}

// GenerateDestroyVaultScript creates a script that withdraws
// tokens from a vault and destroys the tokens
func GenerateDestroyVaultScript(fungibleAddr, tokenAddr flow.Address) []byte {
	code := assets.MustAssetString(destroyVaultFilename)
	return replaceAddressesLeofyCoin(code, fungibleAddr, tokenAddr)
}

// GenerateMintTokensScript creates a script that uses the admin resource
// to mint new tokens and deposit them in a Vault
func GenerateMintTokensScript(fungibleAddr, tokenAddr flow.Address) []byte {
	code := assets.MustAssetString(mintTokensFilename)
	return replaceAddressesLeofyCoin(code, fungibleAddr, tokenAddr)
}

// GenerateBurnTokensScript creates a script that uses the admin resource
// to destroy tokens and deposit them in a Vault
func GenerateBurnTokensScript(fungibleAddr, tokenAddr flow.Address) []byte {
	code := assets.MustAssetString(burnTokensFilename)
	return replaceAddressesLeofyCoin(code, fungibleAddr, tokenAddr)
}

// -----------------------------------------------------------------------
// LeofyCoin Scripts
// -----------------------------------------------------------------------

// GenerateInspectVaultScript creates a script that retrieves a
// Vault from the array in storage and makes assertions about
// its balance. If these assertions fail, the script panics.
func GenerateInspectVaultScript(fungibleAddr, tokenAddr flow.Address) []byte {
	code := assets.MustAssetString(scriptsPath + readBalanceFilename)
	return replaceAddressesLeofyCoin(code, fungibleAddr, tokenAddr)
}

// GenerateInspectSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectSupplyScript(fungibleAddr, tokenAddr flow.Address) []byte {
	code := assets.MustAssetString(scriptsPath + readSupplyFilename)
	return replaceAddressesLeofyCoin(code, fungibleAddr, tokenAddr)
}
