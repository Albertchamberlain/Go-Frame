
`Map主要提供一个快速的查找，插入，删除，具备与存储体量无关的O（1）的性能，并且支持key上面的唯一性`

### map的基本使用

#### 声明
- map[KeyType]ValueType

- var m map[string]string
这里的KeyType代表map的key类型，一定要是 comparable 的，而ValueType可以是任意的类型，甚至包括其他的内嵌的map

`map的声明的时候默认值是nil ，此时进行取值，返回的是对应类型的零值(不存在也是返回零值)`

map在go里是属于reference type，也就是作为方法的型参或者返回类型的是时候，传递也是这个reference的地址。
不是map的本体。其次，这个map在申明的时候是nil map，需要如果没有初始化，那么就是nil

- var m map[string]int = map[string]int{"amos":12,"albert":10}
- m = map[string]int{}


```go
type hmap struct {
    count        int  //元素个数
    flags        uint8   
    B            uint8 //扩容常量
    noverflow    uint16 //溢出 bucket 个数
    hash0        uint32 //hash 种子
    buckets      unsafe.Pointer //bucket 数组指针
    oldbuckets   unsafe.Pointer //扩容时旧的buckets 数组指针
    nevacuate    uintptr  //扩容搬迁进度
    extra        *mapextra //记录溢出相关
}

type bmap struct {
    tophash        [bucketCnt]uint8  
    // Followed by bucketCnt keys 
    //and then bucketan Cnt values  
    // Followed by overflow pointer.
}

```
每个map的底层结构是hmap，是由若干个结构为bmap的bucket组成的数组，每个bucket可以存放若干个元素(通常是8个)，那么每个key会根据hash算法被归到同一个bucket中，当一个bucket中的元素超过8个的时候，hmap会使用extra中的overflow来扩展存储key。

#### 读取

i:=m["jack"]
如果"jack"存在，就返回那个值，如果不存在，返回0值，也就是说，根据这个value的类型，返回缺省值，比如string，就返回“”，int 就返回0

#### 删除

i,ok := delete(m,"amos")
如果amos存在，删除成功，否则什么都没有发生,根据返回值来判断，因为读取不存在的keyy时返回0

#### 无序遍历
for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}

#### 顺序遍历
利用一个slice

```go
import "sort"
var keys []string
// 把key单独抽取出来，放在数组中
for k, _ := range m {
    keys = append(keys, k)
}
// 进行数组的排序
sort.Strings(keys)
// 遍历数组就是有序的了
for _, k := range keys {
    fmt.Println(k, m[k])
}

```
#### GC
delete是不会真正的把map释放的（逻辑删除），所以要回收map还是需要设为nil




