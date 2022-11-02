package ast

import "encoding/json"

type Conditional struct {
	Fact     string      `json:"fact"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type Condition struct {
	Any []Conditional `json:"any"`
	All []Conditional `json:"all"`
}

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type Rule struct {
	Condition Condition `json:"condition"`
	Event     Event     `json:"event"`
}

func ParseJSON(j string) *Rule {
	var rule *Rule
	if err := json.Unmarshal([]byte(j), &rule); err != nil {
		panic("expected valid JSON")
	}
	return rule
}
