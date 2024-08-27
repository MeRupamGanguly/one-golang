package main

import (
	"fmt"
	"sync"
)

func palindrome(s string, ch chan<- map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	m := make(map[string]bool)
	for l < r {
		if arr[l] != arr[r] {
			m[s] = false
			ch <- m
			return
		}
		l++
		r--
	}
	m[s] = true
	ch <- m
}

func main() {
	arr := []string{"lola", "boob", "fidrat", "Ninja"}
	ch := make(chan map[string]bool, len(arr))
	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go palindrome(arr[i], ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	m := make(map[string]bool)
	for i := range ch {
		for k, v := range i {
			m[k] = v
		}
	}
	fmt.Println(m)
}
