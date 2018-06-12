package ruletypes

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// The result of listing the qos rule types
type RuleType struct {
	Type string `json:"type"`
}

type ListRuleTypesPage struct {
	pagination.SinglePageBase
}

func (r ListRuleTypesPage) IsEmpty() (bool, error) {
	v, err := ExtractRuleTypes(r)
	return len(v) == 0, err
}

func ExtractRuleTypes(r pagination.Page) ([]RuleType, error) {
	var s struct {
		RuleTypes []RuleType `json:"rule_types"`
	}

	err := (r.(ListRuleTypesPage)).ExtractInto(&s)
	return s.RuleTypes, err
}

// Result from getting details of a rule type
type GetRuleTypeResult struct {
	gophercloud.Result
}

// Extract the body for rule type details
func (g GetRuleTypeResult) Extract() (*RuleTypeSpec, error) {
	var r RuleTypeSpec
	err := g.ExtractInto(&r)
	return &r, err
}

// The top level response body for GetRuleType
type RuleTypeSpec struct {
	// The drivers for this RuleType
	Drivers []RuleDriver `json:"drivers"`
	// The name of the RuleType
	Type    string       `json:"type"`
}

// A driver for some RuleType
type RuleDriver struct {
	Name                string                `json:"name"`
	SupportedParameters []RuleDriverParameter `json:"supported_parameters"`
}

// A parameter for some RuleDriver
type RuleDriverParameter struct {
	// The name of the parameter
	Name   string `json:"parameter_name"`
	// The type of the parameter (values can be choices, or range)
	Type   string `json:"parameter_type"`
	// The values specified for this parameter
	Values interface{} `json:"parameter_values"`
}

