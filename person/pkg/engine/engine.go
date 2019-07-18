package engine

import (
	"context"
	`google.golang.org/grpc`
	`net`
	"net/http"
	"os"
	"os/signal"
	"sdkeji/person/pkg/app"
	pb `sdkeji/person/pkg/proto`
	"sdkeji/person/pkg/routing"
	`sdkeji/person/pkg/server`
	"sync"
	"syscall"
	"time"
)

var (
	std *Engine
)

type Engine struct {
	server *http.Server
	close  chan struct{}
	wg     sync.WaitGroup
}

func Get() *Engine {
	return std
}

func NewStdInstance() *Engine {
	app.Init()
	std = new(Engine)
	std.close = make(chan struct{})
	std.server = &http.Server{Addr: app.Conf.HTTPAddr}
	return std
}

func (e *Engine) Run() {
	go e.registerSignal()

	e.wg.Add(2)
	go e.serveHTTP()
	go e.serveRPC()
	e.wg.Wait()
}

func (e *Engine)serveRPC()  {
	lis,err := net.Listen("tcp",":50051")
	if err != nil{
		app.Logger.Info("failed to listen: %v",err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterPersonsServer(s,&server.PersonsServer{})
	s.Serve(lis)
}

func (e *Engine) serveHTTP() {
	defer e.wg.Done()

	e.server.Handler = routing.Register(
		app.Logger,
	)

	app.Logger.Info("listen and serve http service.", "addr", app.Conf.HTTPAddr)

	err := e.server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			app.Logger.Error("an error was returned while listen and serve engine.", "error", err)
			return
		}
	}
	app.Logger.Info("engine shutdown successfully.")
}

func (e *Engine) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	return e.server.Shutdown(ctx)
}

func (e *Engine) registerSignal() {
	app.Logger.Info("register signal handler.")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case sig := <-ch:
		signal.Ignore(syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
		app.Logger.Info("received signal, try to shutdown engine.", "signal", sig.String())
		close(ch)
		close(e.close)
		err := e.shutdown()
		if err != nil {
			app.Logger.Error("fail to shutdown engine.", "error", err)
		}
	}
}
