package utils

import (
	memoryDB "github.com/ascenmmo/tcp-server/internal/storage"
	"time"
)

type RateLimit interface {
	IsLimited(id string) bool
}

type RateLimitRow struct {
	Count       int
	LastRequest time.Time
}

type rateLimit struct {
	rateLimitSize int
	storage       memoryDB.IMemoryDB
}

func (r *rateLimit) IsLimited(id string) bool {
	count, lastRequest := r.getCountAndTime(id)
	defer r.setCountAndTime(id, count, lastRequest)

	isBefore := lastRequest.Before(time.Now())
	if !isBefore {
		return true
	}

	if count >= r.rateLimitSize {
		if isBefore {
			count = 0
		}
		lastRequest = time.Now().Add(time.Second * 1)
		return true
	}

	count = count + 1

	return false
}

func (r *rateLimit) getCountAndTime(id string) (count int, timestamp time.Time) {
	data, ok := r.storage.GetData(id)
	if !ok {
		return
	}
	rateLimitRow, ok := data.(RateLimitRow)
	if !ok {
		return
	}
	return rateLimitRow.Count, rateLimitRow.LastRequest
}

func (r *rateLimit) setCountAndTime(id string, count int, lastRequest time.Time) {
	r.storage.SetData(id, RateLimitRow{count, lastRequest})
}

func NewRateLimit(limit int, storage memoryDB.IMemoryDB) RateLimit {
	return &rateLimit{rateLimitSize: limit, storage: storage}
}
