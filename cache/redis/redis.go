package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	// TTL is the default value for time to live it is set to 10 seconds.
	TTL = time.Second * 10
)

type RedisDB struct {
	client *redis.Client
}

func New(dns string) *RedisDB {

	client := redis.NewClient(&redis.Options{
		Addr:     dns, // use default Addr
		Password: "",  // no password set
		DB:       0,   // use default DB
	})

	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis: success connection:", pong)

	return &RedisDB{
		client: client,
	}
}

func (cache *RedisDB) Set(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return cache.client.Set(ctx, key, string(data), TTL).Err()
}

func (cache *RedisDB) Get(ctx context.Context, key string) ([]byte, error) {
	cache.client.Get(ctx, key)

	val, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("key does not exists", key)
			return nil, redis.Nil
		}
		log.Printf("redis: try to get data from the key %v get error %v", key, err)
	}
	return []byte(val), nil
}
