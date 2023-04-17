package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

var rdb *redis.Client

func init() {
	tmp := redis.NewClient(&redis.Options{
		Addr:     "121.37.119.47:6379",
		DB:       2,
		Password: "44913730",

		PoolSize:     5,
		MinIdleConns: 2,

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,
	})
	rdb = tmp
}

func Record(UserID string, num int) error {
	result, err := rdb.Do(context.Background(), "setbit", "record:"+UserID, num*4, 0).Result()
	if err != nil {
		return err
	}
	log.Print(result)
	return nil
}

func getBitPos(id string) (int, error) {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return atoi * 4, nil
}

func UpDateRec(UserID string, policyID string) error {

	bitpos, err := getBitPos(policyID)
	if err != nil {
		return err
	}
	result, err := rdb.Do(context.Background(), "bitfield", "record:"+UserID, "overflow", "sat", "incrby", "u4", strconv.Itoa(bitpos), 1).Result()
	if err != nil {
		return err
	}
	log.Print(result)
	return nil
}

func UpdateRecm(UserID string, policyID string) error {

	bitpos, err := getBitPos(policyID)
	if err != nil {
		return err
	}
	result, err := rdb.Do(context.Background(), "bitfield", "recommend:"+UserID, "overflow", "sat", "incrby", "u4", strconv.Itoa(bitpos), 1).Result()
	if err != nil {
		return err
	}
	log.Print(result)
	return nil
}

func GetRec(UserID string, policyID string) (interface{}, error) {
	bitpos, err := getBitPos(policyID)
	if err != nil {
		return nil, err
	}
	result, err := rdb.Do(context.Background(), "bitfield", "record:"+UserID, "get", "u4", strconv.Itoa(bitpos)).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllRecommend(UserID string) (map[string]string, error) {
	result, err := rdb.HGetAll(context.Background(), "recommend:"+UserID).Result()
	if err != nil {
		return nil, err
	}
	//log.Print(result)
	return result, nil
}
