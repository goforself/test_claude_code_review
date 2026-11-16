package main

import (
	"regexp"
	"sync"
)

// 全局正则每次调用都重新编译,应该定义为包级变量
func matchEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// 并发写 map 未加锁,可能导致 panic
var cache = make(map[string]int)

func setCache(key string, value int) {
	cache[key] = value
}

// 类型断言未判断是否成功
func toInt(v interface{}) int {
	return v.(int)
}

// 声明了锁但从未使用
var mu sync.Mutex

func getCache(key string) int {
	return cache[key]
}
