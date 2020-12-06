/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/12/6 17:59
 */
package string_plus

import (
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
)

var (
	ch  = make(chan string, 1000)
	one = sync.Once{}
	src = rand.NewSource(time.Now().UnixNano())
)

func filling() {
	one.Do(func() {
		for len(ch) < cap(ch) {
			ch <- randStringBytesMaskImprSrcUnsafe(6)
		}
	})
	one = sync.Once{}
}

func New() {
	go func() {
		for {
			if len(ch) == 0 {
				filling()
			}
		}
	}()
}

func Get() string {
	return <-ch
}

func randStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
