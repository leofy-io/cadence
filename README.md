# Leofy

## Introduction

This repository contains the smart contracts and transactions that implement the core functionality Leofy.

The smart contracts are written in Cadence, a new resource oriented
smart contract programming language designed for the Flow Blockchain.

### What is Leofy

Leofy is the first representation agency in the metaverse and exploitation of NFTs for main Artists on Spain and LATAM.

First Spanish-speaking portal for the commercialization of NFT on Flow.

### What is Flow?

Flow is a new blockchain for open worlds. Read more about it [here](https://www.onflow.org/).

### What is Cadence?

Cadence is a new Resource-oriented programming language 
for developing smart contracts for the Flow Blockchain.
Read more about it [here](https://www.docs.onflow.org)

We recommend that anyone who is reading this should have already
completed the [Cadence Tutorials](https://docs.onflow.org/cadence) 
so they can build a basic understanding of the programming language.

Resource-oriented programming, and by extension Cadence, 
is the perfect programming environment for Non-Fungible Tokens (NFTs), because users are able
to store their NFT objects directly in their accounts and transact
peer-to-peer. Please see the [blog post about resources](https://medium.com/dapperlabs/resource-oriented-programming-bee4d69c8f8e)
to understand why they are perfect for digital assets like LeofyNFTs or NBA Top Shot Moments or anothers recetly daps on the market. 

### Contributing

If you see an issue with the code for the contracts, the transactions, scripts, documentation, or anything else, please do not hesitate to make an issue or a pull request with your desired changes. This is an open source project and we welcome all assistance from the community!

## LeofyNFT Contract Addresses

`LeofyNFT.cdc`: This is the main Leofy smart contract that defines
the core functionality of the NFT.

| Network | Contract Address     |
|---------|----------------------|
| Testnet | `0xd3af09bdd3c94553` |
| Mainnet | `TBA` |

`LeofyCoin.cdc`: This is the Leofy Fungible Token contract that allows will allow purchase NFT's

| Network | Contract Address     |
|---------|----------------------|
| Testnet | `0xd3af09bdd3c94553` |
| Mainnet | `TBA` |

TBA: Soon we will launch a Marketplace Contract that will allow users to sell their NFTs.

### Non Fungible Token Standard

The LeofyNFT contract utilize the [Flow NFT standard](https://github.com/onflow/flow-nft)
which is equivalent to ERC-721 or ERC-1155 on Ethereum. If you want to build an NFT contract,
please familiarize yourself with the Flow NFT standard before starting and make sure you utilize it 
in your project in order to be interoperable with other tokens and contracts that implement the standard.

### Fungible Token Standard

The LeofyCoin contract utilize the [Flow FT standard](https://github.com/onflow/flow-ft), which is used to make a purchase of NFT.

## Directory Structure

The directories here are organized into contracts, transactions, scripts and lib/go testing library.

Contracts contain the source code for the Leofy contracts that are deployed to Flow.

Scripts contain read-only transactions to get information about
the state of someones Collection or about the state of the Leofy contract.

Transactions contain the transactions that admins and users can use to perform actions in the smart contract like creating Items, mint NFT's for an specific Item, transfering NFT's, mint Leofy Fungible Token (LeofyCoin), transfer LeofyCoin, etc. 

 - `contracts/` : Where the Leofy related smart contracts live.
 - `transactions/` : This directory contains all the transactions and scripts that are associated with the Leofy smart contracts.
 - `transactions/scripts/`  : This contains all the read-only Cadence scripts that are used to read information from the smart contract
 or from a resource in account storage.
 - `lib/` : This directory contains packages for specific programming languages to be able to read copies of the Leofy smart contracts, transaction templates, and scripts.  Also contains automated tests written in those languages. Currently,
 Go is the only language that is supported, but we are hoping to add javascript
 and other languages soon. See the README in `lib/go/` for more information
 about how to use the Go packages.

## LeofyNFT Contract Overview

Each Leofy NFT represents an Item (can be an Picture, Song, Video, Concert, or any Kind of art represented on NFT digital Asset)

Multiple NFTs can be minted inside the same Item and each receives a serial number.

Each NFT is a resource object 
with roughly the following structure:

```cadence
pub resource NFT: NonFungibleToken.INFT, MetadataViews.Resolver  {
    pub let id: UInt64
    pub let itemID: UInt64
    pub let serialNumber: UInt32

    ...
}
```

The other types that are defined in `LeofyNFT` are as follows:

 - `ItemCollection`: A resource that contains a dictionary with each Item created. This resource are stored on the main account address 'storage' to be used to Create New Items. It's also linked to a public capability asociated to a `ItemCollectionPublic` interface to have the capability to borrow an Item (With an `ItemPublic` interface reference).
 - `Item`: An resource that contains variable data for an Item (like number of NFT's minted on that Item, or price), name, thumbnail or descritpion, and a NFTCollection minted for that Item. Also contains the functionality for Mint NFT's on that Item (only for admins), or borrow the NFTCollection of that Item.
 - `NFT`: A resource type that is the NFT that represents the Item a user owns. It stores its unique ID and other metadata. This is the collectible object that the users can store in their accounts.
 - `Collection`: Similar to the `NFTCollection` resource from the NFT
    example, this resource is a repository for a user's NFTs.  Users can
    withdraw and deposit from this collection and get information about the 
    contained NFTs.

Metadata structs associated with Items are stored in the main smart contract
and can be queried by anyone. 

For example, If players wanted to find out the 
author of the Item that the player has on their Collection, they 
would call a public function in the `Leofy` smart contract 
called `getItemMetaDataByField` inside the `Item` resource.

The power to create new Items, and mint NFT's inside each Item
rests with the owner of the `ItemCollection` and `Item` resource. 
External users only can borrow by reference their public interface without the admin functions.

Each `Item` has a price for all their NFTCollection. Users can purchase an NFT with a LeofyCoin Vault. 

Once a user owns a LeofyNFT object, that LeofyNFT is stored directly 
in their account storage via their `Collection` object. The collection object
contains a dictionary that stores the LeofyNFT and gives utility functions
to move them in and out and to read data about the collection and its LeofyNFT.


## How to Run Transactions Against the Leofy Contracts
This repository contains sample transactions that can be executed against the TopShot contract either via Flow CLI or using VSCode. This section will describe how to create setup a public Collection  on the Flow emulator.

#### Send Transaction with Flow CLI
1. Install the [Flow CLI and emulator](https://docs.onflow.org/flow-cli/install/)
2. Initialize the flow emulator configuration.  
`flow emulator --init`
3. [Configure the contracts & deployment section](https://docs.onflow.org/flow-cli/configuration/) of the initialized flow.json file. 
4. Start the emulator.  
`flow emulator`
5. Deploy the NonFungibleToken & TopShot contracts to the flow emulator.  
`flow project deploy --network=emulator`
6. Use the Flow CLI to execute transactions against the emulator. This transaction creates a new set on the flow emulator called "new set name".   
`flow transactions send ./transactions/LeofyNFT/setup_account.cdc"`

## How to run the automated tests for the contracts

See the `lib/go` README for instructions about how to run the automated tests.

## Instructions create Fungible Token (LeofyCoin) Vault and send transfer

A common order of creating Vault and transfer fungibles.

1. Creating and Setup Vault on new account `transactions/LeofyCoin/setup_account.cdc`.
2. Mint tokens with Admin or Service Account `transactions/LeofyCoin/mint_tokens.cdc`.
3. Transfer Tokens to another account `transactions/LeofyCoin/transfer_tokens.cdc`.
4. Destroy Vault 
   `transactions/LeofyCoin/destroy_vault.cdc`.

You can also see the scripts in `transactions/LeofyCoin/scripts` to see how information
can be read from the real LeofyCoin contract deployed.

## Instructions create Non Fungible Token (LeofyNFT) Collection, mint Items and NFTs

A common order of creating Items and mint NFT's inside that Items.

1. Creating and Setup Collection on new account `transactions/LeofyNFT/setup_account.cdc`.
2. Mint / create a new Item with Admin Account `transactions/LeofyNFT/create_item.cdc`.
3. Mint one NFT for an Item with Admin or Service Account `transactions/LeofyNFT/mint_nft.cdc`.
4. Batch Mint multiple NFTs for an Item with Admin or Service Account `transactions/LeofyNFT/batch_mint_nft.cdc`.
5. Transfer Tokens from Item to public account Collection (Only Admin) `transactions/LeofyNFT/transfer_nft_item.cdc`.
6. Transfer Tokens from public account Collection to another account (NFT Owner) `transactions/LeofyNFT/transfer_nft_account.cdc`.
7. Purchase NFT from Item Collection
   `transactions/LeofyNFT/purchase_nft.cdc`.

You can also see the scripts in `transactions/LeofyNFT/scripts` to see how information
can be read from the real LeofyCoin contract deployed.

### Leofy NFT Metadata

NFT metadata is represented in a flexible and modular way using the [standard proposed in FLIP-0636](https://github.com/onflow/flow/blob/master/flips/20210916-nft-metadata.md). The LeofyNFT contract implements the [`MetadataViews.Resolver`](https://github.com/onflow/flow-nft/blob/master/contracts/MetadataViews.cdc#L21) interface, which standardizes the display of LeofyNFT in accordance with FLIP-0636. The LeofyNFT contract also defines a custom view of moment play data called TopShotMomentMetadataView.


## Leofy Marketplace

Coming Soon. 

## License 

The works in these folders 
/leofy-io/cadence/blob/master/contracts/LeofyCoin.cdc 
/leofy-io/cadence/blob/master/contracts/LeofyNFT.cdc 

are under the Unlicense
https://github.com/leofy-oo/cadence/blob/master/LICENSE

## Inspiration 

- NBA Top Shot: https://nbatopshot.com/
- Viv3:  https://viv3.com/











 