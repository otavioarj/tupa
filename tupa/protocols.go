package tupa

import (
        "errors"
        "net"
)

type Fu func(string,string, []string,net.Conn)

type FuMap map[string]Fu

var FunkProtocol = FuMap{}

func ProtoAdd(FuMap){
}

var tlsVerify bool

func chkFu (f string) (error) {
        if (FunkProtocol[f] == nil){ return errors.New("[-] Unknown protocol name: "+f)}
        return nil                    
}
