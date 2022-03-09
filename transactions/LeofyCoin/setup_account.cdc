import FungibleToken from "../../contracts/standard/FungibleToken.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"

// This transaction is a template for a transaction
// to add a Vault resource to their account
// so that they can use the LeofyCoin


transaction {

    prepare(signer: AuthAccount) {

        if signer.borrow<&LeofyCoin.Vault>(from: LeofyCoin.VaultStoragePath) == nil {
            // Create a new LeofyCoin Vault and put it in storage
            signer.save(<-LeofyCoin.createEmptyVault(), to: LeofyCoin.VaultStoragePath)

            // Create a public capability to the Vault that only exposes
            // the deposit function through the Receiver interface
            signer.link<&LeofyCoin.Vault{FungibleToken.Receiver}>(
                LeofyCoin.ReceiverPublicPath,
                target: LeofyCoin.VaultStoragePath
            )

            // Create a public capability to the Vault that only exposes
            // the balance field through the Balance interface
            signer.link<&LeofyCoin.Vault{FungibleToken.Balance}>(
                LeofyCoin.BalancePublicPath,
                target: LeofyCoin.VaultStoragePath
            )
        }
    }
}