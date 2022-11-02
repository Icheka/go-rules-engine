![The right way to Go](https://raw.githubusercontent.com/Icheka/go-rules-engine/ef41df8a2a2effdb340fc2d352d673e9ca82ad50/gopher-ladder.svg "The right way to Go")

### **"The right way to Go"**

<br />

# Go-Rules-Engine

### A JSON-based rule engine, written in Go.

Go-Rules-Engines is a powerful, lightweight, un-opinionated rules engine written in Go. Rules are expressed in simple JSON, and can be stored anywhere (in standalone files, source code, or as data stored in databases), and edited by anyone (even persons with no technical skill).

## Features

- **Deterministic**: uses JSON as an AST (Abstract Syntax Tree) from which to draw inferences and publish reactive events
- **Supports "any" and "all" context operators**
- **Blazing fast**
- **Secure and sandboxed** - JSON AST is never evaluated
- **Easily extensible** - Perfect for building larger expert systems via composition
- **Easily modifiable** - JSON AST can be modified by anybody -- no technical expertise required

## Installation

Works best with Go >=1.8.

```bash
go get github.com/icheka/go-rules-engine
```

## Synopsis

Go-Rules-Engine is build around the concept of Rules. A rule is an expression of business logic as a combination of one or more conditions and an event to be fired when those conditions are met.

        Go-Rules-Engine
               |
           -----------
          |           |
       Conditions   Event

As an example, a simple rule for a fictional discount engine might be stated as:
"Offer a 10% discount if the customer buys 2 apples". Writing a Rule for this discount is easy enough:

### Conditions

Conditions are groups of statements that are evaluated by Go-Rules-Engine. Evaluating to `true` will cause their corresponding event to be fired. Firing an event, instead of directly executing an action, allows Go-Rules-Engine to remain un-opinionated, leaving full control over results processing in the hands of the engineer. This makes Go-Rules-Engine extremely flexible and easily integratable.

Conditions comprise two parts: `all` and `any`. `all` is used enforce that all statements (enclosed by `all` evaluate to `true`) for the corresponding event to be fired. `any` works a bit differently: it requires just one of its statements to evaluate to `true` for the corresponding event to be fired.

The condition of the discount above will look like:

```json
{
  "condition": {
    "all": [
      {
        "identifier": "applesCount",
        "operator": "=",
        "value": 2
      }
    ]
  }
}
```

### Events

Go-Rules-Engine fires a Rule's event when its Conditions evaluates to true. Events are allowed two properties: `type` and `payload` and they are both up to the engineer to customise.

The event for the discount above could look like:

```json
{
    ...
    "event": {
        "type": "discount",
        "payload": {
            "percentage": 10,
            "item": "apple"
        }
    }
}
```

Thus, the discount Rule can be expressed as:

```json
{
  "condition": {
    "all": [
      {
        "identifier": "applesCount",
        "operator": "=",
        "value": 2
      }
    ]
  },
  "event": {
    "type": "discount",
    "payload": {
      "percentage": 10,
      "item": "apple"
    }
  }
}
```

## Processing Rules

Following the example above, assuming that the discount Rule is stored in the file system, we can process the Rule like so:

```go
package main

import (
    "fmt"
    "os"

    ruleEngine "github.com/Icheka/go-rules-engine/rule_engine"
)

func main() {
    // read discount rule
    jsonBytes, err := os.ReadFile("apple-discount-rule.json")
    if err != nil {
        panic(err)
    }

    // a map[string]interface{} representing a customer's cart at checkout
    // cart contains a key (applesCount) matching the `identifier` in our rule's condition
    cart := map[string]interface{}{
        "applesCount": 3,
        "orangesCount": 5,
        "cookiesCount": 1
    }

    // create a new Rule Engine...
    engine := ruleEngine.New(nil)
    // ... and add the discount rule
    engine.AddRule(string(jsonByres))
    // then process it
    fmt.Printf("%+v", engine.EvaluateRules(cart))
    // [{Type:discount Payload:map[item:apple percentage:10]}]
}
```

## More Complex Rules

A rule for the statement: "player A wins the match if player A has no cards left, or if player B has up to 20 cards left" has two possible paths:

1. Player A has no cards left
2. Player B has up to 20 (i.e greater or equal to 20) cards left

These can be expressed aptly using `any`:

```json
{
  "condition": {
    "any": [
      {
        "identifier": "playerACards",
        "operator": "=",
        "value": 0
      },
      {
        "identifier": "playerBCards",
        "operator": ">=",
        "value": 20
      }
    ]
  },
  "event": {
    "type": "win"
  }
}
```

```go
// [{Type:win Payload:<nil>}]
```

Both `event.type` and `event.payload` are optional and entirely up to the rule creator to specify, provided they are valid JSON structures.

## Credits

Special thanks to [@CacheControl](https://github.com/CacheControl) for his work on [json-rules-engine](https://github.com/CacheControl/json-rules-engine) which inspired this.
