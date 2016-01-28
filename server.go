
package main

import(
    "fmt"
    "net"
)

var p = fmt.Println

const(

ip="127.0.0.1"
port=7777

)

func main(){


    listener, err := net.ListenTCP("tcp",&net.TCPAddr{net.ParseIP(ip),port,""})
    if err!=nil{
        p("listen port failed")
        return
    }
    p("waiting for connections")
    Server(listener)

}


func Server(listener *net.TCPListener){

    for {
        con, err := listener.AcceptTCP();
        if err !=nil{
            p("accept error")
            continue
        }
        p("get connection from ",con.RemoteAddr().String())
        defer con.Close();
    
        go func(){

            data := make([]byte, 1024)
            for{

                i, err := con.Read(data)
                p("got data from ",con.RemoteAddr().String(),", values:",string(data[0:i]))
                if err!=nil{
                    p("read data failed")
                    break

                }
                ret,err :=con.Write(data[0:i])
                p("write back~~~~~",ret, err)

            }


        }()






    }



}
