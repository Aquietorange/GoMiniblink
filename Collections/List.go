package Collections

type List struct {
	data  []interface{}
	count int
}

func (_this *List) Init() *List {
	return _this
}

func (_this List) Add(item interface{}) int {
	_this.data = append(_this.data, item)
	_this.count++
	return _this.count - 1
}

func (_this List) RemoveAt(index int) interface{} {
	if index < 0 || index >= _this.count {
		return nil
	}
	v := _this.data[index]
	tmp := _this.data[:index]
	_this.data = append(tmp, _this.data[index+1:])
	return v
}

func (_this List) Remove(item interface{}) bool {
	if item == nil {
		return false
	}
	for i, v := range _this.data {
		if v == item {
			_this.RemoveAt(i)
			return true
		}
	}
	return false
}

func (_this List) GetItem(index int) interface{} {
	return _this.data[index]
}

func (_this List) GetAll() []interface{} {
	return _this.data
}

func (_this List) Count() int {
	return _this.count
}
