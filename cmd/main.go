//go:build windows

package main

import (
	"context"
	"kakaotalkadblock/internal"
	"kakaotalkadblock/internal/win"
	"log"
)

func main() {
	if err := win.SetStartupEnabled(true); err != nil {
		log.Printf("자동실행 등록 실패: %v", err)
	}
	internal.Run(context.Background())
	select {}
}
