# slices
SPAN Digital Slices

[![Open in Dev Containers](https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/SPANDigital/slices)
![Develop Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/slices/go.yml?branch=develop&label=develop)
![Main Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/slices/go.yml?branch=main&label=main)
![Release status](https://img.shields.io/github/v/tag/SPANDigital/slices)

## Usage

Generic functions which operate on Go slices.

### func Filter[S ~[]V, V any](s S, predicate func(V V) bool) []V

Filter a slice with a predicate

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	
	b := slices.Filter(a, func(v int) bool {
        return v % 2 == 0 // even
	})
	
	println(b) // 2, 4, 6, 8
}
```

### func Flatten[T any](s [][]T) (flattened []T) and 3+ dimensional variants

This function flats a 2 dimensional slice into a 1 dimensional slice,
There are variants `Flatten3`, `Flatten4` for 3 and 4 dimensions respectively.

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	a := [][]int{{1, 2, 3, 4}, {5, 6}, {7, 8}}
	
	b := slices.Flatten(a)
	
	println(b) // 1, 2, 3, 4, 5, 6, 7, 8
}
```

### func FilterNil\[T any\](in []*T) (out []*T)

_Removes nil values from slices_

Usage:

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	type exampleStruct struct{}

	a := new(exampleStruct)
	b := new(exampleStruct)
	c := new(exampleStruct)

	filtered := slices.FilterNil([]*exampleStruct{a, b, nil, c, nil})
	println(len(filtered)) // 3
}
```

### func AppendNotNil

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	type exampleStruct struct{}

	a := new(exampleStruct)
	b := new(exampleStruct)
	c := new(exampleStruct)

	filtered := slices.FilterNil([]*exampleStruct{a, b, nil, c, nil})
	println(len(filtered)) // 3
}    
```

### func Contains[S ~[]E, E comparable](s S, v E) bool

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	println(slices.Contains([]int{1,2,3,4,5}, 4)) //true
	println(slices.Contains([]string{"foo", "bar"}, "baz")) //false
}
```

### func Gro

### func Index[S ~[]E, E comparable](s S, v E) int

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	println(slices.Index([]int{1,2,3,4,5}, 4)) //3
	println(slices.Contains([]string{"foo", "bar"}, "baz")) //-1
}
```

### func Intersection[T cmp.Ordered](pS ...[]T) []T

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	println(slices.Intersection([]int{1,2,3,4,5}, []int{4,5,6,7,8})) // {4, 5}
	println(slices.Intersection([]string{"foo", "bar"}, []string{"baz"})) // {}
}
```

### func Map[S ~[]V, V any, E any](s S, extract func(v V) E) []E

```go
package main

import (
	"github.com/spandigital/slices"
)

func main() {
	println(slices.Map([]int{1,2,3,4,5}, func(v int) {
	    return v*2
	}) // 2, 4, 6, 8, 10
}