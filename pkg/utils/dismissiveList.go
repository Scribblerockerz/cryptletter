package utils

import (
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/database"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type DismissiveList interface {
	Add(value string, ttl int64) error
	Set(value string, ttl int64) error
	Has(value string) bool
	Del(value string) error
	All() ([]string, error)
	Drp() error
}

type dismissiveList struct {
	name string
}

//Add will add an entry to the list with a given ttl
func (d dismissiveList) Add(value string, ttl int64) error {
	_, err := database.RedisClient.ZAdd(d.name, redis.Z{
		Score: float64(time.Now().Unix() + ttl),
		Member: value,
	}).Result()
	return err
}

//Set will set ttl on the given value, relative to the current time
func (d dismissiveList) Set(value string, ttl int64) error {
	_, err := database.RedisClient.ZAddXX(d.name, redis.Z{
		Score: float64(time.Now().Unix() + ttl),
		Member: value,
	}).Result()
	return err
}

//Del will remove a given value from the list
func (d dismissiveList) Del(value string) error {
	_, err := database.RedisClient.ZRem(d.name, redis.Z{
		Member: value,
	}).Result()
	return err
}

//Has will check if the entry is present and is still alive
func (d dismissiveList) Has(value string) bool {
	d.cleanup()

	result, err := database.RedisClient.ZScore(d.name, value).Result()
	fmt.Printf("Score: %f", result)
	if err != nil || result == 0 {
		return false
	}

	return true
}

//All will return a list of all stored values
func (d dismissiveList) All() ([]string, error) {
	d.cleanup()
	return database.RedisClient.ZRange(d.name, 0, -1).Result()
}

//Drp will remove all stored values
func (d dismissiveList) Drp() error {
	_, err := database.RedisClient.ZRemRangeByScore(d.name, "-inf", "+inf").Result()
	return err
}

//Cleanup will remove all dead entries from the list
func (d dismissiveList) cleanup() error {
	nowMax := strconv.FormatInt(time.Now().Unix(), 10)
	_, err := database.RedisClient.ZRemRangeByScore(d.name, "-1", nowMax).Result()
	return err
}

//NewDismissiveList will create a new named list
func NewDismissiveList(name string) DismissiveList {
	return &dismissiveList{name: name}
}
