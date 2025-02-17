package healthCheck

import (
	"context"
	"sync"
	"time"
)

// Server 接口定义服务器健康检查所需的方法
type Server interface {
	GetAddress() string
	IsServerActive() bool
	HealthCheck()
}

type HealthChecker struct {
	servers       map[string]Server
	mu            sync.RWMutex
	checkInterval time.Duration
	ctx           context.Context
}

func NewHealthChecker(checkInterval time.Duration, ctx context.Context) *HealthChecker {
	hc := &HealthChecker{
		servers:       make(map[string]Server),
		checkInterval: checkInterval,
		ctx:           ctx,
	}
	go hc.startHealthCheck()
	return hc
}

func (hc *HealthChecker) AddServer(address string, server Server) {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	if _, exists := hc.servers[address]; !exists {
		hc.servers[address] = server
	}
}

func (hc *HealthChecker) GetActiveServers() []string {
	hc.mu.RLock()
	defer hc.mu.RUnlock()

	var activeServers []string
	for _, server := range hc.servers {
		if server.IsServerActive() {
			activeServers = append(activeServers, server.GetAddress())
		}
	}
	return activeServers
}

func (hc *HealthChecker) startHealthCheck() {
	ticker := time.NewTicker(hc.checkInterval)
	for {
		select {
		case <-hc.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			for _, server := range hc.servers {
				go server.HealthCheck()
			}
		}
	}
}
