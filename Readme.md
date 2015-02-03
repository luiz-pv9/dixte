Dixte Analytics
===============

Events and profiles analytics written in Go using PostgreSQL. The goal of this
project is 

### Overview


### Architecture

### Packages

#### `jsonvalidation`

**Why?**
Go's JSON unmarshal works great with struct type, and you can validate the arguments by simply checking some `nil` values or empty strings to see if there are some missing fields. But when the structure is unkown, (e.g event and profile properties) it becomes hard to validate some unauthorized values, and easy to forget some case.

**Example**
```go
	// The syntax API may be a little overwhelming, so let's see each part
	// AnyObjectByRules is a function that returns a JsonValidator (which is also a function)
	matcher := AnyObjectByRules([]*JsonKeyValuePairValidator{
		&JsonKeyValuePairValidator{
			ExactString("name"), AnyObjectByRules([]*JsonKeyValuePairValidator{
				&JsonKeyValuePairValidator{
					ExactString("first"), AnyString,
				},
				&JsonKeyValuePairValidator{
					ExactString("last"), AnyString,
				},
			}),
		},
		// The 
		&JsonKeyValuePairValidator{
			ExactString("age"), AnyNumber,
		},
	})

	// Now matcher is a function that receives one argument: the object we want
	// to check.

	// This is equivalent to:
	// {
	//     "name": {
	//         "first": "Luiz",
	//         "last": "Vasconcellos"
	//     },
	//     "age": 20
	// }
	obj := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Luiz",
			"last":  "Vasconcellos",
		},
		"age": float64(20),
	}
```

* `databasemanager`
* `environment`
* `properties`
* `apps`

* `events` and `profiles`
	
	In development