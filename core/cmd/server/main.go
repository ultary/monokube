package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/ultary/monokube/core/pkg/api/grpc"
	"github.com/ultary/monokube/core/pkg/k8s"
	"github.com/ultary/monokube/core/pkg/kube/watch"
)

func main() {

	////////////////////////////////////////////////////////////////
	//
	//  PostgreSQL
	//

	// https://www.npgsql.org/doc/connection-string-parameters.html
	const pgdsn = "postgres://postgres:postgrespw@localhost:5432/postgres?application_name=ultary&sslmode=disable&connect_timeout=5"
	pgconfig, err := pgxpool.ParseConfig(pgdsn)
	if err != nil {
		log.Fatal(err)
	}
	pgpool, err := pgxpool.NewWithConfig(context.Background(), pgconfig)
	if err != nil {
		log.Fatal(err)
	}
	defer pgpool.Close()

	////////////////////////////////////////////////////////////////
	//
	//  Server Instances
	//

	client := k8s.NewClient()
	grpcServer := grpc.NewServer()
	watchTower := watch.NewTower(client)

	var wg sync.WaitGroup

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer wg.Done()
		_ = <-sigs

		log.Info("Gracefully shutting down ...")
		grpcServer.Stop()
		watchTower.Shutdown()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		watchTower.Watch()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		network, address := "tcp4", "127.0.0.1:50051"
		//network, address := "unix", "/tmp/kluster.sock"
		if err := grpcServer.Serve(network, address); err != nil {
			log.Fatalf("Failed to serve grpc server: %v", err)
		}
	}()

	wg.Wait()
}
