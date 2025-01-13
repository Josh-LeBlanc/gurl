package store

import (
    "fmt"
    "context"
    "github.com/go-redis/redis"
    "time"
)

// struct wrapper around redis client
type StorageService struct {
    redisClient *redis.Client
}

// top level declarations for the storeService and redix context
var (
    storeService = &StorageService{}
    ctx = context.Background()
)

// Note that in a real world usage, the cache duration shouldn't have  
// an expiration time, an LRU policy config should be set where the 
// values that are retrieved less often are purged automatically from 
// the cache and stored back in RDBMS whenever the cache is full
const CacheDuration = 6 * time.Hour

// initialize the store service and return a store pointer
func InitializeStore() *StorageService {
    redisClient := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    pong, err := redisClient.Ping(ctx).Result()
    if err != nil {
        panic(fmt.Sprintf("Error init Redis: %v", err))
    }

    fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
    storeService.redisClient = redisClient
    return storeService
}

// save mapping between original url and shortened url
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
    err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
    if err != nil {
        panic(fmt.Sprintf("Failed saving key url | Error: %v - short url: %s, long url: %s", err, shortUrl, originalUrl))
    }
}

// retrieve initial url based on short url
func RetrieveInitialUrl(shortUrl string) string {
    result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
    if err != nil {
        panic(fmt.Sprintf("Failed retrieving url | Error: %v - short url: %s", err, shortUrl))
    }
    return result
}
