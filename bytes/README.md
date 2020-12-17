### Bytes

该包定义了一些操作 byte slice 的便利操作。因为字符串可以表示为 `[]byte`，因此，bytes 包定义的函数、方法等和 strings 包很类似

为了方便，会称呼 []byte 为 字节数组

#### 转换
将 s 中的所有字符修改为大写（小写、标题）格式返回。
```go
func ToUpper(s []byte) []byte
func ToLower(s []byte) []byte
func ToTitle(s []byte) []byte
```
使用指定的映射表将 s 中的所有字符修改为大写（`小写`、`标题`）格式返回。

```go
func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte
func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte
func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte
```

将 s 中的所有单词的首字符修改为 Title 格式返回。
BUG: 不能很好的处理以 Unicode 标点符号分隔的单词。

```go
func Title(s []byte) []byte
```

### 比较
比较两个 []byte，nil 参数相当于空 []byte。
`a < b 返回 -1
a == b 返回 0
a > b 返回 1`

```go
func Compare(a, b []byte) int
```

判断 a、b 是否相等，nil 参数相当于空 []byte。
```go
func Equal(a, b []byte) bool
```

判断 s、t 是否相似，忽略`大写`、`小写`、`标题`三种格式的区别。
参考 unicode.SimpleFold 函数。
```go
func EqualFold(s, t []byte) bool
```


#### 修剪
去掉 s 两边（左边、右边）包含在 cutset 中的字符（返回 s 的切片）
```go
func Trim(s []byte, cutset string) []byte
func TrimLeft(s []byte, cutset string) []byte
func TrimRight(s []byte, cutset string) []byte
```
去掉 s 两边（左边、右边）符合 f 要求的字符（返回 s 的切片）
```go
func TrimFunc(s []byte, f func(r rune) bool) []byte
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
```

去掉 s 两边的空白（unicode.IsSpace）（返回 s 的切片）
```go
func TrimSpace(s []byte) []byte
```

去掉 s 的前缀 prefix（后缀 suffix）（返回 s 的切片）
```go
func TrimPrefix(s, prefix []byte) []byte
func TrimSuffix(s, suffix []byte) []byte
```

### 合久必分，分久必合
Split 以 sep 为分隔符将 s 切分成多个子串，结果不包含分隔符。
如果 sep 为空，则将 s 切分成 Unicode 字符列表。
SplitN 可以指定切分次数 n，超出 n 的部分将不进行切分。
```go
func Split(s, sep []byte) [][]byte
func SplitN(s, sep []byte, n int) [][]byte
```

功能同 Split，只不过结果包含分隔符（在各个子串尾部）。
```go
func SplitAfter(s, sep []byte) [][]byte
func SplitAfterN(s, sep []byte, n int) [][]byte
```

以连续空白为分隔符将 s 切分成多个子串，结果不包含分隔符。
```go
func Fields(s []byte) [][]byte
```

以符合 f 的字符为分隔符将 s 切分成多个子串，结果不包含分隔符。
```go
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
```

以 sep 为连接符，将子串列表 s 连接成一个字节串。
```go
func Join(s [][]byte, sep []byte) []byte
```

将子串 b 重复 count 次后返回。
```go
func Repeat(b []byte, count int) []byte
```

示例
```go
func main() {
    b := []byte("  Hello   World !  ")
    fmt.Printf("%q\n", bytes.Split(b, []byte{' '}))
    // ["" "" "Hello" "" "" "World" "!" "" ""]
    fmt.Printf("%q\n", bytes.Fields(b))
    // ["Hello" "World" "!"]
    f := func(r rune) bool {
        return bytes.ContainsRune([]byte(" !"), r)
    }
    fmt.Printf("%q\n", bytes.FieldsFunc(b, f))
    // ["Hello" "World"]
}
```
### 子串

判断 s 是否有前缀 prefix（后缀 suffix）
```go
func HasPrefix(s, prefix []byte) bool
func HasSuffix(s, suffix []byte) bool
```

判断 b 中是否包含子串 subslice（字符 r）
```go
func Contains(b, subslice []byte) bool
func ContainsRune(b []byte, r rune) bool
```

判断 b 中是否包含 chars 中的任何一个字符
```go
func ContainsAny(b []byte, chars string) bool
```

