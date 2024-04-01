package main

import (
	"fmt"
)

type List[T comparable] struct {
	l []T
}

func NewList[T comparable]() List[T] {
	return List[T]{}
}

func (l *List[T]) Add(t T) {
	l.l = append(l.l, t)
}

func (l *List[T]) AddAll(list List[T]) {
	l.l = append(l.l, list.l...)
}

func (l *List[T]) AddMulti(t... T) {
	l.l = append(l.l, t...)
}

func (l *List[T]) Remove(index int) {
	list := *new([]T)

	for i, li := range l.l {
		if i != index {
			list = append(list, li)
		} 
	}

	l.l = list
}

func (l *List[T]) RemoveAll() {
	list := new([]T)
	l.l = *list
}

func (l *List[T]) ToString() string {
	return fmt.Sprintf("%v", l.l)
}

func (l *List[T]) Map(f func (t T) T) *List[T] {
	list := *new([]T)
	
	for _, li := range l.l {
		ls := f(li)
		list = append(list, ls)
	}

	li := List[T]{}
	li.l = list

	return &li
}

func (l *List[T]) Filter(f func (t T) bool) *List[T] {
	list := *new([]T)
	
	for _, li := range l.l {
		if f(li) {
			list = append(list, li)
		}
	}

	li := List[T]{}
	li.l = list

	return &li

}

func (l *List[T]) Contains(t T) bool {
	for _, li := range l.l {
		if li == t {
			return true
		}
	}

	return false
}

func Fold[T comparable, R comparable](i T, l List[R], f func(t T, u R) T ) T {
	start := i

	for _, li := range l.l {
		end := f(start, li)
		if end != start {
			start = end
		}		
	}

	return start
}
