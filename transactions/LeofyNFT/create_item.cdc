import NonFungibleToken from "../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../contracts/LeofyNFT.cdc"

// This transction uses the Item resource to mint a new Item.
// 
// Item resource has to be borrow from admin resource
// stored at path /storage/LeofyNFTMinter.

transaction(metadata: {String: String}, price: UFix64) {
    
    // local variable for storing the item collection reference
    let itemCollection: &LeofyNFT.ItemCollection

    prepare(signer: AuthAccount) {
        
        // borrow a reference to the ItemCollection resource in storage
        self.itemCollection = signer.borrow<&LeofyNFT.ItemCollection>(from: LeofyNFT.ItemStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        self.itemCollection.createItem(metadata: metadata, price: price)
    }
}

/*
flow transactions send ./cadence/transactions/LeofyNFT/create_item.cdc \
--args-json '[{
	"type": "Dictionary",
	"value": [{
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
	]
}, {
	"type": "UFix64",
	"value": "20.00"
}]' --network testnet --signer leofynft-testnet-account
*/
 