# LinearMap
[![github]](https://github.com/cospectrum/linear-map)
[![goref]](https://pkg.go.dev/github.com/cospectrum/linear-map)

[github]: https://img.shields.io/badge/github-cospectrum/linear--map-8da0cb?logo=github
[goref]: https://pkg.go.dev/badge/github.com/cospectrum/linear-map

A map implemented by searching linearly in a slice of (key, value).

Intended for use with a small number of elements.

## Install

```sh
go get github.com/cospectrum/linear-map
```
Requires Go version 1.18 or greater.

## Usage

```go
package main

import "github.com/cospectrum/linear-map/linearmap"

func main() {
	m := linearmap.New[int, string]() // empty
	m.Put(1, "x")                   // 1->x
	m.Put(2, "b")                   // 1->x, 2->b
	m.Put(1, "a")                   // 1->a, 2->b
	_, _ = m.Get(2)                 // b, true
	_, _ = m.Get(3)                 // "", false
	_ = m.Values()                  // []string{"a", "b"}
	_ = m.Keys()                    // []int{1, 2}
	m.Remove(1)                     // 2->b
	m.Clear()                       // empty
	m.Empty()                       // true
	m.Size()                        // 0
}
```

## Benchmarks

Benchmarks are performed on arm64 with Key=int, Value=struct{}.

The lower is better

Length: 10
| Method | LinearMap | HashMap |
| ---- | ------------ | ------ |
| Get | 27 ns/op | 49 ns/op |
| Put | 30 ns/op | 70 ns/op |
| Remove | 7 ns/op | 16 ns/op |

Length: 20
| Method | LinearMap | HashMap |
| ---- | ------------ | ------ |
| Get | 85 ns/op | 97 ns/op |
| Put | 92 ns/op | 140 ns/op |
| Remove | 13 ns/op | 29 ns/op |

Length: 30
| Method | LinearMap | HashMap |
| ---- | ------------ | ------ |
| Get | 174 ns/op | 147 ns/op |
| Put | 184 ns/op | 198 ns/op |
| Remove | 20 ns/op | 67 ns/op |
