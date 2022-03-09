import LeofyNFT from "../../../contracts/LeofyNFT.cdc"
import MetadataViews from "../../../contracts/standard/MetadataViews.cdc"

pub fun main(address: Address, id: UInt64): LeofyNFT.LeofyNFTMetadataView {
    let account = getAccount(address)

    let collectionRef = account
        .getCapability(LeofyNFT.CollectionPublicPath)
        .borrow<&{LeofyNFT.LeofyCollectionPublic}>()
        ?? panic("Could not borrow a reference to the collection")

    let nft = collectionRef.borrowLeofyNFT(id: id)!
    
    // Get the Top Shot specific metadata for this NFT
    let view = nft.resolveView(Type<LeofyNFT.LeofyNFTMetadataView>())!

    let metadata = view as! LeofyNFT.LeofyNFTMetadataView
    
    return metadata
}