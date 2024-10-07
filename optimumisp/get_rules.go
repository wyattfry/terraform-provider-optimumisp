package optimumisp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetPortForwardingRulesResponse struct {
	RouterData RouterData `json:"routerData"`
}
type PortForwardingHostID struct {
	IPConnectionIndex int `json:"ipConnectionIndex"`
	DeviceIndex       int `json:"deviceIndex"`
	ConnectionIndex   int `json:"connectionIndex"`
}
type HostID struct {
	PortForwardingHostID PortForwardingHostID `json:"portForwardingHostId"`
}
type PortForwardingRuleID struct {
	Name     string `json:"name"`
	UniqueID string `json:"uniqueId"`
}
type IPAddress struct {
	IPAddress     string `json:"ipAddress"`
	IPBitLength   any    `json:"ipBitLength"`
	IPV6Address   any    `json:"ipV6Address"`
	Status        any    `json:"status"`
	IPAddressType any    `json:"ipAddressType"`
}
type InternalHost struct {
	IPAddress  IPAddress `json:"ipAddress"`
	Name       any       `json:"name"`
	MacAddress any       `json:"macAddress"`
}
type ExternalPorts struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
type PortForwardings struct {
	Index         int            `json:"index"`
	ExternalPorts *ExternalPorts `json:"externalPorts"`
	InternalPort  int            `json:"internalPort"`
	Protocol      string         `json:"protocol"`
}
type PortForwardingRule struct {
	PortForwardingRuleID PortForwardingRuleID `json:"portForwardingRuleId"`
	Enabled              bool                 `json:"enabled"`
	InternalHost         *InternalHost        `json:"internalHost"`
	Action               string               `json:"action"`
	PortForwardings      []PortForwardings    `json:"portForwardings"`
}
type PortForwardingRules struct {
	PortForwardingRule PortForwardingRule `json:"portForwardingRule"`
}
type RouterData struct {
	Connected           bool                  `json:"connected"`
	DataStatus          string                `json:"dataStatus"`
	HostID              HostID                `json:"hostId"`
	PortForwardingRules []PortForwardingRules `json:"portForwardingRules"`
}

// getPortForwardingRules fetches the current port forwarding rules from the router
func (c *Client) GetPortForwardingRules() ([]PortForwardingRules, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://www.optimum.net/api/user/services/v1/user/router/getPortForwardingRules", nil)
	if err != nil {
		return []PortForwardingRules{}, fmt.Errorf("error creating HTTP request: %v", err)
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
		return []PortForwardingRules{}, fmt.Errorf("error sending request to get port forwarding rules: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []PortForwardingRules{}, fmt.Errorf("error reading response body: %v", err)
	}
	response, err := unmarshalGetPortForwardingRulesResponse(body)
	if err != nil {
		return []PortForwardingRules{}, fmt.Errorf("error reading response body: %v", err)
	}
	return response.RouterData.PortForwardingRules, nil
}

func unmarshalGetPortForwardingRulesResponse(raw []byte) (GetPortForwardingRulesResponse, error) {
	// Unmarshal the response
	var data GetPortForwardingRulesResponse
	err := json.Unmarshal(raw, &data)
	if err != nil {
		return GetPortForwardingRulesResponse{}, fmt.Errorf("error unmarshaling port forwarding rules: %v", err)
	}
	fmt.Printf("Unmarshalled %d port forwarding rules\n", len(data.RouterData.PortForwardingRules))

	return data, nil
}
