import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import LeofyNFT from "../../../contracts/LeofyNFT.cdc"

pub fun main(address: Address, id: UInt64): AnyStruct? {
  let account = getAccount(address)

    let collectionRef = account.getCapability(/public/LeofyNFTCollection)
        .borrow<&{LeofyNFT.LeofyCollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef.borrowNFT(id: id)
}