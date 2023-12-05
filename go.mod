module github.com/morph-l2/bindings

go 1.19

require github.com/scroll-tech/go-ethereum v1.11.4

require (
	github.com/btcsuite/btcd v0.23.3 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)

replace (
	github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.20.1-beta
	github.com/scroll-tech/go-ethereum => github.com/morph-l2/go-ethereum v1.10.14-0.20230705062914-b9d37e7fbd3c
)
