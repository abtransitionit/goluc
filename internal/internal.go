/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package internal

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"

	"os"
	"path/filepath"
	"time"
)

var CliName = filepath.Base(os.Args[0])
var logger = logx.GetLogger()

// defines the functions used by a phase
// func ShowPhase(ctx context.Context, cmd ...string) (string, error) {
// 	fmt.Println("  [SETUP] Performing initial setup...")
// 	time.Sleep(500 * time.Millisecond) // Simulate work
// 	return "Setup complete!", nil
// }

func SetupFunc(ctx context.Context, l logx.Logger, cmd ...string) (string, error) {
	fmt.Println("  [SETUP] Performing initial setup...")
	time.Sleep(500 * time.Millisecond) // Simulate work
	return "Setup complete!", nil
}

func BuildFunc(ctx context.Context, l logx.Logger, cmd ...string) (string, error) {
	fmt.Println("  [BUILD] Compiling the application...")
	time.Sleep(1 * time.Second) // Simulate work
	return "Build successful!", nil
}

func TestFunc(ctx context.Context, l logx.Logger, cmd ...string) (string, error) {
	fmt.Println("  [TEST] Running unit tests...")
	time.Sleep(750 * time.Millisecond) // Simulate work
	return "All tests passed!", nil
}

// Dummy Phase Functions (for demonstration purposes)
func CheckSystemStatus(ctx context.Context, l logx.Logger, targets []phase.Target, cmd ...string) (string, error) {
	logger.Info("Checking system status...")
	time.Sleep(1 * time.Second)
	logger.Info("System status OK.")
	return "System status OK", nil
}

func FetchLatestData(ctx context.Context, l logx.Logger, targets []phase.Target, cmd ...string) (string, error) {
	logger.Info("Fetching latest data...")
	time.Sleep(2 * time.Second)
	logger.Info("Data fetched successfully.")
	return "Data fetched successfully", nil
}

func ProcessData(ctx context.Context, l logx.Logger, targets []phase.Target, cmd ...string) (string, error) {
	logger.Info("Processing data...")
	time.Sleep(3 * time.Second)
	logger.Info("Data processed.")
	return "Data processed", nil
}

func GenerateReport(ctx context.Context, l logx.Logger, targets []phase.Target, cmd ...string) (string, error) {
	logger.Info("Generating report...")
	time.Sleep(1 * time.Second)
	logger.Info("Report generated.")
	return "Report generated", nil
}

func Dummy(ctx context.Context, l logx.Logger, targets []phase.Target, cmd ...string) (string, error) {
	logger.Info("Generating report...")
	time.Sleep(1 * time.Second)
	logger.Info("Report generated.")
	return "Report generated", nil
}
