/**
 * @Author: fxl
 * @Description:
 * @File:  simple.go
 * @Version: 1.0.0
 * @Date: 2022/9/7 19:55
 */
package simple_factory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t string) API {
	if t == "1" {
		return &hiAPI{}
	} else if t == "2" {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("hi,%s", name)
}

type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("hello,%s", name)
}
