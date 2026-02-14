package cache

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/YagorX/go-service-ci/internal/model"
	"github.com/redis/go-redis/v9"

	"github.com/alicebob/miniredis/v2"
)

func TestCache_Set(t *testing.T) {
	s := miniredis.RunT(t)
	defer s.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	defer redisClient.Close()

	satelliteCache := New[*model.Satellite](redisClient)

	err := satelliteCache.Set(t.Context(), "testKey",
		&model.Satellite{Name: "moon"}, 3*time.Second)

	if err != nil {
		t.Fatal(err)
	}

	satelliteFromRedis, err := s.Get("testKey")
	if err != nil {
		t.Fatal(err)
	}

	if satelliteFromRedis == "" {
		t.Fatal("satelliteFromRedis is empty")
	}
}

func TestCache_SetTTLWorks(t *testing.T) {
	s := miniredis.RunT(t)
	defer s.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	defer redisClient.Close()

	satelliteCache := New[*model.Satellite](redisClient)

	err := satelliteCache.Set(t.Context(), "testKey",
		&model.Satellite{Name: "moon"}, 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	s.FastForward(2 * time.Second)

	satelliteCached, err := s.Get("testKey")
	if err == nil {
		t.Fatal(errors.New("expected error, got nil"))
	}

	fmt.Println(satelliteCached)

}
