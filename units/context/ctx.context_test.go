package context

import (
	"reflect"
	"testing"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/types"
)

func TestNewCtx(t *testing.T) {

	got := mock.NewContext("获取content", mock.WithRHeaders(types.XMap{}))

	assert.NotEqual(t, nil, got, "获取ctx对象")
}

func TestCtx_Close(t *testing.T) {
	c := mock.NewContext("获取content", mock.WithRHeaders(types.XMap{}))

	c.Close()

	//对ctx.funcs和ctx.context为空不能进行判断
	if !reflect.ValueOf(c.Response()).IsNil() {
		t.Errorf("Close():c.response is not nil")
		return
	}
	if c.APPConf() != nil {
		t.Errorf("Close():c.APPConf is not nil")
		return
	}
	if !reflect.ValueOf(c.User()).IsNil() {
		t.Errorf("Close():c.user is not nil")
		return
	}
	if c.Context() != nil {
		t.Errorf("Close():c.ctx is not nil")
		return
	}
	if !reflect.ValueOf(c.Request()).IsNil() {
		t.Errorf("Close():c.request is not nil")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
		t.Errorf("context.Del(c.tid) doesn't run")
	}()
	context.Current()
}
