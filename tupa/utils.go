package tupa

import (
        "os"
        "fmt"
        "net"
        "time"
)


func checkErr(err error) {

    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}

func monkey(fu Fu,thrs int,hostp string,user string, pass []string, ext []string,con net.Conn) {
        
        if thrs > len(pass) {thrs=len(pass)}
        intv:= len(pass)/thrs
        remain:=float64(len(pass))/float64(thrs) - float64(intv)
        extra:=0
        var init,end,athr int         
        _=init
// Sleep until there is one free thread 
        for (wgCount() >= thrs){time.Sleep(50 * time.Millisecond)}
        athr=thrs-wgCount()
        wgAdd(athr)          
// Get extra value if pass size isn't int div by thr        
        if (remain>0.0) { 
                extra=int(remain*float64(thrs))  
                if extra <1 {extra+=1}                
        }  
        //fmt.Printf("intv:%d remain:%f extra:%d athr:%d\n",intv,remain,extra,athr)     
        for j,init:=0,0 ; j<athr && wg.pok==0; j++ {                             
                end=init + intv -1
                if (extra >0 && extra>=j){ 
                        end+=1
                        if (j>0) {init+=1}
                }                                     
                //fmt.Printf("init:%d end:%d \n",init,end)                                                                          
                go fu(hostp,user, pass[init:end], ext,con)   
                init+=intv
        }           
}
