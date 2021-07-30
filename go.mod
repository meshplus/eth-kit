module github.com/meshplus/eth-kit

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.4
	github.com/holiman/uint256 v1.2.0
	github.com/meshplus/bitxhub-kit v1.2.1-0.20210616114532-4849447f09e1
	github.com/meshplus/bitxhub-model v1.2.1-0.20210728054131-8a811dccae99
	github.com/meshplus/bitxhub  v1.0.0
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
)

replace (
	github.com/meshplus/bitxhub => ../bitxhub
)
