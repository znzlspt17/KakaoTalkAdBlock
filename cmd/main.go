//go:build windows

package main

import (
	"context"
	"kakaotalkadblock/internal"
)

func main() {
	internal.Run(context.Background())
	select {}
}
