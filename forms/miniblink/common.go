package miniblink

import (
	mb "qq2564874169/goMiniblink"
	"reflect"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var _jsFns = make(map[uintptr]*mb.JsFuncBinding)
var _wkeToCore = make(map[wkeHandle]ICore)

func createWke(core ICore) wkeHandle {
	wke := wkeCreateWebView()
	_wkeToCore[wke] = core
	return wke
}

func toJsValue(core ICore, es jsExecState, value interface{}) jsValue {
	if value == nil {
		return jsUndefined()
	}
	switch value.(type) {
	case int:
		return jsInt(int32(value.(int)))
	case int8:
		return jsInt(int32(value.(int8)))
	case int16:
		return jsInt(int32(value.(int16)))
	case int32:
		return jsInt(value.(int32))
	case int64:
		return jsDouble(float64(value.(int64)))
	case uint:
		return jsInt(int32(value.(uint)))
	case uint8:
		return jsInt(int32(value.(uint8)))
	case uint16:
		return jsInt(int32(value.(uint16)))
	case uint32:
		return jsInt(int32(value.(uint32)))
	case uint64:
		return jsDouble(float64(value.(uint64)))
	case float32:
		return jsFloat(value.(float32))
	case float64:
		return jsDouble(value.(float64))
	case bool:
		return jsBoolean(value.(bool))
	case string:
		return jsString(es, value.(string))
	case time.Time:
		return jsDouble(float64(value.(time.Time).Unix()))
	default:
		break
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		length := rv.Len()
		arr := jsEmptyArray(es)
		jsSetLength(es, arr, uint32(length))
		for i := 0; i < length; i++ {
			v := toJsValue(core, es, rv.Index(i).Interface())
			jsSetAt(es, arr, uint32(i), v)
		}
		return arr
	case reflect.Map:
		obj := jsEmptyObject(es)
		kv := rv.MapRange()
		for kv.Next() && kv.Key().Kind() == reflect.String {
			k := kv.Key().Interface().(string)
			v := toJsValue(core, es, kv.Value().Interface())
			jsSet(es, obj, k, v)
		}
		return obj
	case reflect.Struct:
		obj := jsEmptyObject(es)
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Field(i).Type().Name()
			v := toJsValue(core, es, rv.Field(i).Interface())
			jsSet(es, obj, f, v)
		}
		return obj
	case reflect.Func:
		//todo 没有处理保持引用
		jsFn := jsData{}
		name := "function"
		for i := 0; i < len(name); i++ {
			jsFn.name[i] = name[i]
		}
		var call = func(fnes jsExecState, obj, args jsValue, count uint32) jsValue {
			arr := make([]reflect.Value, count)
			for i := uint32(0); i < count; i++ {
				jv := jsArg(fnes, i)
				arr[i] = reflect.ValueOf(toGoValue(core, fnes, jv))
			}
			rs := rv.Call(arr)
			if len(rs) > 0 {
				return toJsValue(core, fnes, rs[0].Interface())
			}
			return 0
		}
		jsFn.callAsFunction = syscall.NewCallbackCDecl(call)
		return jsFunction(es, &jsFn)
	}
	panic("不支持的go类型：" + rv.Kind().String() + "(" + rv.Type().String() + ")")
}

func toGoValue(core ICore, es jsExecState, value jsValue) interface{} {
	switch jsTypeOf(value) {
	case jsType_NULL, jsType_UNDEFINED:
		return nil
	case jsType_NUMBER:
		return jsToDouble(es, value)
	case jsType_BOOLEAN:
		return jsToBoolean(es, value)
	case jsType_STRING:
		return jsToTempString(es, value)
	case jsType_ARRAY:
		length := jsGetLength(es, value)
		ps := make([]interface{}, length)
		for i := 0; i < length; i++ {
			v := jsGetAt(es, value, uint32(i))
			ps[i] = toGoValue(core, es, v)
		}
		return ps
	case jsType_OBJECT:
		ps := make(map[string]interface{})
		keys := jsGetKeys(es, value)
		for _, k := range keys {
			v := jsGet(es, value, k)
			ps[k] = toGoValue(core, es, v)
		}
		return ps
	case jsType_FUNCTION:
		name := "func" + strconv.FormatInt(mb.NewId(), 10)
		jsSetGlobal(es, name, value)
		return mb.JsFunc(func(param ...interface{}) interface{} {
			jses := wkeGlobalExec(wkeHandle(core.GetHandle()))
			ps := make([]jsValue, len(param))
			for i, v := range param {
				ps[i] = toJsValue(core, jses, v)
			}
			fn := jsGetGlobal(jses, name)
			rs := jsCall(jses, fn, jsUndefined(), ps)
			jsSetGlobal(jses, name, jsUndefined())
			return toGoValue(core, jses, rs)
		})
	default:
		panic("不支持的js类型：" + strconv.Itoa(int(value)))
	}
}

func BindJsFunc(binding *mb.JsFuncBinding) {
	id := uintptr(unsafe.Pointer(binding))
	_jsFns[id] = binding
	wkeJsBindFunction(binding.Name, jsFuncCallback, id, 0)
}

func jsFuncCallback(es jsExecState, state uintptr) jsValue {
	wke := jsGetWebView(es)
	core := _wkeToCore[wke]
	count := jsArgCount(es)
	ps := make([]interface{}, count)
	for i := 0; i < int(count); i++ {
		value := jsArg(es, uint32(i))
		ps[i] = toGoValue(core, es, value)
	}
	if fn, ok := _jsFns[state]; ok {
		rs := fn.OnExecute(ps)
		return toJsValue(core, es, rs)
	}
	return 0
}
