import FungibleToken from "../../../contracts/standard/FungibleToken.cdc"
import LeofyCoin from "../../../contracts/LeofyCoin.cdc"

// This script returns an account's LeofyCoin balance.

pub fun main(address: Address): UFix64 {
    let account = getAccount(address)
    
    let vaultRef = account.getCapability(LeofyCoin.BalancePublicPath).borrow<&LeofyCoin.Vault{FungibleToken.Balance}>()
        ?? panic("Could not borrow Balance reference to the Vault")

    return vaultRef.balance
}
