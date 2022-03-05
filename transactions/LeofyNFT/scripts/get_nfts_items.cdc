import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import LeofyNFT from "../../../contracts/LeofyNFT.cdc"

pub fun main(address: Address): AnyStruct? {
  let account = getAccount(address)

    let collectionRef = account.getCapability(LeofyNFT.CollectionPublicPath)
        .borrow<&{LeofyNFT.LeofyCollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef
}