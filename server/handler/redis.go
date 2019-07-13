package handler

import "github.com/AquoDev/simple-imageboard-golang/redis"

// Redis client for getting from/setting to cache.
var redisClient = redis.GetRedisClient()
