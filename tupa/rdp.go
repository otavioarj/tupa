package tupa

import (
	"fmt"
	"net"	
	"github.com/icodeface/grdp/core"	
	"github.com/icodeface/grdp/protocol/nla"
	"github.com/icodeface/grdp/protocol/pdu"
	"github.com/icodeface/grdp/protocol/sec"
	"github.com/icodeface/grdp/protocol/t125"
	"github.com/icodeface/grdp/protocol/tpkt"
	"github.com/icodeface/grdp/protocol/x224"
	"strings"
	"time"
)


type Client struct {
	Host string // ip:port
	tpkt *tpkt.TPKT
	x224 *x224.X224
	mcs  *t125.MCSClient
	sec  *sec.Client
	pdu  *pdu.Client
}

func Rdpcon(hostp string,user string, pass []string, infos []string,con net.Conn) {        
     var err error
     if (con ==nil){
             con, err = net.DialTimeout("tcp", hostp,time.Duration(Tout) *time.Second)
             checkErr(err)                
     }   
	defer con.Close()
	var g Client
	g.Host=hostp
	domain := strings.Split(hostp, ":")[0]
	
	for _,p := range pass {
		if (wg.pok>0){break}
		g.sec.SetPwd(p)
		g.tpkt = tpkt.New(core.NewSocketLayer(con, nla.NewNTLMv2(domain, user, p)))
		g.x224 = x224.New(g.tpkt)
		g.mcs = t125.NewMCSClient(g.x224)
		g.sec = sec.NewClient(g.mcs)
		g.pdu = pdu.NewClient(g.sec)
		g.sec.SetUser(user)	
		g.sec.SetDomain(domain)
		g.tpkt.SetFastPathListener(g.pdu)
		g.pdu.SetFastPathSender(g.tpkt)
		g.x224.SetRequestedProtocol(x224.PROTOCOL_SSL | x224.PROTOCOL_HYBRID)
		err = g.x224.Connect()
		if err == nil {			
			fmt.Println("[+] Pass: "+p)
			wgGot()
		}	
	}
	wgDone()
}
