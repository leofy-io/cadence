import LeofyNFT from "../../../contracts/LeofyNFT.cdc"

// This scripts returns the number of Items currently in existence.

pub fun main(): UInt64 {    
    return LeofyNFT.totalItemSupply
}
