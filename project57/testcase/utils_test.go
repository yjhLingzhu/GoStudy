package testcase

import "testing"

func TestAdd(t *testing.T) {
	// 调用
	res := add(1, 2)
	if res != 3 {
		t.Fatalf("执行错误！")
	}

	// 正确运行, 输出日志
	t.Logf("执行正确")
}
