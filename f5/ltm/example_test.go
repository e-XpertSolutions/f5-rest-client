// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm_test

import (
	"log"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/e-XpertSolutions/f5-rest-client/f5/ltm"
)

func Example() {
	// setup F5 BigIP client
	f5Client, err := f5.NewBasicClient("https://url-to-bigip", "user", "password")
	if err != nil {
		log.Fatal(err)
	}

	// setup client for the LTM API
	ltmClient := ltm.New(f5Client)

	// query the /ltm/virtual API
	vsConfigList, err := ltmClient.Virtual().ListAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(vsConfigList)

	node1 := ltm.Node{
		Address: "1.1.1.1",
		Name:    "my-node-1",
	}

	err = ltmClient.Node().Create(node1)
	if err != nil {
		log.Fatal(err)
	}
}
