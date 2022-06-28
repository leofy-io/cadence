import NonFungibleToken from "../../../contracts/standard/NonFungibleToken.cdc"
import LeofyMarketPlace from "../../../contracts/LeofyMarketPlace.cdc"

pub fun main(): UFix64 {
  return LeofyMarketPlace.minimumBidIncrement;
}