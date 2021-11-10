package dspattn

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

type extConfig struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func parseConfig(fname string, data []byte) (*extConfig, error) {
	ext := filepath.Ext(fname)
	var (
		conf extConfig
		er   error
	)
	switch ext {
	case "json":
		er = json.Unmarshal(data, &conf)
	case "ini":
	case "yml", "yaml":

	default:
		er = fmt.Errorf("unknown ext %s", ext)
	}

	return &conf, er
}

type IStorage interface {
	Store(k, v string)
	Load(k string) string
}

type redisStorage struct {
}

func (f *redisStorage) Store(k, v string) {
}
func (f *redisStorage) Load(k string) string {
	return ""
}

type fileStorage struct {
}

func (f *fileStorage) Store(k, v string) {
}
func (f *fileStorage) Load(k string) string {
	return ""
}

type boltStorage struct {
}

func (f *boltStorage) Store(k, v string) {
}
func (f *boltStorage) Load(k string) string {
	return ""
}

func simpleFatoryGetStorage(sTyp int) IStorage {
	switch sTyp {
	case 0:
		return &fileStorage{}
	case 1:
		return &redisStorage{}
	case 2:
		return &redisStorage{}
	}

	panic(sTyp)
}

type IStorageFactory interface {
	Create() IStorage
}

type redisStorageFatory struct {
}

func (f *redisStorageFatory) Create() IStorage {
	// conn redis
	return &redisStorage{}
}

type fileStorageFatory struct {
}

func (f *fileStorageFatory) Create() IStorage {
	// open file
	return &fileStorage{}
}

func NewStorageFactory(st int) IStorageFactory {
	switch st {
	case 0:
		return &fileStorageFatory{}
	case 1:
		return &redisStorageFatory{}
	}
	panic(st)
}

type IRedisStorageFactory interface {
	RediGoCreate() IStorage
	GoRedisCreate() IStorage
}

type redisStorageFactory struct {
}

func (f *redisStorageFatory) RediGoCreate() IStorage {
	// https://github.com/gomodule/redigo
	return &redisStorage{}
}

func (f *redisStorageFatory) GoRedisCreate() IStorage {
	// https://github.com/go-redis/redis
	return &redisStorage{}
}
