package gowalk

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	ctx    context.Context
	opt    option
	cancel func()
}

func New(opts ...Option) *App {
	// 初始化基础配置
	o := option{
		ctx:         context.Background(),
		sigs:        []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		stopTimeout: 5 * time.Second, // 默认接收到终止信号5s后停止服务
	}

	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)

	return &App{
		ctx:    ctx,
		cancel: cancel,
		opt:    o,
	}

}

// Run 运行服务
func (a *App) Run() error {

	if len(a.opt.servers) > 0 {
		// ctx不应该传递到下游服务
		eg, ctx := errgroup.WithContext(a.ctx)
		for _, srv := range a.opt.servers {
			//使用新变量,防止多次循环使用同一个变量
			s := srv
			eg.Go(func() error {
				return s.Start(a.ctx)
			})

			eg.Go(func() error {
				// 等待cancel信号
				<-ctx.Done()
				stopCtx, cancel := context.WithTimeout(ctx, a.opt.stopTimeout)
				defer cancel()
				return s.Stop(stopCtx)
			})
		}

		// 通过信号进行关闭
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, a.opt.sigs...)

		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-quit:
				if a.cancel != nil {
					a.cancel()
				}
				return errors.New(fmt.Sprintf("get os signal: %v", sig))
			}
		})

		if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
			return err
		}

		return nil
	}

	return errors.New("no service to start")

}
