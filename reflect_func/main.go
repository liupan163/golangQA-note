package reflect_func

import "reflect"

// value := reflect.New(typ Type )
func main() {
	v := reflect.ValueOf(nil)
	reflect.TypeOf(nil)

	v.Kind() // 值 的 Type，golang里面的基本类型
	v.Type() // Type

	v.Elem() // 是顺序遍历的
	v2 := v.Elem()
	v = v2

	v.IsNil() // 注意 Ptr跟 Infterface类型的处理

	//v.Method(0).Type().AssignableTo(v.Type())
	// value 	 .Type  .AssignableTo( Type )

	v.Call([]reflect.Value{reflect.ValueOf(v), v.Method(0)})

	//token || ast
}
/*
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr || value.Kind() != reflect.Struct{ ... }
	value.NumField()
	value.IsValid()

	valueField := value.Field(i)
	value.Type().AssignableTo(valueField.Type())

	rType := value.Type()

	value.Field(1)    				// Value
	typeField := rType.Field(i)	    // StructField    特别用法

	tagName, ok := typeField.Tag.Lookup(dot.TagDot)
	if typeField.Type.Kind() != reflect.Ptr && typeField.Type.Kind() != reflect.Interface { ... }
	typeField.Name
*/
