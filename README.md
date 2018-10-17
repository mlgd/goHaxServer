# goHaxServer

goHaxServer is a HaxServer alternative written with Go.

## Features

* Cross-platform server to launch kernel exploit on WiiU
* DNS server to filter access to update servers
* Responsive web page for best usage on WiiU internet browser

## Usage

1. Configure your server with file `conf/app.conf`
2. Run goHaxServer from command line :
   * Linux or MacOS : `./goHaxServer` or `sudo ./goHaxServer`
   * Windows : `goHaxServer`
3. Go to homepage goHaxServer with server IP on WiiU internet browser. Example : http://192.168.0.1:8080/

## Recommendations

Run goHaxServer as root or with root privilege to configure HTTP server on port 80.
Check you don't have another web server running, and do the same for DNS server to use the embedded server.

## Source dependencies

* [beego] (https://github.com/astaxie/beego)
* [DNS library] (https://github.com/miekg/dns)
