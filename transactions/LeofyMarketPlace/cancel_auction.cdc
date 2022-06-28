import LeofyMarketPlace from "../../contracts/LeofyMarketPlace.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"
import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import FungibleToken from "../../contracts/standard/FungibleToken.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"

// This transction uses the ItemCollection resource to mint a new NFT inside an Item.
//
// It must be run with the account that has the minter resource
// stored at path /storage/LeofyItemCollection.

transaction(id: UInt64) {
    
    // local variable for storing the item collection reference
    let marketplaceCollection: &LeofyMarketPlace.MarketplaceCollection

    prepare(signer: AuthAccount) {
        //self.depositRef = signer
        //        .getCapability<&{LeofyNFT.CollectionPublic}>(LeofyNFT.CollectionPublicPath)
        //        ?? panic("Could not borrow a reference to the receiver's collection")

       // borrow a reference to the ItemCollection resource in storage
       self.marketplaceCollection = signer.borrow<&LeofyMarketPlace.MarketplaceCollection>(from: LeofyMarketPlace.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the Auction Collection")

    }

    execute {
        if self.marketplaceCollection != nil {
            self.marketplaceCollection.cancelAuction(id)
        }
        else{
            panic("Auction Collection not exists")
        }
    }
}
