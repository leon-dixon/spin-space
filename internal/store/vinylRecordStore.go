package store

import (
	"context"
	"errors"
	"fmt"
	"spin-space/internal/model"
	"sync"
)

type VinylRecordStore struct {
	records map[uint]*model.VinylRecord
	nextID  uint
	mu      sync.Mutex
}

func NewVinylRecordStore() *VinylRecordStore {
	return &VinylRecordStore{
		records: make(map[uint]*model.VinylRecord),
		nextID:  1,
	}
}

func (store *VinylRecordStore) Create(context context.Context, record *model.VinylRecord) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	record.ID = store.nextID
	fmt.Println("Before ID: %d", store.nextID)
	store.records[record.ID] = record
	store.nextID++
	fmt.Println("After ID: %d", store.nextID)
	return nil
}

func (store *VinylRecordStore) GetByID(id uint) (*model.VinylRecord, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	record, exists := store.records[id]
	if !exists {
		return nil, errors.New("record not found")
	}

	return record, nil
}

func (store *VinylRecordStore) GetVinyls() ([]*model.VinylRecord, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	if len(store.records) == 0 {
		return nil, errors.New("no records found")
	}

	records := make([]*model.VinylRecord, 0, len(store.records))
	for _, record := range store.records {
		records = append(records, record)
	}

	return records, nil
}
