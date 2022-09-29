package tupa

import (
        
        "fmt"
        "net"     
        "strings" 
        "crypto/tls"     
)

// Example of a stacked protocol (i.e, FTP/TLS), where it relies on a TLS tunnel
func Ftpscon(hostp string,user string, pass []string, con net.Conn) {
        var err error
               
        if (con ==nil){
                conf := &tls.Config{InsecureSkipVerify: tlsVerify,}
                con, err = tls.Dial("tcp", hostp,conf)
                checkErr(err)                
        }
// Con will be closed in the inner call. Follow this approach for every protocol that does have clear and TLS tunnels
        Ftpcon(hostp,user, pass, con)        
}

   

// Example of a protocol
func Ftpcon(hostp string,user string, pass []string, con net.Conn) {
        
        var err error
        if (con ==nil){
                con, err = net.Dial("tcp", hostp)
                checkErr(err)                
        }
        defer con.Close()
        msg:= []byte("USER "+ user +"\n")
        reply := make([]byte, 256)
        _, err = con.Read(reply)
        checkErr(err)
        if (strings.Contains(string(reply),"FTP")){
                for _,p := range pass {
                        if (wg.pok>0){break}                        
                        _, err = con.Write(msg)
                        checkErr(err)                                   
                        _, err = con.Read(reply)
                        checkErr(err)
                        if (!strings.Contains(string(reply),"331")) {
                                fmt.Println("[-] Anom: " +string(reply)) 
                        }                          
                        _, err = con.Write([]byte("PASS "+ p +"\n"))
                        checkErr(err)      
                        _, err = con.Read(reply)
                        if (!strings.Contains(string(reply),"530")) {
                           fmt.Println("[+] Pass: "+p)
                            wgGot()
                        }
                }
        } 
        //fmt.Println("S: "+string(reply))  
        wgDone()   
        
}


