package acceptancedatastore

import (
	"github.com/kaspanet/kaspad/domain/consensus/model"
)

// acceptanceDataStore represents a store of AcceptanceData
type acceptanceDataStore struct {
}

// New instantiates a new AcceptanceDataStore
func New() model.AcceptanceDataStore {
	return &acceptanceDataStore{}
}

// Insert inserts the given acceptanceData for the given blockHash
func (ads *acceptanceDataStore) Insert(dbTx model.DBTxProxy, blockHash *model.DomainHash, acceptanceData *model.BlockAcceptanceData) error {
	return nil
}

// Get gets the acceptanceData associated with the given blockHash
func (ads *acceptanceDataStore) Get(dbContext model.DBContextProxy, blockHash *model.DomainHash) (*model.BlockAcceptanceData, error) {
	return nil, nil
}

// Delete deletes the acceptanceData associated with the given blockHash
func (ads *acceptanceDataStore) Delete(dbTx model.DBTxProxy, blockHash *model.DomainHash) error {
	return nil
}