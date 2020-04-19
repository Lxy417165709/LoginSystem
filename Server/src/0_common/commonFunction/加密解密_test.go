package commonFunction

import (
	"testing"
)

var key = "0123456789123456"

func TestEncodeAndDecode(t *testing.T) {
	testSlice := []string{
		``,
		`123456789`,
		`4778977`,
		` erqr5wqe5ewqt57e7we7tg7we`,
		`qwj wqewqeq",`,
		`wqewqllfqqwiiovweqwnxc `,
		`wqewqllfqqwigw 4545q iovweqwnxc `,
		`wqewqllwqr241897643#@87wigw 4545q iovweqwnxc `,
		`wqewqllwqr2418wqr2418wqr2418wqr2418wqr2418wqr2418wqr2418wqr241897643#@872wigw 4545q iovweqwnxc `,
		`wqewqllwqr2418wqr2418qw wq  wqr2418wqr2418wqr24w///q*r* q18wqr2418wqr2418wqr241897643#@872wigw 4545q iovweqwnxc `,
		`................................`,
	}

	for _, str := range testSlice {
		encodeStr, _ := Encode(str, key)
		decodeStr, _ := Decode(encodeStr, key)
		if decodeStr != str {
			t.Errorf("%v 案例测试出错", str)
		}
	}
}
