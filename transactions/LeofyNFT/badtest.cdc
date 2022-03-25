import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

transaction(itemID: UInt64) {
    
    execute {
        let itemCollection = LeofyNFT.getItemCollectionPublic()
        let item = itemCollection.borrowItem(itemID: itemID)


        
        //let NFTItemCollection = item.borrowCollection()
        
        //destroy NFTItemCollection.withdraw(withdrawID: 0)
        let metadata = item.getMetadata()
        metadata["author"] = "Item name changed";
        log("START TEST")
        log(itemCollection.getItemMetaDataByField(itemID: itemID, field: "author"))
        let metadata2 = item.getMetadata()
        log(metadata2)


    }
}
