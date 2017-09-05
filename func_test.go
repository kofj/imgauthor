package main

import "testing"

var tests = map[string]int{
	"http://wx1.sinaimg.cn/mw690/9d0d09abgy1fj0wcs7aewj20ij0sn12y.jpg": 2634877355,
	"http://wx1.sinaimg.cn/mw690/006r2HqOgy1fj7dxg3zuxj30p02a1wry.jpg": 5896401674,
}

func TestGetUid(t *testing.T) {
	var uid int
	var err error
	for k, v := range tests {
		uid, err = getUid(k)
		if err != nil {
			t.Error(err)
		}
		if uid != v {
			t.Errorf("url: %s expected %d got %d \n", k, uid, v)
		}
	}
}
