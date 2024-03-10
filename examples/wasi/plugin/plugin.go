//go:build tinygo.wasm

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/khulnasoft-lab/go-plugin/examples/wasi/cat"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	cat.RegisterFileCat(CatPlugin{})
}

type CatPlugin struct{}

var _ cat.FileCat = (*CatPlugin)(nil)

func (CatPlugin) Cat(_ context.Context, request *cat.FileCatRequest) (*cat.FileCatReply, error) {
	// The message is shown in stdout as os.Stdout is attached.
	fmt.Println("File loading...")
	b, err := os.ReadFile(request.GetFilePath())
	if err != nil {
		return nil, err
	}
	return &cat.FileCatReply{
		Content: string(b),
	}, nil
}
