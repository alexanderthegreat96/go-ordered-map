# OrderedMap - A Go Package for Managing Ordered Key-Value Pairs

This package provides a `OrderedMap` data structure in Go that leverages the `linkedhashmap` library from `github.com/emirpasic/gods/maps/linkedhashmap` to maintain key-value pairs in insertion order.

## Features

* **Preserves Insertion Order:** Keys are stored in the order they were added, providing a deterministic way to iterate through the map.
* **JSON Support:** The `FromJSON` and `ToJSON` methods enable easy conversion between `OrderedMap` and JSON format.
* **Efficient Key-Value Operations:** Provides methods for adding, removing, accessing, and iterating over key-value pairs.

## Installation

1. Ensure you have Go installed ([https://go.dev/doc/install](https://go.dev/doc/install)).
2. Use the `go get` command to download the package:

    ```bash
    go get -u github.com/alexanderthegreat96/go-ordered-map
3. Import the package
    ```go
    import (
        "encoding/json"
        "github.com/emirpasic/gods/maps/linkedhashmap"
    )
4. Create a sorted map instance
    ```go
    sm := NewOrderedMap()
5. Start using the functionality
    ```go
    sm.AddPair("name", "John Doe")
    sm.AddPair("age", 30)

    data := map[string]interface{} {
        "is_student": false,
        "gpa": 3.75,
    }
    sm.AddPairs(data)

    // use a json string for the map
    jsonStr := `{"city": "Anytown", "zipcode": "12345"}`
    sm.FromJSON(jsonStr)

    // convert back into json

    jsonStr := sm.ToJSON()
    fmt.Println(jsonStr)


## Licence
MIT Licence