package main

import (
	"context"
	"flag"
	"os"

	"xds/sample"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
)

var (
	logger sample.Logger
	port   uint
	nodeID string
)

func init() {
	logger = sample.Logger{}

	flag.BoolVar(&logger.Debug, "debug", false, "Enable xDS server debug logging")
	flag.UintVar(&port, "port", 18000, "xDS management server port")
	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "sample-xds-node-id", "Node ID")
}

func main() {
	flag.Parse()

	// Create a cache
	cache := cache.NewSnapshotCache(false, cache.IDHash{}, logger)

	// Create the snapshot that we'll serve to Envoy
	snapshot := sample.GenerateSnapshot()
	if err := snapshot.Consistent(); err != nil {
		logger.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
		os.Exit(1)
	}
	logger.Debugf("will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := cache.SetSnapshot(context.Background(), nodeID, snapshot); err != nil {
		logger.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	// Run the xDS server
	ctx := context.Background()
	cb := &sample.Callbacks{Debug: logger.Debug}
	srv := server.NewServer(ctx, cache, cb)
	sample.RunServer(ctx, srv, port)
}
