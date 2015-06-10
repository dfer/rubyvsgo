package main

import (
	"fmt"
	"time"
	"runtime"
	
	// Библиотеки для работы функции node
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

// Функция аналогична функции node1 написанной на Ruby
func node1 (text string) int64 {
	h := sha1.New()
	h.Write([]byte(text))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	
	sha1_hash_len := len(sha1_hash)
	last := sha1_hash[sha1_hash_len-1:sha1_hash_len]
	
	value, _ := strconv.ParseInt(last, 16, 64) //int64
	return value % 8
}

// Этот вариант примерно на 30% быстрее
// Благодаря отказу от лишнего преобразования из одной системы счисления в другую
func node2 (text string) int64 {
	h := sha1.New()
	h.Write([]byte(text))
	mas := h.Sum(nil) // "hello world" -> [42 174 108 53 201 79 207 180 21 219 233 95 64 139 156 233 30 232 70 237]
	return int64(mas[len(mas)-1]) % 8 // Берем последний элемент массива. Это целое десятичное число. И считаем остаток от деления на 8
}

func main() {
	// Используем только 1 ядро процессора. Ruby выполняется на одном ядре. Не будем давать преимущество Go
	runtime.GOMAXPROCS(1)
	
	var (
		date_from int64
		date_to int64
		delta int64
	)
	
	// -------- node v1 -----------
	date_from = time.Now().UnixNano()
	
	for i := 1; i <= 1000000; i++ {
		node1("user:"+string(i))	
	}
	
	date_to = time.Now().UnixNano()
	
	delta = date_to - date_from
	
	fmt.Println("node1: ", delta)
	
	// -------- node v2 -----------
	date_from = time.Now().UnixNano()
	
	for i := 1; i <= 1000000; i++ {
		node2("user:"+string(i))	
	}
	
	date_to = time.Now().UnixNano()
	
	delta = date_to - date_from
	
	fmt.Println("node2: ", delta)
}