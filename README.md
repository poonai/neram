# Neram
   A ligthweight date and time manipulation library written in go
# Installation
```
go get github.com/sch00lb0y/neram
```
# Usage
 ``` go
 package main
 import "github.com/sch00lb0y/neram"
 import "fmt"
 func main(){
 demo:=neram.Now()
 fmt.Print(demo.Format("YYYY"))  //prints the year in four digit
 }
 ```
 # API
 # Time Parser
 ```go
  package main
 import "github.com/sch00lb0y/neram"
 import "fmt"
 func main(){
 demo:=neram.Now()
 fmt.Print(demo.Format("YYYY"))  //prints the year in four digit
 fmt.Print(demo.Format("HHHH YY")) //
 }
 ```
