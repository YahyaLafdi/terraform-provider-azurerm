package services

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkRuleSet struct {
	Bypass  *SearchBypass `json:"bypass,omitempty"`
	IPRules *[]IPRule     `json:"ipRules,omitempty"`
}
