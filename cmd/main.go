package main

import (
	"context"
	"errors"
	"fmt"
	"gostub/config"
	"gostub/flags"
	"gostub/internal/manager"
	"gostub/internal/service"
	"gostub/log"
	"net/http"

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
	m := manager.New(log.NewLogger(""))
	f := service.NewFactory()
	for _, s := range cfg.Services {
		go func() {
			defer wg.Done()

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

			done := make(chan struct{})
			go func() {
				defer close(done)
				<-ctx.Done()
				err := m.StopService(s.Name)
				if err != nil {
					log.Error(err.Error())
				}

				done <- struct{}{}
			}()

			err = m.StartService(s.Name, srv)
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Error(err.Error())
				return
			}

			<-done
		}()
	}
	wg.Wait()
}
