
package main

import(
    "fmt"
    "net"
    "time"
)
var p=fmt.Println

func main(){

    tcpaddr,err :=net.ResolveTCPAddr("tcp4","127.0.0.1:7777")
    checkError(err)
    ch := make(chan int,4)
    for i:=0;i<10;i++{
        go dot(ch, tcpaddr)
    }

    for i:=0;i<10;i++{
        p(<-ch)
    }


}

func dot(ch chan int, tcpaddr *net.TCPAddr){
    con, err := net.DialTCP("tcp",nil,tcpaddr)
    checkError(err)
    defer con.Close()
    for{
         _,err = con.Write([]byte("test,"+time.Now().String()))
        data := make([]byte, 1024)
        _,err := con.Read(data)
        fmt.Println("ppppppp")
        checkError(err)
        fmt.Println(string(data))

        break
    }
    ch<-1


}




func checkError(err error){
    if err != nil{
        fmt.Println("error happended:%s",err.Error())

    }


}
