import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transaction transfers a LeofyNFT Item from one account to another.

transaction(recipient: Address, itemID: UInt64, withdrawID: UInt64) {
    // local variable for storing the item collection reference
    let itemCollection: &LeofyNFT.ItemCollection

    prepare(signer: AuthAccount) {
        // borrow a reference to the ItemCollection resource in storage
        self.itemCollection = signer.borrow<&LeofyNFT.ItemCollection>(from: LeofyNFT.ItemStoragePath)
            ?? panic("Could not borrow a reference to the Item Collection")
    }

    execute{
        if self.itemCollection.items[itemID] != nil {
            let itemRef = &self.itemCollection.items[itemID] as &LeofyNFT.Item
            let collection = &itemRef.NFTsCollection as &LeofyNFT.Collection

            // borrow a public reference to the receivers collection
            let recipient = getAccount(recipient)
            let depositRef = recipient
                .getCapability(LeofyNFT.CollectionPublicPath)
                .borrow<&{NonFungibleToken.CollectionPublic}>()
                ?? panic("Could not borrow a reference to the receiver's collection")

            let nft <- collection.withdraw(withdrawID: withdrawID)
            
            //Deposit the NFT in the recipient's collection
            depositRef.deposit(token: <-nft)
        }
        else{
            panic("Item not exists")
        }
    }
}

