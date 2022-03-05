import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transction uses the NFTMinter resource to mint a new NFT.
//
// It must be run with the account that has the minter resource
// stored at path /storage/LeofyNFTMinter.

transaction(
    recipient: Address,
    itemID: UInt64
) {
    
    // local variable for storing the minter reference
    let admin: &LeofyNFT.Admin

    prepare(signer: AuthAccount) {

        // borrow a reference to the NFTMinter resource in storagele
        self.admin = signer.borrow<&LeofyNFT.Admin>(from: LeofyNFT.AdminStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        // Borrow the recipient's public NFT collection reference
        let receiver = getAccount(recipient)
            .getCapability(LeofyNFT.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        let itemRef = self.admin.borrowItem(itemID: itemID)

        // Mint the NFT and deposit it to the recipient's collection
        itemRef.mintNFT(
            recipient: receiver
        )
    }
}
