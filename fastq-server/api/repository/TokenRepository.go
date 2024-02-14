package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisTokenRepository struct {
	Redis *redis.Client
}

func NewRedisTokenRepository(rdb *redis.Client) *redisTokenRepository {
	return &redisTokenRepository{
		Redis: rdb,
	}
}

func (r *redisTokenRepository) SetRefreshToken(ctx context.Context, userId, tokenId string, expiresIn time.Duration) error {
	key := fmt.Sprintf("%s:%s", userId, tokenId)
	if err := r.Redis.Set(ctx, key, 0, 0).Err(); err != nil {
		log.Printf("error setting token in redis store %v", err)
		return err
	}
	return nil
}

func (r *redisTokenRepository) DeleteRefrashToken(ctx context.Context, userId, tokenId string) error {
	key := fmt.Sprintf("%s:%s", userId, tokenId)

	result := r.Redis.Del(ctx, key)

	if result.Err() != nil {
		log.Printf("error deletig the refresh token/ userId %s/%s : %v", tokenId, userId, result.Err().Error())
		return result.Err()
	}
	if result.Val() < 1 {
		log.Printf("error deletig the refresh token/ userId %s/%s : invalid refresh token", tokenId, userId)
		return errors.New("invalid refresh token")
	}

	return nil
}

// openssl genpkey -algorithm RSA -out ./rsa_private_dev.pem -pkeyopt rsa_keygen_bits:2048
// 	openssl rsa -in ./rsa_private_dev.pem -pubout -out ./rsa_public_dev.pem
