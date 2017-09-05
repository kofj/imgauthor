package main

import (
	"flag"
	"fmt"
	"net/url"
	"regexp"
)

var (
	partern = regexp.MustCompile("[[:alnum:]]{32}")
)

// * example url: https://ws3.sinaimg.cn/large/006tKfTcly1fj6mmuffgkj30id0bcgoz.jpg
func main() {
	flag.Parse()
	imgurl := flag.Arg(0)

	if len(imgurl) < 24 {
		fmt.Println("invalid weibo image url.", imgurl)
		return
	}

	uid, err := getUid(imgurl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("https://weibo.com/u/%d\n", uid)
}

func getUid(imgurl string) (uid int, err error) {
	var carry = 16
	URL, err := url.Parse(imgurl)
	if err != nil {
		return
	}

	hash := partern.FindString(URL.EscapedPath())[:8]

	if hash[0] == '0' && hash[1] == '0' {
		carry = 62
	}

	for _, b := range hash {
		uid = uid*carry + idx(b)
	}

	return
}

func idx(c rune) int {
	i := int(c)
	if i >= 48 && i <= 57 {
		return i - 48
	}

	if c >= 97 && c <= 122 {
		return i - 97 + 10
	}
	return i - 65 + 36
}
