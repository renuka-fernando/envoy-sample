package sample

import (
	"context"
	"log"
	"sync"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

type Callbacks struct {
	Signal         chan struct{}
	Debug          bool
	Fetches        int
	Requests       int
	DeltaRequests  int
	DeltaResponses int
	mu             sync.Mutex
}

func (cb *Callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	log.Printf("server callbacks fetches=%d requests=%d\n", cb.Fetches, cb.Requests)
}
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	logger.Debugf("stream %d open for %s\n", id, typ)
	return nil
}
func (cb *Callbacks) OnStreamClosed(id int64) {
	logger.Debugf("stream %d closed\n", id)
}
func (cb *Callbacks) OnDeltaStreamOpen(_ context.Context, id int64, typ string) error {
	logger.Debugf("delta stream %d open for %s\n", id, typ)
	return nil
}
func (cb *Callbacks) OnDeltaStreamClosed(id int64) {
	logger.Debugf("delta stream %d closed\n", id)
}
func (cb *Callbacks) OnStreamRequest(int64, *discovery.DiscoveryRequest) error {
	logger.Debugf("OnStreamRequest")
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Requests++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}
func (cb *Callbacks) OnStreamResponse(context.Context, int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
	logger.Debugf("OnStreamResponse")
}
func (cb *Callbacks) OnStreamDeltaResponse(id int64, req *discovery.DeltaDiscoveryRequest, res *discovery.DeltaDiscoveryResponse) {
	logger.Debugf("OnStreamDeltaResponse")
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.DeltaResponses++
}
func (cb *Callbacks) OnStreamDeltaRequest(id int64, req *discovery.DeltaDiscoveryRequest) error {
	logger.Debugf("OnStreamDeltaRequest")
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.DeltaRequests++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}

	return nil
}
func (cb *Callbacks) OnFetchRequest(_ context.Context, req *discovery.DiscoveryRequest) error {
	logger.Debugf("OnFetchRequest")
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Fetches++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}
func (cb *Callbacks) OnFetchResponse(*discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
	logger.Debugf("OnFetchResponse")
}
