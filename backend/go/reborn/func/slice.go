package helper

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// 去除重复项
func Unique(elems []string) []string {
	m := make(map[string]struct{})
	n := make([]string, 0)
	for _, elem := range elems {
		if _, ok := m[elem]; !ok {
			m[elem] = struct{}{}
			n = append(n, elem)
		}
	}
	return n
}

// 转为 map
func ToMap(elems []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, elem := range elems {
		if _, ok := m[elem]; ok {
			continue
		}

		m[elem] = struct{}{}
	}

	return m
}

// 判断是否在 []string 里
func InArray(needle string, arr []string) bool {
	for _, elem := range arr {
		if elem == needle {
			return true
		}
	}
	return false
}

// 取元素交集，以 a 的顺序为准
func Intersect(a, b []string) []string {
	bm := map[string]bool{}
	for _, e := range b {
		bm[e] = true
	}

	var n []string
	for _, e := range a {
		if _, ok := bm[e]; ok {
			n = append(n, e)
		}
	}
	return n
}

// 合并元素，保证顺序
func Merge(a ...[]string) []string {
	m := make(map[string]bool)
	var n []string
	for _, itemids := range a {
		for _, itemid := range itemids {
			if _, ok := m[itemid]; !ok {
				m[itemid] = true
				n = append(n, itemid)
			}
		}
	}
	return n
}

// 返回 elems 里排除 a 的项，保证顺序
func Diff(elems []string, a ...[]string) []string {
	mids := Merge(a...)
	mm := make(map[string]bool)
	for _, mid := range mids {
		mm[mid] = true
	}

	var n []string
	for _, elem := range elems {
		if _, ok := mm[elem]; !ok {
			n = append(n, elem)
		}
	}
	return n
}

// 打乱顺序
func Shuffle(arr []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

// 过滤列表
func Filter(elems []string, fn func(a string) bool) []string {
	var past []string
	for _, e := range elems {
		if fn(e) {
			past = append(past, e)
		}
	}
	return past
}