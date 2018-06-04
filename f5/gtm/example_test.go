// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm_test

import (
	"log"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/e-XpertSolutions/f5-rest-client/f5/gtm"
)

func Example() {
	// setup F5 BigIP client
	f5Client, err := f5.NewBasicClient("https://url-to-bigip", "user", "password")
	if err != nil {
		log.Fatal(err)
	}

	// setup client for the GTM/DNS API
	gtmClient := gtm.New(f5Client)

	// query the /gtm/datacenter API
	configList, err := gtmClient.Datacenter().ListAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(configList)
}
