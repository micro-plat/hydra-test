package main

import (
	"github.com/micro-plat/hydra"
)

func init() {
	hydra.OnReady(func() {

		hydra.Conf.API("8070").Render(`
        request := import("request")
response := import("response")
text := import("text")
types :=import("types")

rc_xml:="<response><code>{@status}</code><msg>{@content}</msg></response>"
rc_json:="{\"msg\":\"{@content}\"}"

getContent := func(){  

    input:={status:response.getStatus(),content:response.getRaw()}

    if text.has_prefix(request.getPath(),"/xml"){
        return [200,types.translate(rc_xml,input)]
    }
     if text.has_prefix(request.getPath(),"/json"){
        return [200,types.translate(rc_json,input)]
    }
    if text.has_prefix(request.getPath(),"/plain"){
        return [200,input.content,"text/plain"]
    }
}

render := getContent()
		`)

	})
}
