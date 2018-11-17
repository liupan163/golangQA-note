package main

import (
	"reflect"
)

func IsNil(any interface{}) bool {
	re := false
	if any != nil {
		v := reflect.ValueOf(any)
		if v.Kind() == reflect.Ptr {
			re = v.IsNil()
			if !re {
				for {
					v2 := v.Elem()
					if v2.Kind() != reflect.Ptr {
						break

					}
					re = v2.IsNil()
					if re {
						break
					}
					v = v2
				}
			}
		}
	}
	return re
}

// 测试并发效率
//func BenchmarkLoopsParallel(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		var test ForTest
//		ptr := &test
//		for pb.Next() {
//			ptr.Loops()
//		}
//	}
//}