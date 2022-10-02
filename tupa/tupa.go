package tupa

import (
        "bufio"
        "os"
        "fmt"     
        "flag"
)

func Start() {   
        hostp := flag.String("hp", "127.0.0.1:8080","HostPort String as host:port.")
        buff  := flag.Int("b",4096, "Input buffer size,i.e., file/memory chunks, ideally power of two.")
        thrs  := flag.Int("n",16, "Thread number, prefer n as power of two.") 
        tout  := flag.Int("to",3, "TCP timeout") 
        user  := flag.String("u","admin","User/account for the attack.")
        pfunk := flag.String("p","ftp","Protocol/Service to brute")
        pfile := flag.String("f","","Input file name, if none input=Stdin.")
        tlsvr := flag.Bool("t",true,"Skip TLS TrustChain verification") 
        tail := flag.Bool("xt",false,"Make all tail strings as extra value to protocol parser")       
        flag.Parse()
        
        var extra []string
        if *tail {extra=flag.Args()}
        
        var in *os.File
        var err error   
        initProts()       
        checkErr(chkFu(*pfunk))
        funk:=FuProts[*pfunk]
        fmt.Print("[*] Loaded protocols: ")
        for k, _ := range FuProts {
                fmt.Print(k+" ")
        }
        fmt.Print("\n")   
        Tout=*tout
        tlsVerify= *tlsvr         
        pass:= make([]string,*buff)
        cnt:=0 
        fmt.Print("[*] Input from:")
        if (*pfile== ""){              
                in=os.Stdin
                fmt.Println(" Stdin")
        } else {
                fmt.Println(" "+ *pfile)
                in, err = os.Open(*pfile)
                checkErr(err)                
        }
        scanner := bufio.NewScanner(in)
        scanner.Split(bufio.ScanWords)        
        for scanner.Scan() {
// Old good way of memory optimization, don't need to rely on garbage collector for 4KB of data for every loop :)
                if (cnt==*buff){
                        cnt=0
                        monkey(funk,*thrs,*hostp,*user,pass,extra,nil)                        
                }                        
                pass[cnt]=scanner.Text()                
                cnt++
        }        
        if (cnt >0){
                pass=pass[0:cnt:cnt+1]
                monkey(funk,*thrs,*hostp,*user,pass,extra,nil)     
        } else {fmt.Println("ZERO")}       
        wg.Wait()
}
    
