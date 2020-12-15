###  环形链表(Ring)


- ring的数据结构

ring是一种循环链表，没有头尾，从任意一个节点出发就可以遍历整个链表  Value表示当前节点的值 
```go
type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}
```

### 新建 （New）

```go
func New(n int) *Ring{}
```
新建一个Ring 接收一个int值作为ring的长度


### 遍历
遍历也是Ring的常用操作

被遍历链表的元素不能为空，因此如果为空需要先调用init方法进行`自环`
- init()


- Next()返回当前节点的后一个节点
- Prev()返回当前节点的前一个节点


```go
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}
```


```go
// Next returns the next ring element. r must not be empty.
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev returns the previous ring element. r must not be empty.
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}
```

通过这两个方法均可以对ring进行不同时针的遍历，首先保存当前节点，然后依次访问下一个节点，直到回到起始节点
```go
p := ring.Next()
//  do something with first element
for p != ring {
    // do something with current element
    
    p = p.Next()
}
```

```go
p := ring.Prev()
//  do something with first element
for p != ring {
    // do something with current element
  
    p = p.Prev()
}
```

### Link与Unlink

Link将两个ring连接到一起，而Unlink将一个ring拆分为两个，移除n个元素并组成一个新的ring
func (r *Ring) Link(s *Ring) *Ring
func (r *Ring) Unlink(n int) *Ring

### Len()返回ring内包含的值个数

```go
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
```


### Do()
Do是ring中一个很灵活的方法，可以传入一个函数作为方法的参数
也就这Do方法负责做个操作，这个操作由传入的函数决定来干什么，通过传递不同的函数，可以在同一个ring上实现多种不同的操作
总体上来讲进行了一定程度上的抽象

```go
func (r *Ring) Do(f func(interface{})) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}
```












### Ring与List的区别在哪儿？

本部分来自go语言中文网

container/ring包中的Ring类型实现的是一个循环链表，也就是我们俗称的环。其实List在内部就是一个循环链表。它的根元素永远不会持有任何实际的元素值，而该元素的存在，就是为了连接这个循环链表的首尾两端。

所以也可以说，List的零值是一个只包含了根元素，但不包含任何实际元素值的空链表。两者本质都是循环链表，最主要的不同有下面几种。

- Ring类型的数据结构仅由它自身即可代表，而List类型则需要由它以及Element类型联合表示。这是表示方式上的不同，也是结构复杂度上的不同。
- 一个Ring类型的值严格来讲，只代表了其所属的循环链表中的一个元素，而一个List类型的值则代表了一个完整的链表。这是表示维度上的不同。
- 在创建并初始化一个Ring值的时候，我们可以指定它包含的元素的数量，但是对于一个List值来说，却不能这样做（也没有必要这样做）。循环链表一旦被创建，其长度是不可变的。这是两个代码包中的New函数在功能上的不同，也是两个类型在初始化值方面的第一个不同。
- 仅通过var r ring.Ring语句声明的r将会是一个长度为1的循环链表，而List类型的零值则是一个长度为0的链表。别忘了List中的根元素不会持有实际元素值，因此计算长度时不会包含它。这是两个类型在初始化值方面的第二个不同。
- Ring值的Len方法的算法复杂度是 O(N) 的，而List值的Len方法的算法复杂度则是 O(1)的。这是两者在性能方面最显而易见的差别。