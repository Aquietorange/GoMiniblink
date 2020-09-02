package GoMiniblink

import (
	"reflect"
	"strconv"
)

type EventDispatcher struct {
	_seq    uint64
	_key    string
	_fnkeys []string
	_fnMap  map[string]interface{}
}

func (_this *EventDispatcher) Init(key string) *EventDispatcher {
	_this._key = key
	_this._fnMap = make(map[string]interface{})
	return _this
}

func (_this *EventDispatcher) IsEmtpy() bool {
	return len(_this._fnkeys) == 0
}

func (_this *EventDispatcher) Add(name string, fn interface{}) {
	if _, ok := _this._fnMap[name]; ok == false {
		_this._fnMap[name] = fn
		_this._fnkeys = append(_this._fnkeys, name)
	}
}

func (_this *EventDispatcher) AddEx(fn interface{}) string {
	name := "fn" + strconv.FormatUint(_this._seq, 10)
	_this.Add(name, fn)
	_this._seq++
	return name
}

func (_this *EventDispatcher) Remove(name string) interface{} {
	if v, ok := _this._fnMap[name]; ok {
		delete(_this._fnMap, name)
		for i := 0; i < len(_this._fnkeys); i++ {
			if _this._fnkeys[i] == name {
				_this._fnkeys = append(_this._fnkeys[:i], _this._fnkeys[i+1:]...)
				break
			}
		}
		return v
	}
	return nil
}

func (_this *EventDispatcher) Fire(key string, sender interface{}, param ...interface{}) {
	if key != _this._key {
		panic("key不正确")
	}
	if len(_this._fnkeys) == 0 {
		return
	}
	args := make([]reflect.Value, len(param)+1)
	args[0] = reflect.ValueOf(sender)
	for i, n := range param {
		args[i+1] = reflect.ValueOf(n)
	}
	for _, key := range _this._fnkeys {
		fn := reflect.ValueOf(_this._fnMap[key])
		fn.Call(args[:fn.Type().NumIn()])
	}
}

func (_this *EventDispatcher) Clear() {
	_this._seq = 0
	_this._fnkeys = make([]string, 0)
	_this._fnMap = make(map[string]interface{})
}
