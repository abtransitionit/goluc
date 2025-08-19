/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var CliName = filepath.Base(os.Args[0])

// defines the functions used by a phase
func SetupFunc(ctx context.Context, cmd ...string) (string, error) {
	fmt.Println("  [SETUP] Performing initial setup...")
	time.Sleep(500 * time.Millisecond) // Simulate work
	return "Setup complete!", nil
}

func BuildFunc(ctx context.Context, cmd ...string) (string, error) {
	fmt.Println("  [BUILD] Compiling the application...")
	time.Sleep(1 * time.Second) // Simulate work
	return "Build successful!", nil
}

func TestFunc(ctx context.Context, cmd ...string) (string, error) {
	fmt.Println("  [TEST] Running unit tests...")
	time.Sleep(750 * time.Millisecond) // Simulate work
	return "All tests passed!", nil
}
