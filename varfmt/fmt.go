//go:build !solution

package varfmt

import (
	"fmt"
	"strconv"
	"strings"
)

func Sprintf(format string, args ...any) string {
	nextArgIdx := 0

	argCache := make(map[int]string, len(args))

	var sb strings.Builder
	sb.Grow(len(format))

	for i := 0; i < len(format); {
		if format[i] != '{' {
			sb.WriteByte(format[i])
			i++
			continue
		}

		// find end bracket
		end := i
		for end < len(format) && format[end] != '}' {
			end++
		}
		if end == len(format) {
			panic("open argument parenthesis")
		}

		// get argument string
		argIdx := nextArgIdx
		if end-i > 1 {
			idx, err := strconv.Atoi(format[i+1 : end])
			if err != nil {
				panic("unable to parse argument index")
			}

			if idx >= len(args) {
				panic("invalid argument index")
			}

			argIdx = idx
		}

		// add simple arg caching strategy
		if _, ok := argCache[argIdx]; !ok {
			argCache[argIdx] = fmt.Sprint(args[argIdx])
		}

		sb.WriteString(argCache[argIdx])
		nextArgIdx++
		i = end + 1
	}

	return sb.String()
}

// goos: darwin
// goarch: arm64
// pkg: gitlab.com/slon/shad-go/varfmt
// cpu: Apple M1 Pro
// BenchmarkFormat
// BenchmarkFormat/small_int
// BenchmarkFormat/small_int-10            10789147                98.18 ns/op           56 B/op          4 allocs/op
// BenchmarkFormat/small_string
// BenchmarkFormat/small_string-10          7358712               162.0 ns/op           120 B/op          6 allocs/op
// BenchmarkFormat/big
// BenchmarkFormat/big-10                     30320             38943 ns/op          139578 B/op         27 allocs/op
// BenchmarkSprintf
// BenchmarkSprintf/small
// BenchmarkSprintf/small-10               31490827                38.27 ns/op            2 B/op          1 allocs/op
// BenchmarkSprintf/small_string
// BenchmarkSprintf/small_string-10        23475153                51.38 ns/op           16 B/op          1 allocs/op
// BenchmarkSprintf/big
// BenchmarkSprintf/big-10                    29937             39692 ns/op           16392 B/op          1 allocs/op
