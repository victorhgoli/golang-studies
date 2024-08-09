package main

import (
	"estudo-test/api/controller"
	"estudo-test/api/graphql"
	grpcx "estudo-test/api/grpc"
	"estudo-test/api/grpc/example"
	"estudo-test/config"
	"estudo-test/db"
	"estudo-test/infra/logger"
	"estudo-test/integration"
	"estudo-test/internal/repository"
	"estudo-test/internal/service"
	"estudo-test/pkg/kafka/consumer"
	"estudo-test/routes"
	"net"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := logger.NewLogrusLogger()

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Configuração de conexão com o banco de dados
	dns := viper.GetString("db.dsn")
	if dns == "" {
		log.Fatal("A variável de ambiente DB_DSN não está definida")
	}

	log.Infof("dns: %v", dns)
	db, err := db.Connect(dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dataTest := integration.NewInfoTestIntegration(log)

	// Inicializar o repositório e o serviço
	userRepo := repository.NewUserRepository(db)
	pedidoRepo := repository.NewPedidoRepository(db)
	userService := service.NewUserService(userRepo, log, dataTest)
	pedidoService := service.NewPedidoService(pedidoRepo)

	cadController := controller.NewCadController(userService, pedidoService, log)
	graphql.UserService = userService
	r := routes.NewRouter(cadController)

	// Iniciar servidor gRPC
	grpcServer := grpc.NewServer()
	example.RegisterUserServiceServer(grpcServer, &grpcx.UserServiceServer{UserService: userService})
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Infof("gRPC server started on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		log.Infof("Servidor iniciado na porta 8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal(err)
		}
	}()

	consumer.StartTesteConsumer()

}
