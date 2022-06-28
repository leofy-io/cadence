import NonFungibleToken from "../../../contracts/standard/NonFungibleToken.cdc"
import LeofyMarketPlace from "../../../contracts/LeofyMarketPlace.cdc"

pub fun main(): UInt64 {
  return LeofyMarketPlace.totalMarketPlaceItems;
}