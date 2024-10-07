package optimumisp

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_unmarshal(t *testing.T) {

	raw := `{
    "routerData": {
        "connected": true,
        "dataStatus": "CURRENT",
        "hostId": {
            "portForwardingHostId": {
                "ipConnectionIndex": 1,
                "deviceIndex": -1,
                "connectionIndex": -1
            }
        },
        "portForwardingRules": [
            {
                "portForwardingRule": {
                    "portForwardingRuleId": {
                        "name": "rulename",
                        "uniqueId": "rulename_MAGIC_EXTENSION#@$"
                    },
                    "enabled": true,
                    "internalHost": {
                        "ipAddress": {
                            "ipAddress": "10.0.0.1",
                            "ipBitLength": null,
                            "ipV6Address": null,
                            "status": null,
                            "ipAddressType": null
                        },
                        "name": null,
                        "macAddress": null
                    },
                    "action": "NONE",
                    "portForwardings": [
                        {
                            "index": 1,
                            "externalPorts": {
                                "start": 80,
                                "end": 80
                            },
                            "internalPort": 80,
                            "protocol": "TCP"
                        }
                    ]
                }
            }
        ]
    }
}`
	input := string(regexp.MustCompile(`\s`).ReplaceAll([]byte(raw), []byte("")))

	got, err := unmarshalGetPortForwardingRulesResponse([]byte(input))

	if err != nil {
		t.Errorf("Err should be nil and instead is %v\n", err)
	}

	if len(got.RouterData.PortForwardingRules) != 1 {
		t.Errorf("Wrong length wanted 1 got %d\n", len(got.RouterData.PortForwardingRules))
	}

	if !got.RouterData.Connected {
		t.Errorf("Expected true but got false\n")
	}

	gotIp := got.RouterData.PortForwardingRules[0].PortForwardingRule.InternalHost.IPAddress.IPAddress

	if gotIp != "10.0.0.1" {
		t.Errorf("Expected Address 10.0.0.1 but got %s\n", gotIp)
	}
}

func Test_SerializeDeleteRequest(t *testing.T) {

	input := `{
"portForwardingRule": {
	"portForwardingRuleId": {
		"name": "rulename"
	},
	"action": "DELETE",
	"portForwardings": [
		{
			"index": 1
		}
	]
}
}`
	want := string(regexp.MustCompile(`\s`).ReplaceAll([]byte(input), []byte("")))
	json, _ := SerializeDeleteRequest("rulename", []int{1})
	got := string(json)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got \n%v\n but wanted \n%v\n", got, want)
	}
}
