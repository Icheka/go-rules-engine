package ast

import (
	"fmt"
	"testing"
)

func TestMapify(t *testing.T) {
	// Mapify converts a struct to a map[string]interface{}
	type S struct {
		Key string
	}
	s := S{
		Key: "value",
	}
	m := Mapify(s)

	if m["Key"] == nil || m["Key"] != "value" {
		t.Fatalf("expected m[\"Key\"] to be %s, got %s", s.Key, m["Key"])
	}
}

func TestParseJSON(t *testing.T) {
	j := `{
		"condition": {
			"any": [{
				"fact": "myVar",
				"operator": "eq",
				"value": "hello world"
			}]
		},
		"event": {
			"type": "result",
			"payload": {
				"data": {
					"say": "Hello World!"
				}   
			}
		}
	}`

	rule := ParseJSON(j)

	if fmt.Sprintf("%T", rule) != "*ast.Rule" {
		t.Fatalf("expected rule to be *ast.Rule, got %T", rule)
	}
}
