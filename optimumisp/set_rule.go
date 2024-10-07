package optimumisp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// type SetPortForwardingRuleRequest struct {
// 	PortForwardingRule struct {
// 		PortForwardingRuleID struct {
// 			Name string `json:"name"`
// 		} `json:"portForwardingRuleId"`
// 		Enabled      bool `json:"enabled"`
// 		InternalHost struct {
// 			IPAddress struct {
// 				IPAddress string `json:"ipAddress"`
// 			} `json:"ipAddress"`
// 		} `json:"internalHost"`
// 		Action          string `json:"action"`
// 		PortForwardings []struct {
// 			Index         int `json:"index"`
// 			ExternalPorts struct {
// 				Start string `json:"start"`
// 				End   string `json:"end"`
// 			} `json:"externalPorts"`
// 			InternalPort string `json:"internalPort"`
// 			Protocol     string `json:"protocol"`
// 		} `json:"portForwardings"`
// 	} `json:"portForwardingRule"`
// }

// setPortForwardingRule sends a POST request to set a port forwarding rule
func (c *Client) SetPortForwardingRule(ruleName string, indexes []int, action string) error {
	// Construct the port forwarding rule request body
	request := PortForwardingRule{}

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://www.optimum.net/api/user/services/v1/user/router/setPortForwardingRule", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	for _, c := range c.cookies {
		req.AddCookie(&http.Cookie{
			Name:   c.Name,
			Value:  c.Value,
			Domain: c.Domain,
		})
	}

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending port forwarding request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("port forwarding request failed with status: %v", resp.Status)
	}

	return nil
}
