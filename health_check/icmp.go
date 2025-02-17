package healthCheck

import (
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func IcmpHealthCheck(address string) bool {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return false
	}
	defer conn.Close()

	// 开始构建icmp 请求
	// 1. 构建icmp 请求
	echo := &icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("hello"),
		},
	}
	echoBytes, _ := echo.Marshal(nil)
	// 2. 发送icmp 请求
	ip, _ := net.ResolveIPAddr("ip", address)
	conn.WriteTo(echoBytes, ip)
	// 3. 接收icmp 响应
	reply := make([]byte, 1500)
	// 4. 循环3次，如果3次都失败，则认为服务器不可用
	for i := 0; i < 3; i++ {
		err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			return false
		}
		n, peer, err := conn.ReadFrom(reply)
		if err != nil {
			return false
		}
		msg, err := icmp.ParseMessage(1, reply[:n])
		if err != nil {
			return false
		}
		switch msg.Type {
		case ipv4.ICMPTypeEchoReply:
			echoReply, ok := msg.Body.(*icmp.Echo)
			if !ok {
				return false
			}
			if peer.String() == ip.String() && echoReply.ID == os.Getpid()&0xffff && echoReply.Seq == 1 {
				return true
			}
		default:
			return false
		}
	}
	return false
}
