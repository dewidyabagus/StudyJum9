package main

// Implementasi package di dalam package
// import (
// 	"learning/package/config/env"
// )

// func main() {
// 	env.NewEnv()
// }

import (
	"fmt"
	"learning/package/config"
)

func main() {
	var cfg = config.LoadConfig(".")
	fmt.Println("App name:", cfg.AppName)
	fmt.Println("Environment:", cfg.Env)
}
