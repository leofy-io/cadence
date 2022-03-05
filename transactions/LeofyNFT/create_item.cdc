import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transction uses the Item resource to mint a new NFT.
// 
// Item resource has to be borrow from admin resource
// stored at path /storage/LeofyNFTMinter.

transaction(metadata: {String: String}) {
    
    // local variable for storing the minter reference
    let admin: &LeofyNFT.Admin

    prepare(signer: AuthAccount) {
        
        // borrow a reference to the NFTMinter resource in storagele
        self.admin = signer.borrow<&LeofyNFT.Admin>(from: LeofyNFT.AdminStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        self.admin.createItem(metadata: metadata)
    }
}

/*
flow transactions send ./cadence/transactions/LeofyNFT/create_item.cdc \
--args-json '[{"type": "Dictionary", "value": [
    {
      "key": {
        "type": "String",
        "value": "author"
      },
      "value": {
        "type": "String",
        "value": "Leofy"
      }
    },
    {
      "key": {
        "type": "String",
        "value": "name"
      },
      "value": {
        "type": "String",
        "value": "Our First NFT"
      }
    },
    {
      "key": {
        "type": "String",
        "value": "thumbnail"
      },
      "value": {
        "type": "String",
        "value": "https://leofy.io/leofy-logo-y-2.svg"
      }
    }
]}]' --network testnet --signer leofynft-testnet-account
*/
