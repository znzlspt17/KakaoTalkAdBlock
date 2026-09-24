//go:build windows

package main

import (
	"context"
	"kakaotalkadblock/internal"

	_ "kakaotalkadblock/winres"
)

func main() {
	internal.Run(context.Background())
	select {}
}
