package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storageService = &StorageService{}
	ctx            = context.Background()
)

/*
	Note that in a real world usage, the cache duration shouldn't have
  an expiration time, an LRU policy config should be set where the
  values that are retrieved less often are purged automatically from
  the cache and stored back in RDBMS whenever the cache is full.
*/

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)

	storageService.redisClient = redisClient
	return storageService
}

/*
	We want to be able to save the mapping between the originalUrl
	and the generated shortUrl url
*/

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Error saving url mapping: %v", err))
	}
}

/*
	We should be able to retrieve the initial long URL once the short
	is provided. This is when users will be calling the shortlink in the
	url, so what we need to do here is to retrieve the long url and
	think about redirect.
*/

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storageService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Error retrieving url: %v", err))
	}

	return result
}
