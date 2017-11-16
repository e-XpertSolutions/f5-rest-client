package f5

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// ErrNoTransaction is the error returned when a function related to
// transaction management is called when there is no active transaction.
var ErrNoTransaction = errors.New("no active transaction")

type Transaction struct {
	TransID          int64  `json:"transId"`
	ValidateOnly     bool   `json:"validateOnly"`
	ExecutionTimeout int64  `json:"executionTimeout"`
	SelfLink         string `json:"selfLink"`
	State            string `json:"state"`
	TimeoutSeconds   int64  `json:"timeoutSeconds"`
	AsyncExecution   bool   `json:"asynExecution"`
	FailureReason    string `json:"failureReason"`
	Kind             string `json:"kind"`
}

// Begin starts a transaction.
func (c *Client) Begin() (*Client, error) {
	// Send HTTP request to the API responsible for creating a new transaction.
	resp, err := c.SendRequest("POST", "/mgmt/tm/transaction", map[string]interface{}{})
	if err != nil {
		return nil, errors.New("cannot create request for starting a new transaction: " + err.Error())
	}
	defer resp.Body.Close()

	// Decode and validate transaction creation response.
	var tx struct {
		TransID int64  `json:"transId"`
		State   string `json:"state"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&tx); err != nil {
		return nil, errors.New("cannot read transaction creation response: " + err.Error())
	}
	if tx.State != "STARTED" {
		return nil, fmt.Errorf("invalid transaction sate %q; want %q", tx.State, "STARTED")
	}

	// Create a new client from the current one, but with a transaction ID.
	newClient := c.clone()
	newClient.txID = strconv.FormatInt(tx.TransID, 10)

	return newClient, nil
}

// TransactionID returns the ID of the current transaction. If there is no
// active transaction, an empty string is returned.
func (c *Client) TransactionID() string {
	return c.txID
}

// TransactionState returns the state of the current transaction. If there is no
// active transaction, ErrNoTransaction is returned.
func (c *Client) TransactionState() (*Transaction, error) {
	if c.txID == "" {
		return nil, ErrNoTransaction
	}
	var tx Transaction
	if err := c.ReadQuery("/mgmt/tm/transaction/"+c.txID, &tx); err != nil {
		return nil, errors.New("cannot get current transaction state: " + err.Error())
	}
	return &tx, nil
}

// Rollback aborts the current transaction. If there is no active transaction,
// ErrNoTransaction is returned.
func (c *Client) Rollback() error {
	if c.txID == "" {
		return nil
	}
	txID := c.txID
	c.txID = ""
	if err := c.ModQuery("DELETE", "/mgmt/tm/transaction/"+txID, nil); err != nil {
		c.txID = txID
		return errors.New("cannot rollback current transaction: " + err.Error())
	}
	return nil
}

// Commit commits the transaction.
func (c *Client) Commit() error {
	if c.txID == "" {
		return errors.New("no transaction started")
	}

	txID := c.txID
	c.txID = ""

	data := map[string]interface{}{"state": "VALIDATING"}
	resp, err := c.SendRequest("PATCH", "/mgmt/tm/transaction/"+txID, data)
	if err != nil {
		return errors.New("cannot commit transaction: " + err.Error())
	}
	defer resp.Body.Close()

	return nil
}
