/*
    Description: Central Smart Contract for Leofy

    This smart contract contains the core functionality for 
    Leofy, created by LEOFY DIGITAL S.L.

    The contract manages the data associated with all the items
    that are used as templates for the NFTs

    Then an Admin can create new Items. Items consist of a public struct that 
    contains public information about a item, and a private resource used
    to mint new NFT's linked to the Item.

    The admin resource has the power to do all of the important actions
    in the smart contract. When admins want to call functions in a Item,
    they call their borrowItem function to get a reference 
    to a item in the contract. Then, they can call functions on the item using that reference.
    
    When NFTs are minted, they are initialized with a ItemID and
    are returned by the minter.

    The contract also defines a Collection resource. This is an object that 
    every Leofy NFT owner will store in their account
    to manage their NFT collection.

    The main Leofy account will also have its own Moment collections
    it can use to hold its own moments that have not yet been sent to a user.

    Note: All state changing functions will panic if an invalid argument is
    provided or one of its pre-conditions or post conditions aren't met.
    Functions that don't modify state will simply return 0 or nil 
    and those cases need to be handled by the caller.

*/

import NonFungibleToken from "./NonFungibleToken.cdc"
import MetadataViews from "./MetadataViews.cdc"

pub contract LeofyNFT: NonFungibleToken {

    // -----------------------------------------------------------------------
    // Leofy contract Events
    // -----------------------------------------------------------------------

    // Emitted when the LeofyNFT contract is created
    pub event ContractInitialized()

    // Emitted when a new Item struct is created
    pub event ItemCreated(id: UInt64, metadata: {String:String})
    pub event SetCreated(id: UInt64, name: String)

    pub event Withdraw(id: UInt64, from: Address?)
    pub event Deposit(id: UInt64, to: Address?)
    pub event Minted(id: UInt64)

    // Named Paths
    //
    pub let CollectionStoragePath: StoragePath
    pub let CollectionPublicPath: PublicPath
    pub let AdminStoragePath: StoragePath

    // -----------------------------------------------------------------------
    // TopShot contract-level fields.
    // These contain actual values that are stored in the smart contract.
    // -----------------------------------------------------------------------

    // Variable size dictionary of Item structs
    access(self) var items: @{UInt64: Item}

    pub var totalSupply: UInt64
    pub var totalItemSupply: UInt64

    // -----------------------------------------------------------------------
    // LeofyNFT contract-level Composite Type definitions
    // -----------------------------------------------------------------------
    // These are just *definitions* for Types that this contract
    // and other accounts can use. These definitions do not contain
    // actual stored values, but an instance (or object) of one of these Types
    // can be created by this contract that contains stored values.
    // -----------------------------------------------------------------------

    
    // Item is a Resource that holds metadata associated 
    // with a specific Artist Item, like the picture from Artist John Doe
    //
    // Leofy NFTs will all reference a single item as the owner of
    // its metadata. 
    //
    pub resource Item {

        // The unique ID for the Item
        pub let itemID: UInt64

        // Stores all the metadata about the item as a string mapping
        // This is not the long term way NFT metadata will be stored. It's a temporary
        // construct while we figure out a better way to do metadata.
        //
        pub let metadata: {String: String}

        access(contract) var numberMinted: UInt32

        init(metadata: {String: String}) {
            pre {
                metadata.length != 0: "New Item metadata cannot be empty"
            }
            self.itemID = LeofyNFT.totalItemSupply
            self.metadata = metadata
            self.numberMinted = 0
        }

        pub fun mintNFT(
            recipient: &{NonFungibleToken.CollectionPublic}
        ) {
            // create a new NFT
            var newNFT <- create NFT(
                id: LeofyNFT.totalSupply,
                itemID: self.itemID,
                serialNumber: self.numberMinted + 1
            )

            // deposit it in the recipient's account using their reference
            recipient.deposit(token: <-newNFT)

            emit Minted(id: LeofyNFT.totalSupply)

            self.numberMinted =  self.numberMinted + 1
            LeofyNFT.totalSupply = LeofyNFT.totalSupply + 1
        }

        pub fun batchMintNFT(
            recipient: &{NonFungibleToken.CollectionPublic},
            quantity: UInt64
        ){
            var i: UInt64 = 0
            while i < quantity {
                self.mintNFT(recipient: recipient)
                i = i + 1;
            }
        }
    }

    // This is an implementation of a custom metadata view for Leofy.
    // This view contains the Item metadata.
    //
    pub struct LeofyNFTMetadataView {
        pub let author: String
        pub let name: String
        pub let description: String
        pub let thumbnail: String
        pub let itemID: UInt64
        pub let serialNumber: UInt32

        init(
            author: String,
            name: String,
            description: String,
            thumbnail: AnyStruct{MetadataViews.File},
            itemID: UInt64,
            serialNumber: UInt32
        ){
            self.author = author
            self.name = name
            self.description = description
            self.thumbnail = thumbnail.uri()
            self.itemID = itemID
            self.serialNumber = serialNumber
        }
    }

    pub resource NFT: NonFungibleToken.INFT, MetadataViews.Resolver  {
        pub let id: UInt64
        pub let itemID: UInt64
        pub let serialNumber: UInt32

        init(
            id: UInt64,
            itemID: UInt64,
            serialNumber: UInt32
        ) {
            self.id = id
            self.itemID = itemID
            self.serialNumber = serialNumber
        }

        pub fun description(): String {
            return "NFT: '"
                .concat(LeofyNFT.getItemMetaDataByField(itemID: self.itemID, field: "name") ?? "''")
                .concat("' from Author: '")
                .concat(LeofyNFT.getItemMetaDataByField(itemID: self.itemID, field: "author") ?? "''")
                .concat("' with serial number ")
                .concat(self.serialNumber.toString())
        }

        pub fun getViews(): [Type] {
            return [
                Type<MetadataViews.Display>(),
                Type<LeofyNFTMetadataView>()
            ]
        }

        pub fun resolveView(_ view: Type): AnyStruct? {
            switch view {
                case Type<MetadataViews.Display>():
                    return MetadataViews.Display(
                        name: LeofyNFT.getItemMetaDataByField(itemID: self.itemID, field: "name") ?? "",
                        description: self.description(),
                        thumbnail: MetadataViews.HTTPFile(LeofyNFT.getItemMetaDataByField(itemID: self.itemID, field: "thumbnail") ?? "")
                    )
                case Type<LeofyNFTMetadataView>():
                    return LeofyNFTMetadataView(
                        author: LeofyNFT.getItemMetaDataByField(itemID: self.itemID, field: "author") ?? "", 
                        name: LeofyNFT.getItemMetaDataByField(itemID: self.itemID, field: "name") ?? "",
                        description: self.description(),
                        thumbnail: MetadataViews.HTTPFile(LeofyNFT.getItemMetaDataByField(itemID: self.itemID, field: "thumbnail") ?? ""),
                        itemID: self.itemID,
                        serialNumber: self.serialNumber
                    )
            }

            return nil
        }
    }

    pub resource interface LeofyCollectionPublic {
        pub fun deposit(token: @NonFungibleToken.NFT)
        pub fun getIDs(): [UInt64]
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT
        pub fun borrowLeofyNFT(id: UInt64): &LeofyNFT.NFT? {
            post {
                (result == nil) || (result?.id == id):
                    "Cannot borrow LeofyNFT reference: the ID of the returned reference is incorrect"
            }
        }
    }

    pub resource Collection: LeofyCollectionPublic, NonFungibleToken.Provider, NonFungibleToken.Receiver, NonFungibleToken.CollectionPublic, MetadataViews.ResolverCollection {
        // dictionary of NFT conforming tokens
        // NFT is a resource type with an `UInt64` ID field
        pub var ownedNFTs: @{UInt64: NonFungibleToken.NFT}

        init () {
            self.ownedNFTs <- {}
        }

        // withdraw removes an NFT from the collection and moves it to the caller
        pub fun withdraw(withdrawID: UInt64): @NonFungibleToken.NFT {
            let token <- self.ownedNFTs.remove(key: withdrawID) ?? panic("missing NFT")

            emit Withdraw(id: token.id, from: self.owner?.address)

            return <-token
        }

        // deposit takes a NFT and adds it to the collections dictionary
        // and adds the ID to the id array
        pub fun deposit(token: @NonFungibleToken.NFT) {
            let token <- token as! @LeofyNFT.NFT

            let id: UInt64 = token.id

            // add the new token to the dictionary which removes the old one
            let oldToken <- self.ownedNFTs[id] <- token

            emit Deposit(id: id, to: self.owner?.address)

            destroy oldToken
        }

        // getIDs returns an array of the IDs that are in the collection
        pub fun getIDs(): [UInt64] {
            return self.ownedNFTs.keys
        }

        // borrowNFT gets a reference to an NFT in the collection
        // so that the caller can read its metadata and call its methods
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT {
            return &self.ownedNFTs[id] as &NonFungibleToken.NFT
        }

        pub fun borrowLeofyNFT(id: UInt64): &LeofyNFT.NFT? {
            if self.ownedNFTs[id] != nil {
                // Create an authorized reference to allow downcasting
                let ref = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT
                return ref as! &LeofyNFT.NFT
            }

            return nil
        }

        pub fun borrowViewResolver(id: UInt64): &AnyResource{MetadataViews.Resolver} {
            let nft = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT
            let LeofyNFT = nft as! &LeofyNFT.NFT
            return LeofyNFT as &AnyResource{MetadataViews.Resolver}
        }

        destroy() {
            destroy self.ownedNFTs
        }
    }

    // Resource that an admin or something similar would own to be
    // able to mint new NFTs
    //
    pub resource Admin {

        pub fun createItem(metadata: {String: String}): UInt64 {

            // Create the new Play
            var newItem <- create Item(
               metadata: metadata
            )
            
            let newID = newItem.itemID

            // Store it in the contract storage
            LeofyNFT.items[newID] <-! newItem
            emit ItemCreated(id: LeofyNFT.totalItemSupply, metadata:metadata)

            // Increment the ID so that it isn't used again
            LeofyNFT.totalItemSupply = LeofyNFT.totalItemSupply + UInt64(1)

            return newID            
        }

        pub fun borrowItem(itemID: UInt64): &Item {
            pre {
                LeofyNFT.items[itemID] != nil: "Cannot borrow Item: The Item doesn't exist"
            }

             return &LeofyNFT.items[itemID] as &Item;

        }
    }

    // -----------------------------------------------------------------------
    // LeofyNFT contract-level function definitions
    // -----------------------------------------------------------------------

    // public function that anyone can call to create a new empty collection
    pub fun createEmptyCollection(): @NonFungibleToken.Collection {
        return <- create Collection()
    }

    // getItemsLength 
    // Returns: Int length of items created
    pub fun getItemsLength(): Int {
        return LeofyNFT.items.length
    }

    // getItemMetaDataByField returns the metadata associated with a 
    //                        specific field of the metadata
    //                        Ex: field: "Artist" will return something
    //                        like "John Doe"
    // 
    // Parameters: itemID: The id of the Item that is being searched
    //             field: The field to search for
    //
    // Returns: The metadata field as a String Optional
    pub fun getItemMetaDataByField(itemID: UInt64, field: String): String? {
        // Don't force a revert if the itemID or field is invalid
        if( LeofyNFT.items[itemID] != nil){
           let item = &LeofyNFT.items[itemID] as &Item
           return item.metadata[field]
        }
        else{
            return nil;
        }
    }

    // -----------------------------------------------------------------------
    // LeofyNFT initialization function
    // -----------------------------------------------------------------------

    init() {
        self.CollectionStoragePath = /storage/LeofyNFTCollection
        self.CollectionPublicPath = /public/LeofyNFTCollection
        self.AdminStoragePath = /storage/LeofyNFTMinter

        // Initialize the total supply
        self.totalSupply = 0
        self.totalItemSupply = 0

        self.items <- {}

        // Create a Collection resource and save it to storage
        let collection <- create Collection()
        self.account.save(<-collection, to: /storage/LeofyNFTCollection)

        // create a public capability for the collection
        self.account.link<&LeofyNFT.Collection{NonFungibleToken.CollectionPublic, LeofyNFT.LeofyCollectionPublic}>(
            self.CollectionPublicPath,
            target: self.CollectionStoragePath
        )

        // Create a Minter resource and save it to storage
        self.account.save(<-create Admin(), to: self.AdminStoragePath)

        emit ContractInitialized()
    }
}
