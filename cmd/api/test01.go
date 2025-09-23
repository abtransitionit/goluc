/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package api

import (
	"fmt"
	"os"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
)

func test01(logger logx.Logger) {

	accessToken, err := ovh.GetAccessTokenFromFile()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	clientSecret, err := ovh.GetSaSecret()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	clientId, err := ovh.GetSaId()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}

	fmt.Println("Access Token:", accessToken)
	fmt.Println("Client Secret:", clientSecret)
	fmt.Println("Client Id:", clientId)

}

// // Create client
// basedApiEndpoint := "https://jsonplaceholder.typicode.com"
// client := apicli.NewClient(basedApiEndpoint, nil, nil)
// _ = client.Do(
// 	context.Background(),
// 	"GET",
// 	"/posts",
// 	nil, // body
// 	nil, // output
// 	logger,
// 	map[string]string{"page": "1"},          // query params
// 	map[string]string{"X-Api-Key": "token"}, // headers
// )
