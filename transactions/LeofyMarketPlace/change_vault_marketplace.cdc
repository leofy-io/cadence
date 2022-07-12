import LeofyMarketPlace from "../../contracts/LeofyMarketPlace.cdc"
import LeofyCoin from "../../contracts/LeofyCoin.cdc"
import FungibleToken from "../../contracts/standard/FungibleToken.cdc"


transaction(to: Address) {
    let adminRef: &LeofyMarketPlace.LeofyMarketPlaceAdmin

    prepare(signer: AuthAccount) {
        self.adminRef = signer.borrow<&LeofyMarketPlace.LeofyMarketPlaceAdmin>(from: LeofyMarketPlace.AdminStoragePath)
			?? panic("Could not borrow reference to the owner's AdminLeofyMarketPlace!")
    }
    execute {
        let recipient = getAccount(to)
        let receiverRef = recipient.getCapability<&AnyResource{FungibleToken.Receiver}>(LeofyCoin.ReceiverPublicPath)
        
        self.adminRef.changeMarketplaceVault(receiverRef)
    }
}