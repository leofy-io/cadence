import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transaction is a template for a transaction
// to add a Vault resource to their account
// so that they can use the Collection


transaction {

    prepare(signer: AuthAccount) {

        let collection <- LeofyNFT.createEmptyCollection()
        signer.save(<-collection, to: LeofyNFT.CollectionStoragePath)

        signer.link<&{LeofyNFT.LeofyCollectionPublic}>(
            LeofyNFT.CollectionPublicPath,
            target: LeofyNFT.CollectionStoragePath
        )
    }

}