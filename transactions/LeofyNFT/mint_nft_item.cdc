import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transction uses the NFTMinter resource to mint a new NFT.
//
// It must be run with the account that has the minter resource
// stored at path /storage/LeofyNFTMinter.

transaction(recipient: Address, typeID: UInt64) {
    
    // local variable for storing the minter reference
    let minter: &LeofyNFT.NFTMinter

    prepare(signer: AuthAccount) {

        // borrow a reference to the NFTMinter resource in storagele
        self.minter = signer.borrow<&LeofyNFT.NFTMinter>(from: /storage/LeofyNFTMinter)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        // get the public account object for the recipient
        let recipient = getAccount(recipient)

        // borrow the recipient's public NFT collection reference
        let receiver = recipient
            .getCapability(/public/LeofyNFTCollection)!
            .borrow<&{LeofyNFT.LeofyCollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        let metadata =  {"typeID": typeID.toString()};

        // mint the NFT and deposit it to the recipient's collection
        self.minter.mintNFT(recipient: receiver, metadata: metadata)
    }
}
