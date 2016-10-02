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
 fmt.Print(demo.Format("HHHH [escaped] YY")) // prints  escaped 16
 }
 ```
# suported formats 
 + HHHH (HOURS)
 + YYYY (YEARS)
 + YY   (YEARS IN TWO DIGIT)
 + MMMM (MONTH IN STRING)
 + MM   (MONTH IN DIGIT)
 + [ESCAPED] (Replace the ESCAPED with whatever string you want )

# ADD
   ```go
  package main
 import "github.com/sch00lb0y/neram"
 import "fmt"
 func main(){
 demo:=neram.Now().Add(1,'Days') //add 1 day to the current time
 fmt.Print(demo.Format("YYYY"))  //prints the year in four digit
 fmt.Print(demo.Format("HHHH [escaped] YY")) // prints  escaped 16
 }
 ```
# Supported Manipulator String
 + Days
 + Months
 + Years
 + Hours
 + Minutes
 + Seconds
 
# SUB
   ```go
  package main
 import "github.com/sch00lb0y/neram"
 import "fmt"
 func main(){
 demo:=neram.Now().Sub(1,'Days') //add 1 day to the current time
 fmt.Print(demo.Format("YYYY"))  //prints the year in four digit
 fmt.Print(demo.Format("HHHH [escaped] YY")) // prints  escaped 16
 }
 ```
# Supported Manipulator String
 + Days
 + Months
 + Years
 + Hours
 + Minutes
 + Seconds
 
# StartOf
   ```go
  package main
 import "github.com/sch00lb0y/neram"
 import "fmt"
 func main(){
 demo:=neram.Now().StartOf("Year") //reset the time to the begining of the year
 fmt.Print(demo.Format("YYYY"))  //prints the year in four digit
 fmt.Print(demo.Format("HHHH [escaped] YY")) // prints  escaped 16
 }
 ```
# Supported Manipulator String
+ Day
+ Month
+ Year
+ Hour
+ Minute
