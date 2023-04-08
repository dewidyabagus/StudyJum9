package main

import "fmt"

type Config struct {
	AppName    string
	Env        string
	ListenPort int
}

// Type By Reference
func Unmarshal(cfg *Config) {
	fmt.Printf("Unmarshal => Memory Address: %p \n", cfg)

	cfg.AppName = "svc-payment"
	cfg.Env = "production"
	cfg.ListenPort = 8080
}

// TYpe By Value
func Unmarshal2(cfg Config) {
	fmt.Printf("Unmarshal2 => Memory Address: %p \n", &cfg)

	cfg.AppName = "svc-payment"
	cfg.Env = "production"
	cfg.ListenPort = 8080
}

func main() {
	var cfg Config

	fmt.Printf("Memory Address: %p \n", &cfg)

	// Unmarshal(&cfg)
	Unmarshal2(cfg)

	fmt.Println("App name:", cfg.AppName)
	fmt.Println("Env:", cfg.Env)
	fmt.Println("Listen port:", cfg.ListenPort)
}
