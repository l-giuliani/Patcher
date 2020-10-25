package main
import (
	"fmt"
    "os"
    dns "patcher/dns"
    gui "patcher/gui"
) 

func  main(){
    fmt.Println("Awesome Patcher")
    if(len(os.Args) < 2){
        gui.StartGui()
        return
    }
    service, exist := dns.ServiceMap[os.Args[1]]
    if (exist) {
        /*if (len(os.Args) < 5){
            fmt.Println("Number params error")
            return
        }*/
        res := service(os.Args[2:])
        if (res) {
            fmt.Println("Success")
            //os.Exit(0)
        } else {
            //os.Exit(-1)
        }
    } else if (os.Args[1] == "update-service"){
        fmt.Println("Not implemented yet")
    } else {
        fmt.Println("Option error")
    }
}
