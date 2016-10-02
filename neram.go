// neram:- a light weight time manipulation library
// version:1.0.0
// authon:பாலாஜி(sch00lb0y)
// made with அன்பு love
package neram
import("time"
       "helper"
       "strings"
       "errors"
       "strconv"
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
// Month Array
var Month=[12]string{"January","February","March","April","May","June","July","August","September","October","November","December"}

/*
@params Time object
returns Neram object
*/
func New(Time time.Time)(*Neram)  {
  Year,Month,Day:=Time.Date()
  Hours,Minutes,Seconds:=Time.Clock()
  return &Neram{Time,Hours,Minutes,Seconds,Day,Month.String(),Year}
}
/*
creates Neram object for the current time
*/
func Now()(*Neram)  {
return New(time.Now())
}
/*
@params format string
supported format
HHHH : Hours
MMMM :Month
YYYYY :Year
YY: year (only last two digit)
MM: Month ( gives month in numbers)
returns string based on the format
refer helper package for to know about Indexof func
*/
func (neram *Neram)Format(format string)(string,error) {
   splitedformat:=strings.Split(format," ")
   for i := 0;i <len(splitedformat);i++ {

        if helper.Indexof(formats,splitedformat[i]) < 0{

                if splitedformat[i][0]!=91&&splitedformat[i][len(splitedformat[i])-1]!=93{
                     return "",errors.New("invalid format")
                    }
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
     default:
                result=result+splitedformat[i][1:len(splitedformat[i])-1]+" "
     }
   }
return result[0:len(result)-1],nil
}

// returns day
func (neram *Neram)GetDay()int  {
return neram.Day
}
// returns Hours
func (neram *Neram)GetHours()int  {
return neram.Hours
}
//returns Seconds
func (neram *Neram)GetSeconds()int  {
return neram.Seconds
}
// return Minutes
func (neram *Neram)GetMinutes()int  {
return neram.Minutes
}
//returns Month
func (neram *Neram)GetMonth()string  {
return neram.Month
}
// returns Year
func (neram *Neram)GetYear()int  {
return neram.Year
}
/*@params n integer that you want to add
 *@params manipulator string (Days Months Years Seconds Minutes Hours)
 * returns added Neram object
 */
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
/*
same as add func but it decrements
*/

func (neram *Neram)Sub(n int,manipulator string)(*Neram,error)  {
  return neram.Add(-n,manipulator)
}
/*
StartOf func makes the Neram object to move to start of the manipulator string
that mentioned
*/

func (neram *Neram)StartOf(manipulator string)(*Neram,error)  {
  NewObject:=new(Neram)
  switch manipulator {
  case "Year":
              NewObject=New(time.Date(neram.Year,time.Month(1),1,0,0,0,0,neram.Time.Location()))
  case "Month":
              NewObject=New(time.Date(neram.Year,time.Month(helper.Indexof(Month,neram.Month)+1),1,0,0,0,0,neram.Time.Location()))
  case "Day":
              NewObject=New(time.Date(neram.Year,time.Month(helper.Indexof(Month,neram.Month)+1),neram.Day,0,0,0,0,neram.Time.Location()))
  case "Hour":
              NewObject=New(time.Date(neram.Year,time.Month(helper.Indexof(Month,neram.Month)+1),neram.Day,neram.Hours,0,0,0,neram.Time.Location()))
  case "Minute":
              NewObject=New(time.Date(neram.Year,time.Month(helper.Indexof(Month,neram.Month)+1),neram.Day,neram.Hours,neram.Minutes,0,0,neram.Time.Location()))
  }
  return NewObject,nil
}
