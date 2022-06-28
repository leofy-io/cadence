import NonFungibleToken from "../../../contracts/standard/NonFungibleToken.cdc"
import LeofyMarketPlace from "../../../contracts/LeofyMarketPlace.cdc"

pub fun main(address: Address, id: UInt64): AnyStruct? {
  let account = getAccount(address)

  let marketplaceCollectionRef = account.getCapability(LeofyMarketPlace.CollectionPublicPath)
        .borrow<&{LeofyMarketPlace.MarketplaceCollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
  return marketplaceCollectionRef.getMarketplaceStatus(id);
}