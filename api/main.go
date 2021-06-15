package main

import (
	"flag"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	tmp "github.com/lijiaqigreat/personal_server/protobuf"
	"google.golang.org/grpc"
)

var (
	keyLocation  = flag.String("key", "", "path for key file, empty for no tls")
	certLocation = flag.String("cert", "", "path for cert file, empty for no tls")
	addr         = flag.String("addr", ":8000", "http service address, ignored when certLocation not empty")
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message
	// from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period.
	// Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on
	// connection.
	closeGracePeriod = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

/*
 */
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.URL.Path != "/" {
		message := "Not Found.\nDo you mean this?\n" + "TODO"
		http.Error(w, message, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeContent(w, r, "", time.Time{}, strings.NewReader("It works!"))
}

func main() {
	flag.Parse()
	server := grpc.NewServer()
	roomHub := NewRoomHub()
	tmp.RegisterRoomServiceServer(server, roomHub)
	wrappedGrpc := grpcweb.WrapServer(server)
	tlsHttpServer := http.NewServeMux()
	tlsHttpServer.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//if wrappedGrpc.IsGrpcWebRequest(req) {
		//wrappedGrpc.ServeHTTP(resp, req)
		wrappedGrpc.HandleGrpcWebRequest(resp, req)
		// Fall back to other servers.
		//http.DefaultServeMux.ServeHTTP(resp, req)
	})
	http.ListenAndServe(*addr, tlsHttpServer)

	/*
		mux.Handle("/ws", roomHub.GetHandler())
		if *certLocation != "" {
			log.Print(fmt.Sprintf("now serving :443\n"))
			log.Fatal(http.ListenAndServeTLS(":443", *certLocation, *keyLocation, mux))
		} else {
			log.Print(fmt.Sprintf("now serving %s\n", *addr))
			log.Fatal(http.ListenAndServe(*addr, mux))
		}
	*/
}
