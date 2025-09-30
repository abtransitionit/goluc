/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/go-resty/resty/v2"
)

func test02(ctx context.Context, logger logx.Logger) {
	SaId, err := ovh.GetSaId()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	SaSecret, err := ovh.GetSaSecret()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}

	// creata client
	client := resty.New()

	domain := "www.ovh.com"
	endpoint := "/auth/oauth2/token"
	url := fmt.Sprintf("https://%s%s", domain, endpoint)

	// Send request
	resp, err := client.R().
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     SaId,
			"client_secret": SaSecret,
			"scope":         "all",
		}).
		SetResult(&ovh.AccessToken{}). // auto-unmarshal JSON into struct
		Post(url)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	// Access the parsed response
	result := resp.Result().(*ovh.AccessToken)
	fmt.Println("Access Token:", result.Token)
	fmt.Println("Expires In:", result.ExpiresIn)

}

// cid="EU.ad5a4d2d8166e832"
// cs="a2ae0f8dafc59ac1a257f9624bee315e"
// domain="www.ovh.com"
// endpoint="/auth/oauth2/token"
// url="https://${domain}${endpoint}"
// verb="POST"
// curl --request ${verb} \
//   --url "${url}" \
//   --header 'content-type: application/x-www-form-urlencoded' \
//   --data scope=all \
//   --data client_id="${cid}" \
//   --data client_secret="${cs}" \
//   --data grant_type=client_credentials
