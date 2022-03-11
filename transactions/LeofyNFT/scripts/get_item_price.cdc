import NonFungibleToken from "../../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../../contracts/LeofyNFT.cdc"

pub fun main(itemID: UInt64): UFix64 {
    let itemCollection = LeofyNFT.getItemCollectionPublic()
    let item = itemCollection.borrowItem(itemID: itemID)
    // Borrow a reference to a specific NFT in the collection
    return item.price;
}