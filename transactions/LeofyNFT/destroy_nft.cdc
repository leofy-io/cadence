import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

transaction(id: UInt64) {
    prepare(signer: AuthAccount) {
        let collectionRef = signer.borrow<&LeofyNFT.Collection>(from: LeofyNFT.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the owner's collection")

        // withdraw the NFT from the owner's collection
        let nft <- collectionRef.withdraw(withdrawID: id)

        destroy nft
    }
}