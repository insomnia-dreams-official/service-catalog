package server

import (
	"database/sql"
	"fmt"
	"github.com/insomnia-dreams-official/service-catalog/internal/delivery_grpc"
	"github.com/insomnia-dreams-official/service-catalog/internal/repo_postgres"
	"github.com/insomnia-dreams-official/service-catalog/internal/ucase"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"path"
)

type Server struct {
	grpcServer *grpc.Server
}

// Load environment variables from a config file corresponding to stack (dev/test/prod).
func init() {
	// Define config file's path
	if dir, err := os.Getwd(); err == nil {
		viper.AddConfigPath(path.Join(dir, "config"))
	} else {
		log.Fatalf("can't get process's working directory for loading viper config file")
	}

	// Define config file's name
	switch os.Getenv("ENVIRONMENT") {
	case "dev":
		viper.SetConfigName("dev.config")
	case "test":
		viper.SetConfigName("test.config")
	case "prod":
		viper.SetConfigName("prod.config")
	default:
		log.Fatalf("define environment variable \"ENVIRONMENT\" with one of values: \"dev\", \"test\", \"prod\"; then restart service")
	}

	// Load config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("can't load viper config file, reason: %v", err)
	}
}

// New function create all repositories, usecases and inject them in our implementation of grpc server;
// to call them (usecases) in grpc methods
func New(db *sql.DB) Server {
	grpcServer := grpc.NewServer()

	// Create repositories
	categoryRepo := repo_postgres.NewCategoryRepo(db)

	// Create usecases
	navigationUcase := ucase.NewNavigationUcase(categoryRepo)

	// Inject ucases in out grpc server
	delivery_grpc.Register(grpcServer, navigationUcase)
	return Server{grpcServer}
}

func (s *Server) Run() {
	// Create tcp listener for grpc dial
	address := fmt.Sprintf("%s:%s", viper.GetString(`catalog.host`), viper.GetString(`catalog.port`))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Run grpc-server
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
