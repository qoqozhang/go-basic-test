package healthCheck

import (
	"context"
	"log"
	"testing"
	"time"
)

// DefaultServer 实现 Server 接口
type DefaultServer struct {
	address      string
	failCount    int
	successCount int
	isActive     bool
}

// NewDefaultServer 创建一个新的 DefaultServer 实例
func NewDefaultServer(address string) *DefaultServer {
	return &DefaultServer{
		address:      address,
		isActive:     false,
		failCount:    0,
		successCount: 0,
	}
}

func (s *DefaultServer) GetAddress() string {
	return s.address
}

func (s *DefaultServer) IsServerActive() bool {
	return s.isActive
}

func (s *DefaultServer) incrementSuccess() {
	s.successCount++
	s.failCount = 0
}

func (s *DefaultServer) incrementFailure() {
	s.failCount++
	s.successCount = 0
}

func (s *DefaultServer) updateStatus() {
	if s.successCount >= 2 && !s.isActive {
		s.isActive = true
	}
	if s.failCount >= 1 {
		s.isActive = false
	}
}

func (s *DefaultServer) HealthCheck() {
	if IcmpHealthCheck(s.address) {
		s.incrementSuccess()
	} else {
		s.incrementFailure()
	}
	s.updateStatus()
}

var servers = []string{"www.baidu.com", "114.114.114.114", "223.5.5.5", "127.0.0.1"}

func TestDefaultServer_HealthCheck_WithContext(t *testing.T) {
	// 创建带有上下文的测试
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	HealthChecker := NewHealthChecker(1*time.Second, ctx)
	for _, server := range servers {
		HealthChecker.AddServer(server, NewDefaultServer(server))
	}
	go func() {
		for {
			log.Printf("servers: %+v\n", HealthChecker.GetActiveServers())
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(2 * time.Minute)
}
