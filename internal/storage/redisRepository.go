package storage

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr, password string) *RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return &RedisRepository{client: rdb}
}

func (r *RedisRepository) GetModeCounts(ctx context.Context, areaCode string) (map[string]int, error) {
	modeCounts := make(map[string]int)
	modes := []string{"solo", "multiplayer", "tdm"}

	for _, mode := range modes {
		countStr, err := r.client.Get(ctx, fmt.Sprintf("%s:%s", areaCode, mode)).Result()
		if err == redis.Nil {
			modeCounts[mode] = 0
		} else if err != nil {
			return nil, err
		} else {
			modeCounts[mode], _ = strconv.Atoi(countStr)
		}
	}
	return modeCounts, nil
}

func (r *RedisRepository) UpdateModeCount(ctx context.Context, areaCode, mode string, delta int) error {
	_, err := r.client.IncrBy(ctx, fmt.Sprintf("%s:%s", areaCode, mode), int64(delta)).Result()
	return err
}

func (r *RedisRepository) GetAllModeCounts() (map[string]int, error) {
	ctx := context.Background()
	result := make(map[string]int)

	keys, err := r.client.Keys(ctx, "mode_count:*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		count, err := r.client.Get(ctx, key).Int()
		if err == nil {
			modeName := key[len("mode_count:"):] 
			result[modeName] = count
		}
	}

	return result, nil
}


func (r *RedisRepository) GetModeCount(ctx context.Context, areaCode, mode string) (int, error) {
	count, err := r.client.Get(ctx, fmt.Sprintf("%s:%s", areaCode, mode)).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil 
		}
		return 0, err 
	}


	intCount, err := strconv.Atoi(count)
	if err != nil {
		return 0, err 
	}

	return intCount, nil
}
