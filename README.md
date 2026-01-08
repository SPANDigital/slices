# 🚀 slices

A powerful, type-safe Go library providing generic utility functions for working with slices.

[![Open in Dev Containers](https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/SPANDigital/slices)
![Develop Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/slices/go.yml?branch=develop&label=develop)
![Main Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/slices/go.yml?branch=main&label=main)
![Tag](https://img.shields.io/github/v/tag/SPANDigital/slices)

## Why use this library?

Working with slices is fundamental to Go programming, but the standard library doesn't provide many helper functions for common slice operations. This library fills that gap by providing type-safe, generic functions for filtering, mapping, flattening, grouping, and more.

Built with **Go generics**, this library offers:
- 🔒 **Type safety** - Catch errors at compile time, not runtime
- ⚡ **Zero dependencies** - Just pure Go (except for testing)
- 🧪 **BDD tested** - Comprehensive test coverage using Cucumber/Godog
- 📦 **Easy to use** - Simple, intuitive API

## Installation

```bash
go get github.com/spandigital/slices
```

**Requirements:** Go 1.24.0 or later

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/spandigital/slices"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

    // Filter even numbers
    evens := slices.Filter(numbers, func(n int) bool {
        return n % 2 == 0
    })
    fmt.Println(evens) // [2, 4, 6, 8]

    // Double all numbers
    doubled := slices.Map(numbers, func(n int) int {
        return n * 2
    })
    fmt.Println(doubled) // [2, 4, 6, 8, 10, 12, 14, 16]

    // Get unique values
    unique := slices.Unique([]int{1, 2, 2, 3, 3, 3})
    fmt.Println(unique) // [1, 2, 3]
}
```

## Available Functions

### Filtering & Selection

#### `Filter[S ~[]V, V any](s S, predicate func(V) bool) []V`

Filter a slice based on a predicate function.

```go
evens := slices.Filter([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
    return v % 2 == 0
})
// Result: [2, 4, 6]
```

#### `FilterNil[T any](in []*T) []*T`

Remove nil pointer values from a slice of pointers.

```go
type Person struct{ Name string }
people := []*Person{
    {Name: "Alice"},
    nil,
    {Name: "Bob"},
    nil,
}
filtered := slices.FilterNil(people)
// Result: [{Name: "Alice"}, {Name: "Bob"}]
```

#### `RemoveNil[T any](in []T) []T`

Remove any nil-able values (pointers, interfaces, slices, maps, channels, functions) from a slice.

```go
values := []any{1, nil, "hello", nil, 3.14}
clean := slices.RemoveNil(values)
// Result: [1, "hello", 3.14]
```

#### `Contains[S ~[]E, E comparable](s S, v E) bool`

Check if a slice contains a specific value.

```go
hasValue := slices.Contains([]int{1, 2, 3, 4, 5}, 3) // true
noValue := slices.Contains([]string{"foo", "bar"}, "baz") // false
```

#### `Index[S ~[]E, E comparable](s S, v E) int`

Find the index of a value in a slice. Returns -1 if not found.

```go
idx := slices.Index([]int{1, 2, 3, 4, 5}, 3) // 2
notFound := slices.Index([]string{"foo", "bar"}, "baz") // -1
```

#### `Unique[T comparable](s []T) []T`

Get unique values from a slice (removes duplicates).

```go
unique := slices.Unique([]int{1, 2, 2, 3, 3, 3, 4})
// Result: [1, 2, 3, 4] (order may vary)
```

#### `Intersection[T cmp.Ordered](slices ...[]T) []T`

Find the intersection of multiple slices (values present in all slices).

```go
common := slices.Intersection(
    []int{1, 2, 3, 4, 5},
    []int{3, 4, 5, 6, 7},
    []int{4, 5, 8, 9},
)
// Result: [4, 5]
```

### Transformation

#### `Map[S ~[]V, V any, E any](s S, extract func(V) E) []E`

Transform each element in a slice using a mapping function.

```go
doubled := slices.Map([]int{1, 2, 3, 4}, func(v int) int {
    return v * 2
})
// Result: [2, 4, 6, 8]

lengths := slices.Map([]string{"a", "ab", "abc"}, func(s string) int {
    return len(s)
})
// Result: [1, 2, 3]
```

#### `SyncMap[S ~[]V, V any, E any](ctx context.Context, s S, extract func(context.Context, V) (E, error)) ([]E, error)`

Concurrently map over a slice using goroutines. All operations run in parallel.

```go
import "context"

urls := []string{"https://api.example.com/1", "https://api.example.com/2"}
results, err := slices.SyncMap(context.Background(), urls, func(ctx context.Context, url string) (string, error) {
    // Fetch data from URL concurrently
    return fetchData(ctx, url)
})
```

#### `MapFrom[S ~[]I, I any, K comparable, V any](s S, extractKey func(I) K, extractValue func(I) V) map[K]V`

Convert a slice to a map by extracting keys and values from each element.

```go
type User struct {
    ID   int
    Name string
}

users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
userMap := slices.MapFrom(users,
    func(u User) int { return u.ID },
    func(u User) string { return u.Name },
)
// Result: map[int]string{1: "Alice", 2: "Bob"}
```

#### `Flatten[T any](s [][]T) []T`

Flatten a 2-dimensional slice into a 1-dimensional slice.

```go
nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
flat := slices.Flatten(nested)
// Result: [1, 2, 3, 4, 5, 6]
```

**Note:** Also available: `Flatten3`, `Flatten4` for 3D and 4D slices.

### Grouping & Organization

#### `GroupBy[S ~[]V, V any, K comparable](s S, extract func(V) K) map[K][]V`

Group slice elements by a key extracted from each element.

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {Name: "Alice", Age: 25},
    {Name: "Bob", Age: 30},
    {Name: "Charlie", Age: 25},
}

byAge := slices.GroupBy(people, func(p Person) int {
    return p.Age
})
// Result: map[int][]Person{
//     25: [{Name: "Alice", Age: 25}, {Name: "Charlie", Age: 25}],
//     30: [{Name: "Bob", Age: 30}],
// }
```

#### `GroupByLen[S ~[]V, V any](input S, length int) [][]V`

Split a slice into chunks of a specified length.

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
chunks := slices.GroupByLen(numbers, 3)
// Result: [[1, 2, 3], [4, 5, 6], [7, 8, 9]]

partialChunk := slices.GroupByLen([]int{1, 2, 3, 4, 5}, 2)
// Result: [[1, 2], [3, 4], [5]]
```

### Pagination

#### `Page[V any](s []V, pageSize int, pageIndex int) []V`

Get a specific page from a slice (zero-indexed).

```go
items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
page1 := slices.Page(items, 3, 0) // [1, 2, 3]
page2 := slices.Page(items, 3, 1) // [4, 5, 6]
page4 := slices.Page(items, 3, 3) // [10]
```

#### `NumPages[V any](s []V, pageSize int) int`

Calculate the number of pages needed for a given page size.

```go
items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
totalPages := slices.NumPages(items, 3) // 4
```

### Appending

#### `AppendNotNil[T any](in []*T, values ...*T) []*T`

Append only non-nil pointers to a slice.

```go
type Item struct{ Value int }
items := []*Item{{Value: 1}}
items = slices.AppendNotNil(items, &Item{Value: 2}, nil, &Item{Value: 3})
// Result: [{Value: 1}, {Value: 2}, {Value: 3}]
```

## Testing

This library uses **Behavior-Driven Development (BDD)** with [Cucumber/Godog](https://github.com/cucumber/godog). All functions are tested with comprehensive scenarios written in Gherkin.

```bash
# Run all tests
go test -v ./...

# Run specific function tests
go test -v -run TestFeatures/Filter
```

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

This project uses:
- **Conventional Commits** for commit messages
- **Pre-commit hooks** for code quality (go-fmt, golangci-lint, tests)
- **BDD testing** with Godog

## License

See [LICENSE](LICENSE) file for details.

## Maintainers

This library is maintained by [SPAN Digital](https://github.com/SPANDigital).
