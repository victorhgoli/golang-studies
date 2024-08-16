/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"estudo-test/api/controller"
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

	"github.com/golobby/container/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var instance = container.New()

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		instance.Singleton(func() repository.UserRepository {
			return repository.NewUserRepository(db)
		})
		instance.Singleton(func() repository.PedidoRepository {
			return repository.NewPedidoRepository(db)
		})
		instance.Singleton(func(userRepo repository.UserRepository) service.UserService {
			return service.NewUserService(userRepo, log, dataTest)
		})
		instance.Singleton(func(pedidoRepo repository.PedidoRepository) service.PedidoService {
			return service.NewPedidoService(pedidoRepo)
		})
		instance.Singleton(func(userService service.UserService, pedidoService service.PedidoService) controller.CadController {
			return controller.NewCadController(userService, pedidoService, log)
		})
		instance.Singleton(func(userService service.UserService, pedidoService service.PedidoService) controller.AsyncCadController {
			return controller.NewAsyncCadController(userService, pedidoService, log)
		})

		/*instance.Call(func(userService service.UserService) {
			graphql.UserService = userService
		})*/

		/* var userService service.UserService
		instance.Resolve(&userService)

		graphql.UserService = userService */
		r := routes.NewRouter(instance)

		// Iniciar servidor gRPC
		grpcServer := grpc.NewServer()
		example.RegisterUserServiceServer(grpcServer, &grpcx.UserServiceServer{UserService: nil})
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

		consumer := consumer.NewTesteConsumer(nil)
		consumer.StartTesteConsumer()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
