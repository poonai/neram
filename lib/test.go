package lib
import("fmt")
type Basic interface{
  Demo() int
}

func Summa(ba Basic)  {
  fmt.Print(ba.Demo())
}
