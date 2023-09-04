package testing

import (
	"context"
	"github.com/bloxapp/ssv/operator/validator"
)

// NetworkFactory is a generic factory for network instances
type NetworkFactory func(pctx context.Context, keys NodeKeys) validator.P2PNetwork

// NewLocalTestnet creates a new local network
func NewLocalTestnet(ctx context.Context, n int, factory NetworkFactory) ([]validator.P2PNetwork, []NodeKeys, error) {
	nodes := make([]validator.P2PNetwork, n)
	keys, err := CreateKeys(n)
	if err != nil {
		return nil, nil, err
	}

	for i, k := range keys {
		nodes[i] = factory(ctx, k)
	}

	return nodes, keys, nil
}
