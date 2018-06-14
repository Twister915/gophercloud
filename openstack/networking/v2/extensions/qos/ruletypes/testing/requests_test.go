package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/qos/ruletypes"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListRuleTypes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListRuleTypesResponse)
	})

	page, err := ruletypes.ListRuleTypes(fake.ServiceClient()).AllPages()
	if err != nil {
		t.Errorf("Failed to list rule types pages: %v", err)
		return
	}

	rules, err := ruletypes.ExtractRuleTypes(page)
	if err != nil {
		t.Errorf("Failed to list rule types: %v", err)
		return
	}

	expected := []ruletypes.RuleType{{Type: "bandwidth_limit"}, {Type: "dscp_marking"}, {Type: "minimum_bandwidth"}}
	th.AssertDeepEquals(t, expected, rules)
}

func TestGetRuleTypes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/qos/rule-types/bandwidth_limit", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetRuleTypeResponse)
	})

	ruleType, err := ruletypes.GetRuleType(fake.ServiceClient(), "bandwidth_limit").Extract()
	if err != nil {
		t.Errorf("Failed to get rule type: %v", err)
		return
	}

	expected := &ruletypes.RuleTypeSpec{
		Type: "bandwidth_limit",
		Drivers: []ruletypes.RuleDriver{
			{
				Name: "openvswitch",
				SupportedParameters: []ruletypes.RuleDriverParameter{
					{
						Name: "max_kbps",
						Type: "range",
						Values: map[string]interface{}{
							"end": 2147483647.0,
							"start": 0.0,
						},
					},
					{
						Name: "direction",
						Type: "choices",
						Values: []string{"ingress", "egress"},
					},
					{
						Name: "max_burst_kbps",
						Type: "range",
						Values: map[string]interface{}{
							"end": 2147483647.0,
							"start": 0.0,
						},
					},
				},
			},
			{
				Name: "linuxbridge",
				SupportedParameters: []ruletypes.RuleDriverParameter{
					{
						Name: "max_kbps",
						Type: "range",
						Values: map[string]interface{}{
							"end": 2147483647.0,
							"start": 0.0,
						},
					},
					{
						Name: "direction",
						Type: "choices",
						Values: []string{"ingress", "egress"},
					},
					{
						Name: "max_burst_kbps",
						Type: "range",
						Values: map[string]interface{}{
							"end": 2147483647.0,
							"start": 0.0,
						},
					},
				},
			},
		},
	}

	th.AssertDeepEquals(t, expected, ruleType)
}