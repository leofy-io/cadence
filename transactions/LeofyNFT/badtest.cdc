import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

transaction(itemID: UInt64) {
    
    // local variable for storing the item collection reference
    //let itemCollection: &LeofyNFT.ItemCollection
    let depositRef: &AnyResource{NonFungibleToken.CollectionPublic}

    prepare(signer: AuthAccount) {
       // borrow a reference to the ItemCollection resource in storage
        //self.itemCollection = signer.borrow<&LeofyNFT.ItemCollection>(from: LeofyNFT.ItemStoragePath)
        //    ?? panic("Could not borrow a reference to the Item Collection")
        self.depositRef = signer
            .getCapability(LeofyNFT.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not borrow a reference to the receiver's collection") 
    }

    execute {
        
        let itemCollection = LeofyNFT.getItemCollectionPublic()
        let item = itemCollection.borrowItem(itemID: itemID)

        let itemRef = item as &LeofyNFT.Item

        log(itemRef.mintNFT())


    }
}
