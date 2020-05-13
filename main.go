package main

import (
	"context"
	"fmt"
	"github.com/wenwenxiong/go-kubesphere/cmd"
	"os"
)

func main() {
	ctx := context.Background()
	rootCmd := cmd.NewKubesphereCommand(ctx)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}