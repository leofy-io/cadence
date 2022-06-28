import LeofyMarketPlace from "../../contracts/LeofyMarketPlace.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"
import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import FungibleToken from "../../contracts/standard/FungibleToken.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"


transaction(id: UInt64, address: Address, amount: UFix64) {

    // The Vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault

    // The Vault capability to FT
    let ftReceiverCap: Capability<&{FungibleToken.Receiver}>

    // The NFT Collection capability to reeive NFT
    let nftReceiverCap: Capability<&{LeofyNFT.LeofyCollectionPublic}>

    prepare(signer: AuthAccount) {

        // Get a reference to the signer's stored vault
        let vaultRef = signer.borrow<&FungibleToken.Vault>(from: LeofyCoin.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        // Withdraw tokens from the signer's stored vault
        self.sentVault <- vaultRef.withdraw(amount: amount)

        self.ftReceiverCap = signer.getCapability<&AnyResource{FungibleToken.Receiver}>(LeofyCoin.ReceiverPublicPath)

        self.nftReceiverCap = signer
                .getCapability<&{LeofyNFT.LeofyCollectionPublic}>(LeofyNFT.CollectionPublicPath)
    }

    execute {
        let account = getAccount(address)

        let marketplaceCollectionRef = account.getCapability(LeofyMarketPlace.CollectionPublicPath)
        .borrow<&{LeofyMarketPlace.MarketplaceCollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")

        marketplaceCollectionRef.placePurchase(
            id: id,
            payment: <- self.sentVault,
            collectionCap: self.nftReceiverCap
        );
    }
}