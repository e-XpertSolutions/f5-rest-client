package f5

import (
	"encoding/json"
	"fmt"
)

const (
	QKViewPath = "/mgmt/cm/autodeploy/qkview"
)

type QKViewResponse struct {
	// Unique ID to identify uniquely the qkview.
	ID string `json:"id"`

	// Name of the file in which the qkview is saved.
	Name string `json:"name"`

	//Status of qkview. Possible values are:
	//	- SUCCEDED
	//	- FAILED
	//	- IN_PROGRESS
	Status string `json:"status"`
}

func (c *Client) GenerateQKView(filename string) (*QKViewResponse, error) {
	data := map[string]interface{}{
		"name": filename,
	}
	resp, err := c.SendRequest("POST", QKViewPath, data)
	if err != nil {
		return nil, fmt.Errorf("error while requesting qkview generation: %w", err)
	}
	defer resp.Body.Close()

	if err := c.ReadError(resp); err != nil {
		return nil, fmt.Errorf("qkview api returned an error: %w", err)
	}

	var qkviewResp QKViewResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&qkviewResp); err != nil {
		return nil, fmt.Errorf("qkview api returned a malformed json response: %w", err)
	}
	return &qkviewResp, nil
}

func (c *Client) ListQKViews() ([]QKViewResponse, error) {
	resp, err := c.SendRequest("GET", QKViewPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error while requesting qkview list: %w", err)
	}
	defer resp.Body.Close()

	if err := c.ReadError(resp); err != nil {
		return nil, fmt.Errorf("qkview api returned an error: %w", err)
	}

	qkviewResp := struct {
		Items []QKViewResponse `json:"items"`
	}{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&qkviewResp); err != nil {
		return nil, fmt.Errorf("qkview api returned a malformed json response: %w", err)
	}
	return qkviewResp.Items, nil
}

func (c *Client) CheckQKView(id string) (*QKViewResponse, error) {

	resp, err := c.SendRequest("GET", QKViewPath+"/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("error while checking qkview: %w", err)
	}
	defer resp.Body.Close()

	if err := c.ReadError(resp); err != nil {
		return nil, fmt.Errorf("qkview api returned an error: %w", err)
	}

	var qkviewResp QKViewResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&qkviewResp); err != nil {
		return nil, fmt.Errorf("qkview api returned a malformed json response: %w", err)
	}
	return &qkviewResp, nil
}

func (c *Client) DeleteQKView(id string) (*QKViewResponse, error) {
	resp, err := c.SendRequest("DELETE", QKViewPath+"/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("error while checking qkview: %w", err)
	}
	defer resp.Body.Close()

	if err := c.ReadError(resp); err != nil {
		return nil, fmt.Errorf("qkview api returned an error: %w", err)
	}

	var qkviewResp QKViewResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&qkviewResp); err != nil {
		return nil, fmt.Errorf("qkview api returned a malformed json response: %w", err)
	}
	return &qkviewResp, nil
}
