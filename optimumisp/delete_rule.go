package optimumisp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SerializeDeleteRequest(ruleName string, indexes []int) ([]byte, error) {
	var fwds []PortForwardings

	for _, i := range indexes {
		fwds = append(fwds, PortForwardings{
			Index: i,
		})
	}

	rules := PortForwardingRules{
		PortForwardingRule: PortForwardingRule{
			PortForwardingRuleID: PortForwardingRuleID{
				Name: ruleName,
			},
			Action:          "DELETE",
			PortForwardings: fwds,
		},
	}

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(rules)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// setPortForwardingRule sends a POST request to Delete a port forwarding rule
func (c *Client) DeletePortForwardingRule(ruleName string, indexes []int) error {
	// Construct the port forwarding rule request body
	jsonBody, _ := SerializeDeleteRequest(ruleName, indexes)

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
	fmt.Printf("Sending delete request: %s\n\n%v\n", string(jsonBody), req)
	resp, err := client.Do(req)
	fmt.Println("Sent delete request. Response:", resp)
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
