package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Rom4eg/gostub/config"
	"github.com/Rom4eg/gostub/flags"
	"github.com/Rom4eg/gostub/internal/manager"
	"github.com/Rom4eg/gostub/internal/service"
	"github.com/Rom4eg/gostub/log"

	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	SetLogging()
	LoadConfig()

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

func LoadConfig() {
	f := flags.Get()
	if f.HasConfig() {
		config.SetConfigLocation(f.Config())
	} else {
		cp, err := config.FindConfig()
		if err != nil {
			panic(err)
		}

		config.SetConfigLocation(cp)
	}

	log.Debug(fmt.Sprintf("load config from: %s", config.GetConfigLocation()))
	config.Reload()
}

func SetLogging() {
	f := flags.Get()
	log.SetLevelS(f.Logging())
}

func runServices(ctx context.Context) {
	cfg := config.Get()

	wg := new(sync.WaitGroup)
	wg.Add(len(cfg.Services))
	m := manager.New(log.NewLogger(""))
	f := service.NewFactory()
	for _, s := range cfg.Services {
		go func() {
			defer wg.Done()

			go func() {
				<-ctx.Done()
				err := m.StopService(s.Name)
				if err != nil {
					log.Error(err.Error())
				}
			}()

			for {
				if ctx.Err() != nil {
					log.Info(ctx.Err().Error())
					return
				}

				opts := service.FactoryOpt{
					Type:       service.ServiceType(s.Type),
					Logger:     log.NewLogger(fmt.Sprintf("[%s]", s.Name)),
					ServiceOpt: s.Options,
				}
				srv, err := f.MakeService(s.Name, opts)
				if err != nil {
					log.Error(err.Error())
					return
				}

				err = m.StartService(s.Name, srv)
				if err == nil || errors.Is(err, http.ErrServerClosed) {
					continue
				}

				msg := fmt.Errorf("service \"%s\" crashed with error - %w", s.Name, err)
				log.Error(msg.Error())

				err = m.StopService(s.Name)
				if err != nil {
					log.Error(msg.Error())
				}

				log.Info(fmt.Sprintf("restarting \"%s\" in 5 seconds", s.Name))
				time.Sleep(5 * time.Second)
			}
		}()
	}
	wg.Wait()
}
