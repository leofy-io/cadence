module github.com/leofy-io/cadence/lib/go/test

go 1.14

require (
	github.com/onflow/cadence v0.21.3
	github.com/onflow/flow-emulator v0.28.2
	github.com/onflow/flow-go-sdk v0.24.0
	github.com/onflow/flow-nft/lib/go/contracts v0.0.0-20220119224830-7a2e698160ea
	github.com/onflow/flow-nft/lib/go/templates v0.0.0-20220119224830-7a2e698160ea
	github.com/stretchr/testify v1.7.0
)

replace github.com/onflow/flow-nft/lib/go/contracts => ../contracts

replace github.com/onflow/flow-nft/lib/go/templates => ../templates
