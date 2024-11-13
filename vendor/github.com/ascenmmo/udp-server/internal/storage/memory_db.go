package memoryDB

import (
	"context"
	"sync"
	"time"
)

type IMemoryDB interface {
	GetData(key string) (any, bool)
	SetData(key string, value any)
	Remove(key string)

	AddConnection(id string)
	RemoveConnection(id string)
	GetAllConnection() (ids []string)
	CountConnection() int
}

type MemoryDb struct {
	userData    *userData
	connections *connections
	dataTTL     time.Duration
}

type rowType struct {
	value  any
	usedAt time.Time
}

type userData struct {
	storage sync.Map
	count   int
}

type connections struct {
	storage sync.Map
	count   int
}

func (db *MemoryDb) GetData(key string) (any, bool) {
	value, ok := db.userData.storage.Load(key)
	if !ok {
		return nil, false
	}

	row := value.(*rowType)
	row.usedAt = time.Now()
	return row.value, true
}

func (db *MemoryDb) SetData(key string, value any) {
	db.userData.storage.Store(key, &rowType{value: value, usedAt: time.Now()})
	db.userData.count++
}

func (db *MemoryDb) AddConnection(id string) {
	db.connections.storage.Store(id, &rowType{value: id, usedAt: time.Now()})
	db.connections.count++
}

func (db *MemoryDb) CountConnection() int {
	return db.connections.count
}

func (db *MemoryDb) GetAllConnection() (ids []string) {
	db.connections.storage.Range(func(key, _ any) bool {
		if id, ok := key.(string); ok {
			ids = append(ids, id)
		}
		return true
	})

	return ids
}

func (db *MemoryDb) RemoveConnection(id string) {
	db.connections.storage.Delete(id)
	db.connections.count--
}

func (db *MemoryDb) Remove(key string) {
	db.userData.storage.Delete(key)
}

func (db *MemoryDb) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			db.removeOldData()
		}
	}
}

func (db *MemoryDb) removeOldData() {
	now := time.Now().Add(time.Second * -db.dataTTL)

	db.userData.storage.Range(func(key, value any) bool {
		row := value.(*rowType)
		if !row.usedAt.IsZero() && row.usedAt.Before(now) {
			db.userData.storage.Delete(key)
			db.userData.count--
		}
		return true
	})

	db.connections.storage.Range(func(key, value any) bool {
		row := value.(*rowType)
		if !row.usedAt.IsZero() && row.usedAt.Before(now) {
			db.connections.storage.Delete(key)
			db.connections.count--
		}
		return true
	})

}

func NewMemoryDb(ctx context.Context, dataTTL time.Duration) *MemoryDb {
	db := &MemoryDb{
		userData:    &userData{storage: sync.Map{}},
		connections: &connections{storage: sync.Map{}},
		dataTTL:     dataTTL,
	}
	go db.Run(ctx)
	return db
}
