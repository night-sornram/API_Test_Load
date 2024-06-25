package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/guitarpawat/mock-lookup/protos"
	"github.com/spf13/viper"
	"go.openly.dev/pointy"
	_ "go.uber.org/automaxprocs"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//go:embed data.csv
var data string

var cache map[string]string

func init() {
	cache = make(map[string]string, 500)
	viper.SetDefault("HTTP1_PORT1", "8011")
	viper.SetDefault("HTTP1_PORT2", "8012")
	viper.SetDefault("HTTP1_PORT3", "8013")
	viper.SetDefault("HTTP2_PORT1", "8021")
	viper.SetDefault("HTTP2_PORT2", "8022")
	viper.SetDefault("HTTP2_PORT3", "8023")
	viper.SetDefault("GRPC_PORT1", "8031")
	viper.SetDefault("GRPC_PORT2", "8032")
	viper.SetDefault("GRPC_PORT3", "8033")
	viper.SetDefault("GRPC_PORT4", "8034")
	viper.SetDefault("GRPC_PORT5", "8035")
	viper.AutomaticEnv()
	parseCache()
}

type KafkaConfig struct {
	Url   string
	Topic string
}

func main() {
	// HTTP 1.1
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /phone", handleHttpLookup)
		mux.HandleFunc("GET /health", health)
		panic(http.ListenAndServe(":"+viper.GetString("HTTP1_PORT1"), mux))
	}()
	// HTTP 1.1
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /phone", handleHttpLookup)
		mux.HandleFunc("GET /health", health)
		panic(http.ListenAndServe(":"+viper.GetString("HTTP1_PORT2"), mux))
	}()
	// HTTP 1.1
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /phone", handleHttpLookup)
		mux.HandleFunc("GET /health", health)
		panic(http.ListenAndServe(":"+viper.GetString("HTTP1_PORT3"), mux))
	}()

	// HTTP 2
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /phone", handleHttpLookup)
		mux.HandleFunc("GET /health", health)
		h2s := new(http2.Server)
		server := &http.Server{
			Addr:    ":" + viper.GetString("HTTP2_PORT1"),
			Handler: h2c.NewHandler(mux, h2s),
		}
		panic(server.ListenAndServe())
	}()
	// HTTP 2
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /phone", handleHttpLookup)
		mux.HandleFunc("GET /health", health)
		h2s := new(http2.Server)
		server := &http.Server{
			Addr:    ":" + viper.GetString("HTTP2_PORT2"),
			Handler: h2c.NewHandler(mux, h2s),
		}
		panic(server.ListenAndServe())
	}()
	// HTTP 2
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /phone", handleHttpLookup)
		mux.HandleFunc("GET /health", health)
		h2s := new(http2.Server)
		server := &http.Server{
			Addr:    ":" + viper.GetString("HTTP2_PORT3"),
			Handler: h2c.NewHandler(mux, h2s),
		}
		panic(server.ListenAndServe())
	}()

	// gRPC
	go func() {
		lis, err := net.Listen("tcp", ":"+viper.GetString("GRPC_PORT1"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		protos.RegisterLookupServiceServer(grpcServer, new(GrpcLookupServer))
		panic(grpcServer.Serve(lis))
	}()
	// gRPC
	go func() {
		lis, err := net.Listen("tcp", ":"+viper.GetString("GRPC_PORT2"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		protos.RegisterLookupServiceServer(grpcServer, new(GrpcLookupServer))
		panic(grpcServer.Serve(lis))
	}()
	// gRPC
	go func() {
		lis, err := net.Listen("tcp", ":"+viper.GetString("GRPC_PORT3"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		protos.RegisterLookupServiceServer(grpcServer, new(GrpcLookupServer))
		panic(grpcServer.Serve(lis))
	}()
	// gRPC
	go func() {
		lis, err := net.Listen("tcp", ":"+viper.GetString("GRPC_PORT4"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		protos.RegisterLookupServiceServer(grpcServer, new(GrpcLookupServer))
		panic(grpcServer.Serve(lis))
	}()
	// gRPC
	go func() {
		lis, err := net.Listen("tcp", ":"+viper.GetString("GRPC_PORT5"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		protos.RegisterLookupServiceServer(grpcServer, new(GrpcLookupServer))
		panic(grpcServer.Serve(lis))
	}()

	//kafka

	//cfgReq := KafkaConfig{
	//	Url:   "localhost:9092",
	//	Topic: "shopReq",
	//}
	//
	//cfgRes := KafkaConfig{
	//	Url:   "localhost:9092",
	//	Topic: "shopRes",
	//}
	//
	//connReq, err := kafka.DialLeader(context.Background(), "tcp", cfgReq.Url, cfgReq.Topic, 0)
	//if err != nil {
	//	log.Printf("Error connReq dialing leader: %v", err)
	//}
	//defer func() {
	//	if err := connReq.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//connRes, err := kafka.DialLeader(context.Background(), "tcp", cfgRes.Url, cfgRes.Topic, 0)
	//if err != nil {
	//	log.Printf("Error connRes dialing leader: %v", err)
	//}
	//defer func() {
	//	if err := connRes.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//fmt.Println("Connected Kafka Successful!")
	//
	//if _, err := connReq.Seek(0, kafka.SeekEnd); err != nil {
	//	log.Fatal(err)
	//}
	//
	//for {
	//	message, err := connReq.ReadMessage(10e3)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(string(message.Value))
	//	name := lookup(string(message.Value))
	//	if name == "" {
	//		fmt.Println(`{"error": "Phone number not found", "name": null}`)
	//		return
	//	}
	//
	//	data := func() []kafka.Message {
	//		message := make([]kafka.Message, 0)
	//		message = append(message, kafka.Message{Value: []byte(name)})
	//		return message
	//	}()
	//
	//	_, err = connRes.WriteMessages(data...)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-gracefulStop
}

func parseCache() {
	// Split the data into lines
	newLine := fmt.Sprintln()
	lines := strings.Split(data, newLine)
	cache = make(map[string]string, len(lines))

	// Iterate through the lines and parse the data
	for _, line := range lines {
		// Split the line into fields
		fields := strings.Split(line, ",")

		// Extract the key and value from the fields
		key := fields[0]
		value := fields[1]

		// Add the key-value pair to the cache
		cache[key] = value
	}
}

func handleHttpLookup(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("number")
	w.Header().Set("Content-Type", "application/json")

	if phoneNumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Missing phone number", "name": null}`)
		return
	}

	name := lookup(phoneNumber)
	if name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "Phone number not found", "name": null}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"error": null, "name": "%s"}`, name)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "ok"}`)
}

func lookup(phone string) string {
	time.Sleep(3 * time.Second)
	return cache[phone]
}

type GrpcLookupServer struct {
	protos.UnimplementedLookupServiceServer
}

func (*GrpcLookupServer) Lookup(ctx context.Context, req *protos.LookupReq) (*protos.LookupRes, error) {
	if req.PhoneNumber == "" {
		return &protos.LookupRes{
			Error: pointy.String("Missing phone number"),
			Name:  nil,
		}, nil
	}

	name := lookup(req.PhoneNumber)
	if name == "" {
		return &protos.LookupRes{
			Error: pointy.String("Phone number not found"),
			Name:  nil,
		}, nil
	}

	return &protos.LookupRes{
		Error: nil,
		Name:  pointy.String(name),
	}, nil
}
