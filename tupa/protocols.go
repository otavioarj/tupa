package tupa

import (
        "errors"
        "net"
)

type Fu func(string,string, []string,[] string,net.Conn)

type FuMap map[string]Fu

var FuProts= FuMap{}

var Tout int

func ProtoAdd(FuMap){
}

var tlsVerify bool


func initProts(){
        FuProts["ftp"]= Ftpcon
//        FuProts["ftps"]= Ftpscon // commented as it's used as example in main.go
}
func chkFu (f string) (error) {
        if (FuProts[f] == nil){ return errors.New("[-] Unknown protocol name: "+f)}
        return nil                    
}
