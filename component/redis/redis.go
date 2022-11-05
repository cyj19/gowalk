package redis

import (
	"context"
	"github.com/cyj19/gowalk/core"
	v8 "github.com/go-redis/redis/v8"
)

var (
	componentName = "redis"
	clients       = make(map[string]*v8.Client)
)

type Instance struct {
	Settings map[string]setting `json:"settings"`
}

func (i *Instance) Run() error {
	ctx := context.Background()
	for k, s := range i.Settings {
		client := v8.NewClient(&v8.Options{
			Addr:     s.Addr,
			Password: s.Password,
			DB:       s.DB,
			PoolSize: s.PoolSize,
		})
		if err := client.Ping(ctx).Err(); err != nil {
			return err
		}
		clients[k] = client
	}
	return nil
}

func (i *Instance) Name() string {
	return componentName
}

type setting struct {
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

var _ core.Component = (*Instance)(nil)

func Main() *v8.Client {
	return clients["main"]
}

func Get(name string) *v8.Client {
	return clients[name]
}
