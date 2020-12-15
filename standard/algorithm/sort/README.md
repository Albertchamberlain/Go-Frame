### 排序与查找

排序体现在sort包中`sort`文件中，查找体现在sort包中的`search`文件里


### 接口实现

如果要使用sort包的各个函数，就需要实现sort.Interface  即定义规则


```go
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

### Sort
sort包最核心的函数，`Sort`，用于对一个列表上的元素进行排序，Sort函数会在原有列表上进行排序，函数声明如下：

```go
func Sort(data Interface)
```
### Stable
相较于Sort函数，Stable函数也用于对一个列表进行排序，但是它额外提供保证排序算法是稳定的，也就是排序前后值相同的两个元素相对位置不发生变化，函数声明和Sort类似。

```go
func Stable(data Interface)
```
### Slice
Slice函数用于对一个Slice进行排序，这是实际使用中更为常用的一个函数，函数接收两个参数。第一个是需要排序的Slice；第二个是Slice元素比较函数，它类似于前面sort.Interface里的Less方法。函数声明如下：
```go
func Slice(slice interface{}, less func(i, j int) bool)
```
### Reverse
Reverse函数用于翻转一个列表并返回翻转后的列表，函数声明如下：

```go
func Reverse(data Interface) Interface
```
### IsSorted
IsSorted函数用于判断一个列表是否有序，函数声明如下：
```go
func IsSorted(data Interface) bool
```
### Search
Search函数可以在一个有序列表上进行二分查找操作，它接收两个参数，第一个为从第一个元素开始搜索的元素个数；第二个参数是一个函数，通过接收一个函数f作为参数，找到使得f(x)==true的元素，函数声明如下：

```go
func Search(n int, f func(int) bool) int

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // 避免h溢出
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
```