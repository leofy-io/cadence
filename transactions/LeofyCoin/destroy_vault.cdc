import FungibleToken from "../../contracts/standard/FungibleToken.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"

transaction(amount: UFix64) {
    prepare(signer: AuthAccount) {
        let vaultRef = signer.borrow<&FungibleToken.Vault>(from: LeofyCoin.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        // Withdraw tokens from the signer's stored vault
        let sentVault <- vaultRef.withdraw(amount: amount)

		destroy sentVault
	}
}
 