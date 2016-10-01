package neram
import("time"
       "helper"
       "strings"
       "errors"
       "strconv"
       //"fmt"
    )
type Neram struct {
  Time time.Time
  Hours int
  Minutes int
  Seconds int
  Day int
  Month string
  Year int
}
var formats=[5]string{"HHHH","MMMM","YYYY","YY","MM"}
var Month=[12]string{"January","February","March","April","May","June","July","August","September","October","November","December"}
var Manipulator=[4]string{"days","hours","s","s"}
func New(Time time.Time)(*Neram)  {
  Year,Month,Day:=Time.Date()
  Hours,Minutes,Seconds:=Time.Clock()
  return &Neram{Time,Hours,Minutes,Seconds,Day,Month.String(),Year}
}
func Now()(*Neram)  {
return New(time.Now())
}

func (neram *Neram)Format(format string)(string,error) {
   splitedformat:=strings.Split(format," ")
   for i := 0;i <len(splitedformat);i++ {
        if helper.Indexof(formats,splitedformat[i]) < 0{
          return "",errors.New("invalid format")
        }
   }
   var result string =""
   for i:=0;i<len(splitedformat);i++{
     switch splitedformat[i] {
     case "HHHH":
               result=result+strconv.Itoa(neram.Hours)+" "
     case "MMMM":
               result=result+neram.Month+" "
     case "YYYY":
               result=result+strconv.Itoa(neram.Year)+" "
     case "YY":
               year:=strconv.Itoa(neram.Year)
               Splityear:=strings.Split(year,"")
               result=result+Splityear[2]+Splityear[3]+" "
     case "MM":
               month:=helper.Indexof(Month,neram.Month)+1
               result=result+strconv.Itoa(month)+" "
     }
   }
return result,nil
}

func (neram *Neram)GetDay()int  {
return neram.Day
}

func (neram *Neram)GetHours()int  {
return neram.Hours
}

func (neram *Neram)GetSeconds()int  {
return neram.Seconds
}

func (neram *Neram)GetMinutes()int  {
return neram.Minutes
}

func (neram *Neram)GetMonth()string  {
return neram.Month
}

func (neram *Neram)GetYear()int  {
return neram.Year
}

func (neram *Neram)Add(n int,manipulator string)(*Neram,error)  {
  var NewObject Neram
  switch manipulator {
  case "Days":
             NewObject.Time=neram.Time.AddDate(0,0,n)
  case "Months":
             NewObject.Time=neram.Time.AddDate(0,n,0)
  case "Years":
             NewObject.Time=neram.Time.AddDate(n,0,0)
  case "Seconds":
             NewObject.Time=neram.Time.Add(time.Second*time.Duration(n))
  case "Minutes":
             NewObject.Time=neram.Time.Add(time.Minute*time.Duration(n))
  case "Hours":
             NewObject.Time=neram.Time.Add(time.Hour*time.Duration(n))
  default:
             return nil,errors.New("undefined manipulator format")
  }
  NewObject.Hours,NewObject.Minutes,NewObject.Seconds=NewObject.Time.Clock()
  NewObject.Year,_,NewObject.Day=NewObject.Time.Date()
  _,Month,_:=neram.Time.Date()
  NewObject.Month=string(Month)
  return &NewObject,nil
}

func (neram *Neram)Sub(n int,manipulator string)(*Neram,error)  {
  return neram.Add(-n,manipulator)
}

func (neram *Neram)StartOf(manipulator string)(*Neram,error)  {
  NewObject:=new(Neram)
  switch manipulator {
  case "Year":
              NewObject=New(time.Date(neram.Year,time.Month(1),1,0,0,0,0,neram.Time.Location()))
  }
  return NewObject,nil
}
