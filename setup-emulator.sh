flow transactions send cadence/transactions/LeofyCoin/setup_account.cdc --signer emulator-test
flow transactions send cadence/transactions/LeofyNFT/setup_account.cdc --signer emulator-test
flow transactions send cadence/transactions/LeofyMarketPlace/setup_account.cdc --signer emulator-test

flow transactions send cadence/transactions/LeofyMarketPlace/change_cut_percentage.cdc 15.00
flow transactions send cadence/transactions/LeofyMarketPlace/change_bid_increment.cdc 10.50
flow transactions send cadence/transactions/LeofyMarketPlace/change_extends_time.cdc 60.00
flow transactions send cadence/transactions/LeofyMarketPlace/change_extends_time_lower_than.cdc 60.00


<< 'MULTILINE-COMMENT'
flow transactions send ./cadence/transactions/LeofyNFT/create_item.cdc \
--args-json '[{
	"type": "Dictionary",
	"value": [{
			"key": {
				"type": "String",
				"value": "author"
			},
			"value": {
				"type": "String",
				"value": "Leofy"
			}
		},
		{
			"key": {
				"type": "String",
				"value": "name"
			},
			"value": {
				"type": "String",
				"value": "Our First NFT"
			}
		},
		{
			"key": {
				"type": "String",
				"value": "thumbnail"
			},
			"value": {
				"type": "String",
				"value": "https://leofy.io/leofy-logo-y-2.svg"
			}
		}
	]
}, {
	"type": "UFix64",
	"value": "20.00"
}]'

flow transactions send cadence/transactions/LeofyNFT/batch_mint_nft.cdc 0 100 --gas-limit 3000

flow transactions send cadence/transactions/LeofyCoin/mint_tokens.cdc 01cf0e2f2f715450 5000.00
flow scripts execute  cadence/transactions/LeofyCoin/scripts/get_balance.cdc 01cf0e2f2f715450
flow scripts execute  cadence/transactions/LeofyCoin/scripts/get_balance.cdc f8d6e0586b0a20c7

flow transactions send cadence/transactions/LeofyNFT/purchase_nft.cdc 0 --signer emulator-test

MULTILINE-COMMENT

flow transactions send cadence/transactions/LeofyMarketPlace/sell_item.cdc 3 0.00 30.00 0.00 0.00 --signer emulator-test
flow scripts execute cadence/transactions/LeofyMarketPlace/scripts/get_auction_account.cdc 01cf0e2f2f715450 5
flow scripts execute cadence/transactions/LeofyMarketPlace/scripts/get_all_auctions_account.cdc 01cf0e2f2f715450
flow scripts execute cadence/transactions/LeofyMarketPlace/scripts/getTime.cdc
flow scripts execute cadence/transactions/LeofyMarketPlace/scripts/get_bid_increment.cdc
flow scripts execute cadence/transactions/LeofyMarketPlace/scripts/get_total_auctions.cdc
flow scripts execute cadence/transactions/LeofyMarketPlace/scripts/get_storage_used.cdc 01cf0e2f2f715450
flow transactions send cadence/transactions/LeofyMarketPlace/cancel_auction.cdc 10 --signer emulator-test

flow scripts execute cadence/transactions/LeofyNFT/scripts/get_nft_account.cdc 01cf0e2f2f715450 3
flow scripts execute cadence/transactions/LeofyNFT/scripts/get_nft_account.cdc f8d6e0586b0a20c7 1

flow scripts execute cadence/transactions/LeofyMarketPlace/scripts/get_cut_percentage.cdc 

flow transactions send cadence/transactions/LeofyMarketPlace/bid.cdc 2 01cf0e2f2f715450 10.50
flow transactions send cadence/transactions/LeofyMarketPlace/purchase.cdc 5 01cf0e2f2f715450 0.00

flow transactions send cadence/transactions/LeofyMarketPlace/settle_bid.cdc 5 01cf0e2f2f715450