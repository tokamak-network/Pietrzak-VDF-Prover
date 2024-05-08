package node

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type NodeServer struct {
	httpServer *http.Server
	rpcServer  *rpc.Server
}

func NewRPCServer(addr string, apis []rpc.API) *NodeServer {
	srv := rpc.NewServer()
	for _, api := range apis {
		if err := srv.RegisterName(api.Namespace, api.Service); err != nil {
			log.Fatalf("Failed to register RPC API: %v", err)
		}
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade failed:", err)
			return
		}
		defer conn.Close()
	})

	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &NodeServer{
		httpServer: httpServer,
		rpcServer:  srv,
	}
}

func (ns *NodeServer) Start() {
	log.Printf("Starting server at %s\n", ns.httpServer.Addr)
	if err := ns.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func (ns *NodeServer) Stop() {
	if err := ns.httpServer.Shutdown(nil); err != nil {
		log.Println("Failed to shut down HTTP server:", err)
	}
}
