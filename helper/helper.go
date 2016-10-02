package helper
import ("reflect"
        )
func Indexof(object interface{},val interface{})int  {
  switch reflect.TypeOf(object).Kind() {
  case reflect.Array:
                    object:=reflect.ValueOf(object)
                    for i := 0;i < object.Len();i++ {
                         if object.Index(i).String()==reflect.ValueOf(val).String() {
                            return i
                         }
                    }
                    return -1
  }
  //fmt.Print(reflect.TypeOf(object).Kind())
  return -1
}
