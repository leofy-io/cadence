pub fun main(address: Address) {
    let account = getAccount(address)
    log(account.storageUsed)
    log(account.storageCapacity)
}