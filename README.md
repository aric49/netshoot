# netshoot
Quick and dirty endpoint troubleshooting CLI utility.  For any endpoint pass as a CLI argument, netshoot will perform local and public DNS Lookups, as well as `ping` and HTTP connectivity!

## Usage
Simply build the code and provide an endpoint you're trying to troubleshoot:

```
$ go build -o netshoot main.go

$ ./netshoot https://google.com
NetTroubleshoot Starting......
Troubleshooting connection to google.com
=================================================
 Performing DNS Lookup on  google.com
DNS lookup successfully completed on  google.com
Server:         192.168.1.1
Address:        192.168.1.1#53

Non-authoritative answer:
Name:   google.com
Address: 172.217.15.110
Name:   google.com
Address: 2607:f8b0:4004:811::200e


=================================================
Performing Public DNS Lookup on  google.com using 1.1.1.1......
DNS Lookup successfully completed on google.com
Server:         1.1.1.1
Address:        1.1.1.1#53

Non-authoritative answer:
Name:   google.com
Address: 108.177.122.102
Name:   google.com
Address: 108.177.122.101
Name:   google.com
Address: 108.177.122.100
Name:   google.com
Address: 108.177.122.139
Name:   google.com
Address: 108.177.122.138
Name:   google.com
Address: 108.177.122.113
Name:   google.com
Address: 2607:f8b0:4002:c09::64


=================================================
Performing Ping on google.com
Ping Successfully Completed
PING google.com (172.217.15.110) 56(84) bytes of data.
64 bytes from iad30s21-in-f14.1e100.net (172.217.15.110): icmp_seq=1 ttl=54 time=17.5 ms
64 bytes from iad30s21-in-f14.1e100.net (172.217.15.110): icmp_seq=2 ttl=54 time=17.5 ms
64 bytes from iad30s21-in-f14.1e100.net (172.217.15.110): icmp_seq=3 ttl=54 time=17.7 ms
64 bytes from iad30s21-in-f14.1e100.net (172.217.15.110): icmp_seq=4 ttl=54 time=17.5 ms
64 bytes from iad30s21-in-f14.1e100.net (172.217.15.110): icmp_seq=5 ttl=54 time=17.5 ms

--- google.com ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4006ms
rtt min/avg/max/mdev = 17.455/17.522/17.679/0.080 ms

=================================================
Calling endpoint,  https://google.com
HTTP Request to, https://google.com HTTP Status Code:  200
```


Running netshoot with the -v flag shows the verbose output of the HTTP call:

```
=================================================
Calling endpoint,  https://google.com
HTTP Request to, https://google.com HTTP Status Code:  200
Calling endpoint,  https://google.com
Verbose HTTP Request Output:  <!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en"><head><meta content="Search the world's information, including webpages, images, videos and more. Google has many special features to help you find exactly what you're looking for." name="description"><meta content="noodp" name="robots"><meta content="text/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>Google</title>.....
```