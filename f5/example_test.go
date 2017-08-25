package f5_test

import (
	"log"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/e-XpertSolutions/f5-rest-client/f5/ltm"
)

func ExampleClient_Begin() {
	f5Client, err := f5.NewBasicClient("https://127.0.0.1", "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
	f5Client.DisableCertCheck()

	// Start new transaction.
	tx, err := f5Client.Begin()
	if err != nil {
		log.Fatal(err)
	}

	ltmClient := ltm.New(tx)

	// Node 1
	nodeConfig := ltm.NodeConfig{
		Name:    "test-node-1",
		Address: "1.1.1.1",
	}
	if err := ltmClient.Node().Create(nodeConfig); err != nil {
		log.Fatal(err)
	}

	// Node 2
	nodeConfig = ltm.NodeConfig{
		Name:    "test-node-2",
		Address: "2.2.2.2",
	}
	if err := ltmClient.Node().Create(nodeConfig); err != nil {
		log.Fatal(err)
	}

	// Commit to make the changes persistent.
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
