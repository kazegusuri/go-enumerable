# GO implementation of Ruby Enumerable

**Under Development**

## Example

```go

// map
var out1 []int
_ = enumerable.Map3([]int{3, 4, 5}, &out1, func(a int) int {
	return a + 1
})
fmt.Printf("out1: %#v\n", out1) // => out1: []int{4, 5, 6}

// map
out1 = Map([]int{3, 4, 5}, func(a int) int {
	return a + 1
}).([]int)
fmt.Printf("out1: %#v\n", out1) // => out1: []int{4, 5, 6}

// reduce
var out2 int
_ = enumerable.Reduce([]int{3, 4, 5}, &out2, func(a, b int) int {
	return a + b
})
fmt.Printf("out2: %#v\n", out2) // => out2: 12

// filter
var out3 []int
_ = enumerable.Filter([]int{2, 4, 3, 5}, &out3, func(a int) bool {
	return a > 3
})
fmt.Printf("out3: %#v\n", out3) // => out3: []int{4, 5}

// min
var out4 int
_ = enumerable.Min([]int{5, 4, 3, 7}, &out4, func(a, b int) int {
	if a < b {
		return -1
	}
	return 1
})
fmt.Printf("out4: %#v\n", out4) // => out4: 3

// max
var out5 int
_ = enumerable.Max([]int{5, 4, 3, 7}, &out5, func(a, b int) int {
	if a < b {
		return -1
	}
	return 1
})
fmt.Printf("out5: %d\n", out5) // => out5: 7

// find
var out6 int
_ = enumerable.Find([]int{5, 4, 3, 7}, &out6, func(a int) bool {
	return a == 3
})
fmt.Printf("out6: %#v\n", out6) // => out6: 3


// any
ok, _ := enumerable.Any([]int{5, 4, 3, 7}, func(a int) bool {
	return a == 5
})
fmt.Printf("ok: %v\n", ok) // => ok: true

// all
ok, _ := enumerable.All([]int{5, 4, 3, 7}, func(a int) bool {
	return a > 0
})
fmt.Printf("ok: %v\n", ok) // => ok: true
```

## Copyright

<table>
  <tr>
    <td>Author</td><td>Masahiro Sano <sabottenda@gmail.com></td>
  </tr>
  <tr>
    <td>Copyright</td><td>Copyright (c) 2015- Masahiro Sano</td>
  </tr>
  <tr>
    <td>License</td><td>MIT License</td>
  </tr>
</table>
