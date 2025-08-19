/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"
	"log"
	"os"

	corefilex "github.com/abtransitionit/gocore/filex"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/filex"
	"github.com/abtransitionit/golinux/user"
)

func testCanBeSudoAndIsNotRoot() {
	canBeSudo, err := user.CanBeSudoAndIsNotRoot()
	if err != nil {
		// handle unexpected failure by logging and exiting.
		log.Printf("Failed to check if user can be sudo: %v", err)
		os.Exit(1)
	}

	if canBeSudo {
		// Expected outcome for a non-root, sudo-enabled user.
		log.Println("✅ Current user can be sudo but is not root.")
	} else {
		// Expected outcome for a non-sudo user or a root user.
		log.Println("❌ Current user can NOT be sudo and is not root.")
	}
}

// func testCanBeSudoAndIsNotRoot() {
// 	canBeSudo, err := user.CanBeSudoAndIsNotRoot()
// 	if err != nil {
// 		// handle generic error explicitly: unexpected failure
// 		logx.Error(fmt.Sprintf("Failed to check if user can be sudo: %v", err))
// 		os.Exit(1)
// 	}
// 	if canBeSudo {
// 		// expected outcome
// 		logx.Info("✅ Current user can be sudo.")
// 	} else {
// 		// expected outcome
// 		logx.Info("❌ Current user can NOT be sudo.")
// 	}
// }

// testTouchAsSudo verifies the behavior of the TouchAsSudo function.
// It attempts to create a file with sudo privileges and reports the result.
func testTouchAsSudo() {
	var testFilePath = "/tmp/max"
	// Defer a cleanup function to remove the test file regardless of the outcome.
	// This ensures the test is idempotent.
	defer func() {
		success, err := filex.DeleteAsSudo(testFilePath)
		if err != nil {
			// handle generic error explicitly: unexpected failure
			logx.Error(fmt.Sprintf("❌ Failed to clean up test file: %v", err))
		} else if success {
			logx.Info("✅ Successfully cleaned up test file.")
		} else {
			// function explicitly returned 'false' without an error: expected outcome
			logx.Info("❌ Failed to clean up test file (DeleteAsSudo returned false).")
		}
	}()

	// Call the TouchAsSudo function to attempt creating the file.
	success, err := filex.TouchAsSudo(testFilePath)
	if err != nil {
		// Log the error if the function failed.
		logx.Error(fmt.Sprintf("❌ Failed to create file with sudo: %v", err))
		os.Exit(1)
	}

	// Log the final result based on the boolean return value.
	if success {
		logx.Info("✅ Successfully created file with sudo.")
	} else {
		logx.Info("❌ Failed to create file with sudo.")
	}
}

func testDeleteFile() {
	var testFilePath = "/tmp/mx"

	// Defer a cleanup function to remove the test file regardless of the outcome.
	// This ensures the test is idempotent.
	defer func() {
		success, err := corefilex.DeleteFile(testFilePath)
		if err != nil {
			// handle generic error explicitly: unexpected failure
			log.Printf("❌ Failed to clean up test file: %v", err)
		} else if success {
			log.Println("✅ Successfully cleaned up test file.")
		} else {
			// function explicitly returned 'false' without an error: expected outcome
			log.Println("❌ Failed to clean up test file (DeleteFile returned false).")
		}
	}()

	// // Create a file for the test.
	// if err := os.WriteFile(testFilePath, []byte("test content"), 0644); err != nil {
	// 	log.Printf("❌ Failed to create test file: %v", err)
	// 	return
	// }

	// Call the DeleteFile function to attempt deleting the file.
	success, err := corefilex.DeleteFile(testFilePath)
	if err != nil {
		// Log the error if the function failed.
		log.Printf("❌ Failed to delete file: %v", err)
		os.Exit(1)
	}

	// Log the final result based on the boolean return value.
	if success {
		log.Println("✅ Successfully deleted file.")
	} else {
		log.Println("❌ Failed to delete file.")
	}
}
