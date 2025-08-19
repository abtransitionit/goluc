/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package internal

import (
	"context"
	"fmt"
	"log"

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

// Dummy Phase Functions (for demonstration purposes)
func CheckSystemStatus(ctx context.Context, cmd ...string) (string, error) {
	log.Println("Checking system status...")
	time.Sleep(1 * time.Second)
	log.Println("System status OK.")
	return "System status OK", nil
}

func FetchLatestData(ctx context.Context, cmd ...string) (string, error) {
	log.Println("Fetching latest data...")
	time.Sleep(2 * time.Second)
	log.Println("Data fetched successfully.")
	return "Data fetched successfully", nil
}

func ProcessData(ctx context.Context, cmd ...string) (string, error) {
	log.Println("Processing data...")
	time.Sleep(3 * time.Second)
	log.Println("Data processed.")
	return "Data processed", nil
}

func GenerateReport(ctx context.Context, cmd ...string) (string, error) {
	log.Println("Generating report...")
	time.Sleep(1 * time.Second)
	log.Println("Report generated.")
	return "Report generated", nil
}
