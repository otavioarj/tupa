# Tupa

 **Golang manager for bruteforce/password spray protocols attack**

## Build:

     go build

## Use:

     $ ./Tupã -h
    Usage of ./Tupã:
      -b int
        	Input buffer size,i.e., file/memory chunks, ideally power of two. (default 4096)
      -f string
        	Input file name, if none input=Stdin.
      -hp string
        	HostPort String as host:port. (default "127.0.0.1:8080")
      -n int
        	Thread number, prefer n as power of two. (default 16)
      -p string
        	Protocol/Service to brute (default "ftp")
      -t	Skip TLS TrustChain verification (default true)
      -u string
        	User/account for the attack. (default "admin")

### Examples:

    $ ./Tupã -hp 10.2.240.1:21 -n 10 -f passwordlist.txt -p ftp
	...	...	...
    $ cat wordlist.txt | ./Tupã -hp 10.2.240.1:21 -n 10  -u Administrator -p ftps

