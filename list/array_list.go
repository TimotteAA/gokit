package list

import (
	"gokit/internal"
)

// 基于切片的简单arrayList封装
type ArrayList[T any] struct {
	vals []T
}	

// 创建一个arrayList，容器为cap，长度为0
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, cap),
	}
}

// 从已有的切片创建arrayList
func NewArrayListOf[T any](s []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: s,
	}
}

func (a *ArrayList[T]) Add(idx int, val T) error {
	if idx < 0 || idx >= a.Len() {
		return internal.NewErrorIndexOutOfRange(idx, a.Len())
	}
	// 加到末尾
	a.vals = append(a.vals, val)
	// 从最后一个元素开始，把前面的往后移
	for i := len(a.vals) - 1; i > idx; i-- {
		a.vals[i] = a.vals[i-1]
	}
	return nil
}

func (a *ArrayList[T]) Append(val ...T) error {
	a.vals = append(a.vals, val...)
	return nil
}

func (a *ArrayList[T]) Set(idx int, val T) error {
	if idx < 0 || idx >= a.Len() {
		return internal.NewErrorIndexOutOfRange(idx, a.Len())
	}
	a.vals[idx] = val
	return nil
}

// 使用具名返回值，默认返回零值
func (a *ArrayList[T]) Get(idx int) (t T, err error) {
	if idx < 0 || idx >= a.Len() {
		err = internal.NewErrorIndexOutOfRange(idx, a.Len())
		return
	}
	t = a.vals[idx]
	return 
}

func (a *ArrayList[T]) Delete(idx int) (t T,  err error) {
	if idx < 0 || idx >= a.Len() {
		err = internal.NewErrorIndexOutOfRange(idx, a.Len())
		return
	}
	t = a.vals[idx]
	for i := idx; i < a.Len() - 1; i++ {
		a.vals[i] = a.vals[i + 1]
	}
	// 删除最后一个元素
	a.vals = a.vals[:a.Len() - 1]

	return t, err
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}