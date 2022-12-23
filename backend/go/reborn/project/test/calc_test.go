package main

import "testing"
import "fmt"
import "os"

// 执行测试前
func setup() {
	fmt.Println("Before all tests.")
}

// 执行测试后
func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	// 调用 m.Run() 触发所有测试用例的执行
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

type calcCase struct {
	Name           string
	A, B, Expected int
}

// 帮助函数
func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	t.Run(c.Name, func(t *testing.T) {
		if ans := Mul(c.A, c.B); ans != c.Expected {
			t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
		}
	})

}

func TestMul(t *testing.T) {
	// t.Run("pos", func(t *testing.T) {
	// 	if Mul(2, 3) != 6 {
	// 		// 遇错会停止
	// 		t.Fatal("fail")
	// 	}
	// })

	// t.Run("neg", func(t *testing.T) {
	// 	if Mul(-2, 3) != -6 {
	// 		t.Fatal("fail")
	// 	}
	// })
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"nge", -2, 3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		createMulTestCase(t, &calcCase{c.Name, c.A, c.B, c.Expected})
		// t.Run(c.Name, func(t *testing.T) {
		// if ans := Mul(c.A, c.B); ans != c.Expected {
		// 	t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
		// }
		// })
	}
}
