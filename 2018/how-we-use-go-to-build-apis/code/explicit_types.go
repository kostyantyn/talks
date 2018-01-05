package main

import (
	"github.com/go-redis/redis"
	"github.com/honeybadger-io/honeybadger-go"
)

type App struct {
	Redis *redis.Client       // explicit name
	HB    *honeybadger.Client //commonly used abbreviation
}
