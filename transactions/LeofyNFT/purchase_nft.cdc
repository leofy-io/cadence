import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import FungibleToken from "../../contracts/standard/FungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"

// This transaction borrow the ItemPublic from ItemCollection selected with itemID parameter
// Withdrawal an amount for LeofyCoin (Fungible) Vault with the same value than Item.Price

transaction(itemID: UInt64) {
    
    // The Vault resource that holds the tokens that are being transferred
    let sentVault: @FungibleToken.Vault
    let item: &LeofyNFT.Item{LeofyNFT.ItemPublic}
    let depositRef: &AnyResource{NonFungibleToken.CollectionPublic}

    prepare(signer: AuthAccount) {
        
        // Get a reference to the signer's stored vault
        let vaultRef = signer.borrow<&FungibleToken.Vault>(from: LeofyCoin.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

        // borrow a public reference to the receivers collection
        self.depositRef = signer
            .getCapability(LeofyNFT.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not borrow a reference to the receiver's collection")    

        let itemCollection = LeofyNFT.getItemCollectionPublic()
        self.item = itemCollection.borrowItem(itemID: itemID)!
        
        // Withdraw tokens from the signer's stored vault
        self.sentVault <- vaultRef.withdraw(amount: self.item.price)
    }

    execute {
        let nft <- self.item.purchase(payment: <- self.sentVault)
        // Deposit the NFT in the recipient's collection
        self.depositRef.deposit(token: <-nft)
    }
}