import LeofyMarketPlace from "../../contracts/LeofyMarketPlace.cdc"

// This transaction is a template for a transaction
// to add a Auction Collection resource to their account
// so that they can use it


transaction {

    prepare(signer: AuthAccount) {
        // Return early if the account already has a collection
        if signer.borrow<&LeofyMarketPlace.MarketplaceCollection>(from: LeofyMarketPlace.CollectionStoragePath) != nil {
            return
        }

        // Create a new empty collection
        let marketplaceCollection <- LeofyMarketPlace.createEmptyCollection()

        // save it to the account
        signer.save(<-marketplaceCollection, to: LeofyMarketPlace.CollectionStoragePath)

        signer.link<&LeofyMarketPlace.MarketplaceCollection{LeofyMarketPlace.MarketplaceCollectionPublic}>(
            LeofyMarketPlace.CollectionPublicPath,
            target: LeofyMarketPlace.CollectionStoragePath
        )
    }

}