查找子串 `sep（字节 c、字符 r）`在 s 中第一次出现的位置，找不到则返回 -1。
```go
func Index(s, sep []byte) int
func IndexByte(s []byte, c byte) int
func IndexRune(s []byte, r rune) int
```

查找 `chars` 中的任何一个字符在 s 中第一次出现的位置，找不到则返回 -1。
```go
func IndexAny(s []byte, chars string) int
```

查找符合 f 的字符在 s 中第一次出现的位置，找不到则返回 -1。
```go
func IndexFunc(s []byte, f func(r rune) bool) int
```

功能同上，只不过查找最后一次出现的位置。
```go
func LastIndex(s, sep []byte) int
func LastIndexByte(s []byte, c byte) int
func LastIndexAny(s []byte, chars string) int
func LastIndexFunc(s []byte, f func(r rune) bool) int
```

获取 sep 在 s 中出现的次数（sep 不能重叠）。
```go
func Count(s, sep []byte) int
```

### 替换
将 s 中前 n 个 old 替换为 new，n < 0 则替换全部。
```go
func Replace(s, old, new []byte, n int) []byte
```

将 s 中的字符替换为 mapping(r) 的返回值，
如果 mapping 返回负值，则丢弃该字符。
```go
func Map(mapping func(r rune) rune, s []byte) []byte
```

将 s 转换为 []rune 类型返回
```go
func Runes(s []byte) []rune
type Reader struct { ... }
```

将 b 包装成 bytes.Reader 对象。
```go
func NewReader(b []byte) *Reader
```

bytes.Reader 实现了如下接口：

`io.ReadSeeker
io.ReaderAt
io.WriterTo
io.ByteScanner
io.RuneScanner`
// 返回未读取部分的数据长度
```go
func (r *Reader) Len() int
```

// 返回底层数据的总长度，方便 ReadAt 使用，返回值永远不变。
```go
func (r *Reader) Size() int64
```

// 将底层数据切换为 b，同时复位所有标记（读取位置等信息）。
```go
func (r *Reader) Reset(b []byte)
```

type Buffer struct { ... }

将 buf 包装成 bytes.Buffer 对象。
```go
func NewBuffer(buf []byte) *Buffer
```


将 s 转换为 []byte 后，包装成 bytes.Buffer 对象。
```go
func NewBufferString(s string) *Buffer
```

Buffer 本身就是一个缓存（内存块），没有底层数据，缓存的容量会根据需要自动调整。大多数情况下，使用 new(Buffer) 就足以初始化一个 Buffer 了。
bytes.Buffer 实现了如下接口：
`io.ReadWriter
io.ReaderFrom
io.WriterTo
io.ByteWeriter
io.ByteScanner
io.RuneScanner`

未读取部分的数据长度
```go
func (b *Buffer) Len() int
```


缓存的容量
```go
func (b *Buffer) Cap() int
```

读取前 n 字节的数据并以切片形式返回，如果数据长度小于 n，则全部读取。
切片只在`下一次读写操作前`合法。
```go
func (b *Buffer) Next(n int) []byte
```

读取第一个 delim 及其之前的内容，返回遇到的错误（一般是 io.EOF）。
```go
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
func (b *Buffer) ReadString(delim byte) (line string, err error)
```

写入 r 的 UTF-8 编码，返回写入的字节数和 nil。
保留 err 是为了匹配 bufio.Writer 的 WriteRune 方法「接口实现?」。
```go
func (b *Buffer) WriteRune(r rune) (n int, err error)
```

写入 s，返回写入的字节数和 nil。
```go
func (b *Buffer) WriteString(s string) (n int, err error)
```

引用未读取部分的数据切片（不移动读取位置）
```go
func (b *Buffer) Bytes() []byte
```


返回未读取部分的数据字符串（不移动读取位置）
```go
func (b *Buffer) String() string
```


自动增加缓存容量，以保证有 n 字节的剩余空间。
如果 n 小于 0 或无法增加容量则会 panic。
```go
func (b *Buffer) Grow(n int)
```


将数据长度截短到 n 字节，如果 `n 小于 0 或大于 Cap` 则 panic。
```go
func (b *Buffer) Truncate(n int)
```


重设缓冲区，清空所有数据（包括初始内容）。
```go
func (b *Buffer) Reset()
```





