import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transction uses the Item resource to mint a new Item.
// 
// Item resource has to be borrow from admin resource
// stored at path /storage/LeofyNFTMinter.

transaction(itemID: UInt64, price: UFix64) {
    
    // local variable for storing the item collection reference
    let itemCollection: &LeofyNFT.ItemCollection

    prepare(signer: AuthAccount) {
        
        // borrow a reference to the ItemCollection resource in storage
        self.itemCollection = signer.borrow<&LeofyNFT.ItemCollection>(from: LeofyNFT.ItemStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        let item = (&self.itemCollection.items[itemID] as &LeofyNFT.Item?)!
		item.setPrice(price: price)
		
    }
}