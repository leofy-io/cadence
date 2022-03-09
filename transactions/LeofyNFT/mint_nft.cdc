import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transction uses the ItemCollection resource to mint a new NFT inside an Item.
//
// It must be run with the account that has the minter resource
// stored at path /storage/LeofyItemCollection.

transaction(itemID: UInt64) {
    
    // local variable for storing the item collection reference
    let itemCollection: &LeofyNFT.ItemCollection

    prepare(signer: AuthAccount) {
       // borrow a reference to the ItemCollection resource in storage
        self.itemCollection = signer.borrow<&LeofyNFT.ItemCollection>(from: LeofyNFT.ItemStoragePath)
            ?? panic("Could not borrow a reference to the Item Collection")
    }

    execute {
        if self.itemCollection.items[itemID] != nil {
            let itemRef = &self.itemCollection.items[itemID] as &LeofyNFT.Item
            itemRef.mintNFT()
        }
        else{
            panic("Item not exists")
        }
    }
}
