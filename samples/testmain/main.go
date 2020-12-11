package main

import (
	"fmt"

	"github.com/nelsonken/xmltomap-go"
)

func main() {
	data := []byte(`<xml>
    <a>1</a>
    <b>hi hello world</b>
    <c>xxxx</c>
    <c>rrr</c>
    <dd>
        <ee>rrr</ee>
        <ff>rrr</ff>
    </dd>
</xml>`)
	strmap, err := xmltomap.Unmarshal(data)
	if err != nil {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(strmap)

}
