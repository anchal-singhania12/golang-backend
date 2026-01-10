package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Database struct {
	Driver         string `yaml:"driver"   env:"DATABASE_DRIVER"   env-required:"true"`
	Host           string `yaml:"host"     env:"DATABASE_HOST"     env-required:"true"`
	Port           string `yaml:"port"     env:"DATABASE_PORT"     env-required:"true"`
	Username       string `yaml:"username" env:"DATABASE_USERNAME" env-required:"true"`
	Password       string `yaml:"password" env:"DATABASE_PASSWORD"`
	Name           string `yaml:"database" env:"DATABASE_NAME"     env-required:"true"`
	URL            string `yaml:"url"      env:"DATABASE_URL"      env-required:"true"`
	MigrationsPath string `yaml:"migrations_path"  env:"MIGRATIONS_PATH"   env-default:"pkg/database/migrations"`
}

type TokenKeys struct {
	PublicKey    string `yaml:"public_key" env:"PUBLIC_KEY" env-required:"true"`
	PrivateKey   string `yaml:"private_key" env:"PRIVATE_KEY" env-required:"true"`
	AESKeyBase64 string `yaml:"encryption_key" env:"ENCRYPTION_KEY" env-required:"true"`
}

type Config struct {
	Env             string     `yaml:"env" env:"ENV" env-required:"true"`
	HTTP            HTTPServer `yaml:"http_server"`
	Database        Database   `yaml:"database"`
	AccessTokenKeys TokenKeys  `yaml:"access_token"`
}

func LoadConfig() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "config.yaml", "Path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Configuration file path is required. Please set the CONFIG_PATH environment variable or use the --config flag.")

		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist at path: %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}
	log.Printf("Configuration loaded successfully from %s", configPath)
	// You can use cfg here or return it if needed
	return &cfg
}
