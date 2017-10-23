// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apm_test

import (
	"log"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/e-XpertSolutions/f5-rest-client/f5/apm"
)

func Example() {
	// setup F5 BigIP client
	f5Client, err := f5.NewBasicClient("https://url-to-bigip", "user", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Start new transaction.
	tx, err := f5Client.Begin()
	if err != nil {
		log.Fatal(err)
	}

	apmClient := apm.New(tx)

	oauthAppConfig := &apm.OAuthClientAppConfig{
		Name:          "my_oauth_app_client",
		AppName:       "my_oauth_app_client",
		AuthType:      "secret",
		GrantPassword: "enabled",
		Scopes:        "scope1",
	}

	oauthAppConfig, err = apmClient.OAuthClientApp().Create(*oauthAppConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = apmClient.OAuthProfile().AppendAppClient("my_oauth_profile", oauthAppConfig.Name)
	if err != nil {
		log.Fatal(err)
	}

	err = apmClient.AccessProfile().ApplyPolicy("my_access_policy")
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

}
