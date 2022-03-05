import LeofyCoin from "../../../contracts/LeofyCoin.cdc"

// This script returns the total amount of LeofyCoin currently in existence.

pub fun main(): UFix64 {

    let supply = LeofyCoin.totalSupply

    log(supply)

    return supply
}
