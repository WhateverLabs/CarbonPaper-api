package repository

import (
	"sync"
	"time"
)

type RatelimitRepository struct {
	ipMap *sync.Map
}

func NewRatelimitRepository() *RatelimitRepository {
	ipMap := sync.Map{}
	return &RatelimitRepository{
		ipMap: &ipMap,
	}
}

func (repo *RatelimitRepository) Increment(ip string) {
	val, _ := repo.ipMap.LoadOrStore(ip, 1)
	repo.ipMap.Store(ip, val.(int)+1)

	time.AfterFunc(time.Minute, func() {
		repo.decrement(ip)
	})
}

func (repo *RatelimitRepository) decrement(ip string) {
	val, ok := repo.ipMap.Load(ip)
	if !ok {
		return
	}
	repo.ipMap.Store(ip, val.(int)-1)

	if val.(int) <= 0 {
		repo.ipMap.Delete(ip)
	}
}

func (repo *RatelimitRepository) Get(ip string) int {
	val, ok := repo.ipMap.Load(ip)
	if !ok {
		return 0
	}
	return val.(int)
}

func (repo *RatelimitRepository) Delete(ip string) {
	repo.ipMap.Delete(ip)
}
