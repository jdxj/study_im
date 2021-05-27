package main

import (
	"math/rand"
	"net"
	"testing"
)

type mockConn struct {
}

func (mc *mockConn) Context() interface{}    { return nil }
func (mc *mockConn) SetContext(interface{})  {}
func (mc *mockConn) LocalAddr() net.Addr     { return nil }
func (mc *mockConn) RemoteAddr() net.Addr    { return nil }
func (mc *mockConn) Read() []byte            { return nil }
func (mc *mockConn) ResetBuffer()            {}
func (mc *mockConn) ReadN(int) (int, []byte) { return 0, nil }
func (mc *mockConn) ShiftN(int) int          { return 0 }
func (mc *mockConn) BufferLength() int       { return 0 }
func (mc *mockConn) SendTo([]byte) error     { return nil }
func (mc *mockConn) AsyncWrite([]byte) error { return nil }
func (mc *mockConn) Wake() error             { return nil }
func (mc *mockConn) Close() error            { return nil }

var (
	gCM = &ClientManager{
		clients: make(map[uint32]*Client),
	}

	gGM = &GroupManager{
		groups: make(map[uint32]*Group),
	}
)

func BenchmarkCM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.Int() % 3 // 只测试 0, 1, 2
		var i uint32
		for pb.Next() {
			if r == 0 {
				gCM.AddClient(i, &Client{
					userID: i,
					conn:   &mockConn{},
				})
				i++
			} else if r == 1 {
				gCM.GetClient(i)
				i++
			} else if r == 2 {
				gCM.DelClient(i)
				i++
			} else if r == 3 {
				gCM.Range(func(u uint32, client *Client) {
					client.conn.SetContext(u)
				})
			}
		}
	})
}

func BenchmarkGM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.Int() % 2 // 只测试 0, 1
		var i uint32
		for pb.Next() {
			if r == 0 {
				gGM.AddMember(uint32(r), i, &Client{
					userID: i,
					conn:   &mockConn{},
				})
				i++
			} else if r == 1 {
				gGM.DelMember(0, i)
				i++
			} else if r == 2 {
				gGM.Range(0, func(u uint32, client *Client) {
					client.conn.SetContext(u)
				})
			}
		}
	})
}
