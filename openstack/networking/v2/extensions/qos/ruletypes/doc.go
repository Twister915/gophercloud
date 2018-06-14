/*
Package ruletypes contains functionality for working with Neutron 'quality of service' rule-type resources.

Example: You can list rule-types in the following way:

	page, err := ruletypes.ListRuleTypes(client).AllPages()
	if err != nil {
		return
	}

	rules, err := ruletypes.ExtractRuleTypes(page)
	if err != nil {
		return
	}

	fmt.Printf("%v <- Rule Types\n", rules)


If you'd like to get the details on a specific rule-type, you can do something like this:

	ruleType, err := ruletypes.GetRuleType(client, "bandwidth_limit").Extract()
	if err != nil {
		t.Errorf("Failed to get rule type: %v", err)
		return
	}

	fmt.Printf("%v <- bandwidth_limit\n", ruleType)

*/
package ruletypes
