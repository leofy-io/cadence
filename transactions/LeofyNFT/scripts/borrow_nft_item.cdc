import NonFungibleToken from "../../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../../contracts/LeofyNFT.cdc"

// This script borrows an NFT from a collection
pub fun main(itemID: UInt64, id: UInt64): &NonFungibleToken.NFT {
    let itemCollection = LeofyNFT.getItemCollectionPublic()
    let item = itemCollection.borrowItem(itemID: itemID)
    // Borrow a reference to a specific NFT in the collection
    let nft = item.borrowCollection().borrowNFT(id: id)

    return nft
}