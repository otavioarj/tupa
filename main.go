package main

import "./tupa"

func main(){
// Example of adding a new protocol	
	tupa.FuProts["ftps"]=tupa.Ftpscon // Ftps is at ftp.go	
	tupa.Start() 
}

