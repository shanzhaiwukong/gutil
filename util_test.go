package gutil_test

import (
	"gutil"
	"testing"
)

func Test(t *testing.T) {
	var key = []byte("1234567800009999")
	var datas = []string{"123456", "abcdefg", "中国", `{"grpc_addr":"localhost:8888"}`}
	var miwen = make([]string, 0, len(datas))
	for _, s := range datas {
		if _s, e := gutil.AesEncrypt([]byte(s), key); e != nil {
			miwen = append(miwen, "")
			t.Error(e)
		} else {
			miwen = append(miwen, _s)
		}
	}

	for i, s := range miwen {
		ors, _ := gutil.AesDecrypt(s, key)
		if string(ors) != datas[i] {
			t.Error("not eq")
		}
	}
}
