import LeofyMarketPlace from "../../contracts/LeofyMarketPlace.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"
import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import FungibleToken from "../../contracts/standard/FungibleToken.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"

// This transction uses the ItemCollection resource to mint a new NFT inside an Item.
//
// It must be run with the account that has the minter resource
// stored at path /storage/LeofyItemCollection.

transaction(id: UInt64, auctionStartTime: UFix64, auctionLength: UFix64, startPrice: UFix64, purchasePrice: UFix64) {
    
    // local variable for storing the item collection reference
    let marketplaceCollection: &LeofyMarketPlace.MarketplaceCollection
    let itemCollection: &LeofyNFT.Collection
    let ownerDepositCap: Capability<&AnyResource{LeofyNFT.LeofyCollectionPublic}>
    let ownerVaultCap: Capability<&AnyResource{FungibleToken.Receiver}>

    prepare(signer: AuthAccount) {
        //self.depositRef = signer
        //        .getCapability<&{LeofyNFT.CollectionPublic}>(LeofyNFT.CollectionPublicPath)
        //        ?? panic("Could not borrow a reference to the receiver's collection")

       // borrow a reference to the ItemCollection resource in storage
        self.marketplaceCollection = signer.borrow<&LeofyMarketPlace.MarketplaceCollection>(from: LeofyMarketPlace.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the Auction Collection")

        self.itemCollection = signer
            .borrow<&LeofyNFT.Collection>(from: LeofyNFT.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the owner's collection")

        self.ownerDepositCap = signer
            .getCapability<&AnyResource{LeofyNFT.LeofyCollectionPublic}>(LeofyNFT.CollectionPublicPath)
        
        self.ownerVaultCap = signer
            .getCapability<&AnyResource{FungibleToken.Receiver}>(LeofyCoin.ReceiverPublicPath)


    }

    execute {
        if self.marketplaceCollection != nil {
            let nft <- self.itemCollection.withdraw(withdrawID: id) as! @LeofyNFT.NFT

            self.marketplaceCollection.sellItem(
                token: <-nft,
                auctionStartTime: getCurrentBlock().timestamp,
                auctionLength: auctionLength,
                startPrice: startPrice,
                ownerCollectionCap: self.ownerDepositCap,
                ownerVaultCap: self.ownerVaultCap,
                purchasePrice: purchasePrice
            )
        }
        else{
            panic("Auction Collection not exists")
        }
    }
}
