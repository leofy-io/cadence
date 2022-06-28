import LeofyMarketPlace from "../../contracts/LeofyMarketPlace.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"
import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import FungibleToken from "../../contracts/standard/FungibleToken.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"


transaction(id: UInt64, address: Address) {
    execute {
        let account = getAccount(address)

        let marketplaceCollectionRef = account.getCapability(LeofyMarketPlace.CollectionPublicPath)
        .borrow<&{LeofyMarketPlace.MarketplaceCollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")

        marketplaceCollectionRef.settleAuction(id);
    }
}