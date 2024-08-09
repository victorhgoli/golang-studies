package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("config")
	// Defina o tipo do arquivo de configuração
	viper.SetConfigType("yml")
	// Adicione o diretório onde o arquivo de configuração está localizado
	viper.AddConfigPath(".")
	// Leia o arquivo de configuração
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
		return err
	}

	// Substitua as configurações com variáveis de ambiente, se existirem
	viper.AutomaticEnv()

	// Defina um prefixo para as variáveis de ambiente (opcional)
	//viper.SetEnvPrefix("APP")

	// Substitua os pontos (.) por underscores (_) nas variáveis de ambiente
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil

}
