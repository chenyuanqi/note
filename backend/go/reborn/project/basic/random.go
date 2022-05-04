package main

import (
	crand "crypto/rand"

	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func main() {
	var src cryptoSource
	rnd := rand.New(src)
	fmt.Println(rnd.Intn(1000)) // a truly random number 0 to 999

	// basics
	rand.Seed(time.Now().UnixNano())
	n := rand.Int63() // for example 9010053057479643344
	fmt.Println(n)
	x := rand.Float64() // for example 0.8058824439070219
	fmt.Println(x)

	// use methods to generate random numbers
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	n1 := generator.Int63()
	fmt.Println(n1)
	x1 := generator.Float64()
	fmt.Println(x1)

	// random element from slice
	chars := []rune("AB⌘")
	c := chars[rand.Intn(len(chars))] // for example '⌘'
	fmt.Println(c)

	// generate password
	fmt.Println(generatePassword(32))

	// generate uuid
	fmt.Println(generateUuid())

	// shuffle a slice
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	fmt.Println(a) // [3 2 6 1 8 5 4 7]
}

// @title  生成指定位数密码
// @param  密码位数 int
// @return 密码 string
func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

// @title  生成uuid
// @return uuid string
func generateUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
