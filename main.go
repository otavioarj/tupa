package main

import "./tupa"

func main(){
// Example of adding a new protocol	
	tupa.FuProts["ftps"]=tupa.Ftpscon // Ftps ftp.go	
	tupa.FuProts["rdp"]=tupa.Rdpcon   // RDP  rdp.go
	tupa.Start() 
}

