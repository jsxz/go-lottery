package main

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)

func TestMvc(t *testing.T) {
	e := httptest.New(t, newApp())
	var wg sync.WaitGroup
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前参与抽奖用户数:0")
	for i:=0 ; i<100;i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e.POST("/import").WithFormField("users",fmt.Sprintf("test_u%d",i)).Expect().
				Status(httptest.StatusOK)
		}(i)
	}
	wg.Wait()
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal("当前参与抽奖用户数:100")
}
