/**
 * @Author: fxl
 * @Description:
 * @File:  simple_test.go
 * @Version: 1.0.0
 * @Date: 2022/9/7 19:59
 */
package simple_factory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI("1")
	s := api.Say("tom")
	if s != "hi,tom" {
		t.Fatalf("type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI("2")
	s := api.Say("tom")
	if s != "hello,tom" {
		t.Fatalf("type1 test fail")
	}
}
