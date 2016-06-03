package dsl

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"

	"github.com/mefellows/pact-go/daemon"
	"github.com/mefellows/pact-go/utils"
)

// Use this to wait for a daemon to be running prior
// to running tests
func waitForPortInTest(port int, t *testing.T) {
	timeout := time.After(1 * time.Second)
	for {
		select {
		case <-timeout:
			t.Fatalf("Expected server to start < 1s.")
		case <-time.After(50 * time.Millisecond):
			_, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
			if err == nil {
				return
			}
		}
	}
}

// Use this to wait for a daemon to stop after running a test.
func waitForDaemonToShutdown(port int, t *testing.T) {
	req := ""
	res := ""
	// var req interface{}

	waitForPortInTest(port, t)

	fmt.Println("Sending remote shutdown signal...")
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf(":%d", port))

	err = client.Call("Daemon.StopDaemon", &req, &res)
	// err = client.Call("Daemon.StopDaemon", req, &res)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Println(res)

	t.Logf("Waiting for deamon to shutdown before next test")
	timeout := time.After(1 * time.Second)
	for {
		select {
		case <-timeout:
			t.Fatalf("Expected server to shutdown < 1s.")
		case <-time.After(50 * time.Millisecond):
			conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
			conn.SetReadDeadline(time.Now())
			defer conn.Close()
			if err != nil {
				return
			}
			buffer := make([]byte, 8)
			_, err = conn.Read(buffer)
			if err != nil {
				return
			}
		}
	}
}

func createDaemon(port int) *daemon.Daemon {
	s := &daemon.PactMockService{}
	_, svc := s.NewService()
	d := daemon.NewDaemon(svc)
	go d.StartDaemon(port)
	return d
}

// func TestClient_Fail(t *testing.T) {
// 	client := NewPactClient{ /* don't supply port */ }
//
// }

// Integration style test: Can a client hit each endpoint?
func TestRPCClient_List(t *testing.T) {
	port, _ := utils.GetFreePort()
	createDaemon(port)
	waitForPortInTest(port, t)
	defer waitForDaemonToShutdown(port, t)
	client := &PactClient{Port: port}
	server := client.StartServer()

	waitForPortInTest(server.Port, t)

	s := client.ListServers()

	if len(s.Servers) != 1 {
		t.Fatalf("Expected 1 server to be running, got %d", len(s.Servers))
	}

	// client, err := rpc.DialHTTP("tcp", fmt.Sprintf(":%d", port))
	// var res daemon.PactMockServer
	// err = client.Call("Daemon.StartServer", daemon.PactMockServer{}, &res)
	// if err != nil {
	// 	log.Fatal("rpc error:", err)
	// }
	//
	// waitForPortInTest(res.Port, t)
	//
	// client, err = rpc.DialHTTP("tcp", fmt.Sprintf(":%d", port))
	// var res2 daemon.PactListResponse
	// err = client.Call("Daemon.ListServers", daemon.PactMockServer{}, &res2)
	// if err != nil {
	// 	log.Fatal("rpc error:", err)
	// }
}

// Integration style test: Can a client hit each endpoint?
func TestRPCClient_StartServer(t *testing.T) {
	port, _ := utils.GetFreePort()
	createDaemon(port)
	waitForPortInTest(port, t)

	client, err := rpc.DialHTTP("tcp", fmt.Sprintf(":%d", port))
	var res daemon.PactMockServer
	err = client.Call("Daemon.StartServer", daemon.PactMockServer{}, &res)
	if err != nil {
		log.Fatal("rpc error:", err)
	}

	<-time.After(10 * time.Second)
	waitForDaemonToShutdown(port, t)
}