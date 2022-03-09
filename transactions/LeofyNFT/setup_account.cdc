import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transaction is a template for a transaction
// to add a Vault resource to their account
// so that they can use the Collection


transaction {

    prepare(signer: AuthAccount) {
        // Return early if the account already has a collection
        if signer.borrow<&LeofyNFT.Collection>(from: LeofyNFT.CollectionStoragePath) != nil {
            return
        }

        // Create a new empty collection
        let collection <- LeofyNFT.createEmptyCollection()

        // save it to the account
        signer.save(<-collection, to: LeofyNFT.CollectionStoragePath)

        // create a public capability for the collection
        signer.link<&{NonFungibleToken.CollectionPublic, LeofyNFT.LeofyCollectionPublic}>(
            LeofyNFT.CollectionPublicPath,
            target: LeofyNFT.CollectionStoragePath
        )
    }

}