package main

import (
	"context"
	"flag"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	log "github.com/sirupsen/logrus"
	"github.com/ilkerispir/envoy-xds-server/internal/processor"
	"github.com/ilkerispir/envoy-api-based/internal/server"
	"github.com/ilkerispir/envoy-api-based/internal/watcher"
)

var (
	l log.FieldLogger

	watchDirectoryFileName string
	port                   uint
	basePort               uint
	mode                   string

	nodeID string
)

func init() {
	l = log.New()
	log.SetLevel(log.DebugLevel)

	flag.UintVar(&port, "port", 9002, "xDS management server port")

	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")

	flag.StringVar(&watchDirectoryFileName, "watchDirectoryFileName", "config/config.yaml", "full path to directory to watch for files")
}

func main() {
	flag.Parse()

	cache := cache.NewSnapshotCache(false, cache.IDHash{}, l)
	
	proc := processor.NewProcessor(
	    cache, nodeID, log.WithField("context", "processor"))
	
	proc.ProcessFile(watcher.NotifyMessage{
	    Operation: watcher.Create,
	    FilePath:  watchDirectoryFileName,
	})

	notifyCh := make(chan watcher.NotifyMessage)
	
	go func() {
	    watcher.Watch(watchDirectoryFileName, notifyCh)
	}()
	
	go func() {
	    ctx := context.Background()
	    srv := serverv3.NewServer(ctx, cache, nil)
	    server.RunServer(ctx, srv, port)
	}()
	
	for {
	    select {
	    case msg := <-notifyCh:
	        proc.ProcessFile(msg)
	    }
	}
}