import NonFungibleToken from "../../../contracts/standard/NonFungibleToken.cdc"
import LeofyNFT from "../../../contracts/LeofyNFT.cdc"

// This script borrows an NFT from a collection
pub fun main(x: UFix64, y: UFix64): UInt64 {
    let test:UFix64 = x*y
    let integer:UInt64 = UInt64(test)

    return integer


}