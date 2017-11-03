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
	nodeConfig := ltm.Node{
		Name:    "test-node-1",
		Address: "1.1.1.1",
	}
	if err := ltmClient.Node().Create(nodeConfig); err != nil {
		log.Fatal(err)
	}

	// Node 2
	nodeConfig = ltm.Node{
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

func Example_transaction() {
	f5Client, err := f5.NewBasicClient("https://192.168.10.40", "admin", "admin")
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

	// Create a HTTP monitor
	log.Print("Create a HTTP monitor")

	monitorConfig := ltm.MonitorHTTPConfig{
		Name: "http_monitor_" + tx.TransactionID(),
		Send: "GET / HTTP/1.0\r\n\r\n",
		Recv: "Hello",
	}

	if err := ltmClient.MonitorHTTP().Create(monitorConfig); err != nil {
		log.Fatal(err)
	}

	// Create a Pool
	log.Print("Create a pool")

	poolConfig := ltm.Pool{
		Name:    "pool_" + tx.TransactionID(),
		Monitor: "/Common/http_monitor_" + tx.TransactionID(),
		Members: []string{"10.1.10.10:80", "10.1.10.11:80"},
	}

	if err := ltmClient.Pool().Create(poolConfig); err != nil {
		log.Fatal(err)
	}

	// Create a Virtual Server
	log.Print("Create a Virtual Server")

	vsConfig := ltm.VirtualServer{
		Name:        "vs_http_" + tx.TransactionID(),
		Destination: "10.1.20.130:80",
		IPProtocol:  "tcp",
		Pool:        "pool_" + tx.TransactionID(),
		SourceAddressTranslation: ltm.SourceAddressTranslation{
			Type: "automap",
		},
		Profiles: []ltm.Profile{
			{
				Name:    "tcp-mobile-optimized",
				Context: "all",
			},
			{
				Name: "http",
			},
		},
	}

	if err := ltmClient.Virtual().Create(vsConfig); err != nil {
		log.Fatal(err)
	}

	// Commit to make the changes persistent.
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	// Output:
}
