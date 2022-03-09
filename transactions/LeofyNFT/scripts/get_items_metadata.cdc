import LeofyNFT from "../../../contracts/LeofyNFT.cdc"

// This scripts returns the number of Items currently in existence.

pub fun main(itemID: UInt64): {String: String} { 
    let itemCollection = LeofyNFT.getItemCollectionPublic()
    let item = itemCollection.borrowItem(itemID: itemID)
    //let NFTItemCollection = item.borrowCollection()
    
    //destroy NFTItemCollection.withdraw(withdrawID: 0)

    return item.metadata


}
