package main

import (
	"context"
	"fmt"
	"gostub/config"
	"gostub/flags"
	"gostub/internal/service"
	"gostub/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	f := flags.Get()
	config.SetConfigLocation(f.Config())
	log.SetLevelS(f.Logging())

	log.Debug(fmt.Sprintf("load config from: %s", f.Config()))
	config.Reload()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sigchan := make(chan os.Signal, 1)
	defer close(sigchan)

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigchan
		cancel()
	}()

	runServices(ctx)
}

func runServices(ctx context.Context) {
	cfg := config.Get()

	wg := new(sync.WaitGroup)
	wg.Add(len(cfg.Services))
	for _, s := range cfg.Services {
		go func() {
			defer wg.Done()

			log.Info(fmt.Sprintf("start service: %s", s.Name))
			opts := service.ServiceOpts{
				Name:     s.Name,
				Host:     s.Host,
				Port:     s.Port,
				StubRoot: cfg.StubRoot,
			}
			service := service.NewService(opts)
			err := service.Run(ctx)
			if err != nil {
				log.Error(err.Error())
			}
		}()
	}
	wg.Wait()
}
