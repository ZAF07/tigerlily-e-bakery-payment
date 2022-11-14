package main

// SET UP DB
// SET UP GRPC / PROTOBUF
// SET UP CONTROLLERS / ROUTER
// SET UP SERVICES
// SET UP REPOSITORY

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rest/router"
	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/app"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/db"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/injection"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/service/checkout"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func main() {
	logs := logger.NewLogger()
	logs.InfoLogger.Println("Starting up server ...")

	// Set ENV vars
	// env.SetEnv()

	app.InitApplication()

	// config := config.LoadConfig().GeneralConfig
	config := injection.GetGeneralConfig()
	port := fmt.Sprintf(":%s", config.Port)

	// INIT APPLICATION IN app.go

	// Spin up the main server instance
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logs.ErrorLogger.Println("Something went wrong in the server startup")
		log.Fatalf("Error connecting tcp port %s", port)
	}
	logs.InfoLogger.Println("Successfull server init")

	// Start a new multiplexer passing in the main server
	m := cmux.New(lis)

	// Listen for HTTP requests first
	// If request headers don't specify HTTP, next mux would handle the request
	httpListener := m.Match(cmux.HTTP1Fast())
	grpclistener := m.Match(cmux.Any())

	// Run GO routine to run both servers at diff processes at the same time
	go serveGRPC(grpclistener)
	go serveHTTP(httpListener)

	fmt.Printf("Payment Service Running@%v\n", lis.Addr())

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		log.Fatalf("MUX ERR : %+v", err)
	}

}

// GRPC Server initialisation
func serveGRPC(l net.Listener) {
	grpcServer := grpc.NewServer()

	// Register GRPC stubs (pass the GRPCServer and the initialisation of the service layer)
	rpc.RegisterCheckoutServiceServer(grpcServer, checkout.NewCheckoutService(db.NewDB()))
	// 🚨 TODO: Implement DI for Database instance
	// rpc.RegisterCheckoutServiceServer(grpcServer, checkout.NewCheckoutService(injection.GetPaymentDBInstance()))

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("error running GRPC server %+v", err)
	}
}

// HTTP Server initialisation (using gin)
func serveHTTP(l net.Listener) {
	h := gin.Default()
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		log.Fatalf("error serving HTTP : %+v", err)
	}
}
