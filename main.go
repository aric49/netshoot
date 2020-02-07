package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"net/http"
	"net"
	"net/url"
	"strings"
	"io/ioutil"
	"flag"
)


//Let's Troubleshoot some endpoints! 
func main() {

	verbose := flag.Bool("v", false, "Turn on verbose output")
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("ERROR! Need to provide an endpoint to test. IE: https://google.com")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var url_endpoint string = flag.Arg(0)
	parsedUrl, err := url.Parse(url_endpoint)
	if err != nil {
		log.Fatal(err, "ERROR! Check primary URL!")
	}

	if strings.Contains(url_endpoint, "://") {
		fmt.Println("NetTroubleshoot Starting......")
	} else {
		log.Fatal("The URL must contain a scheme! IE:  http://, https://, or tcp://")
	}

	parsedUrlHost, parsedUrlPort, _ := net.SplitHostPort(parsedUrl.Host)

	if parsedUrlPort == "" {
		if parsedUrl.Scheme == "https" {
			parsedUrlPort = "443"
		} else if parsedUrl.Scheme == "http" {
			parsedUrlPort = "80"
		} else {
			fmt.Println("No port provided, defaulting to 80")
			parsedUrlPort = "80"
		}
	}
	if parsedUrlHost == "" {
		parsedUrlHost = parsedUrl.Host
	}

	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}



	fmt.Println("Troubleshooting connection to", parsedUrlHost)
	fmt.Println("=================================================")
	fmt.Println(dnsLookup(parsedUrlHost))

	fmt.Println("=================================================")
	fmt.Println(dnsLookupPublic(parsedUrlHost))
	fmt.Println("=================================================")

	fmt.Println(pingEndpoint(parsedUrlHost))

	fmt.Println("=================================================")
	fmt.Println("HTTP Request to,", url_endpoint, "HTTP Status Code: ", checkHTTPResponse(url_endpoint))

	if *verbose == true {
		fmt.Println("Verbose HTTP Request Output: ", checkHTTTPResponseVerbose(url_endpoint))
	}

}

	
func checkHTTPResponse(host string) int {
	fmt.Println("Calling endpoint, ", host)
	response, err := http.Get(host)
	if err != nil {
			log.Fatal("An error has occurred!")
	}
	defer response.Body.Close()

	return response.StatusCode
}

func checkHTTTPResponseVerbose(host string) string {
	fmt.Println("Calling endpoint, ", host)
	response, err := http.Get(host)
	if err != nil {
			log.Fatal("An error has occurred!")
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return string(body)
}

func dnsLookup(endpoint string) string {
	fmt.Println(" Performing DNS Lookup on ", endpoint)
	dns, err := exec.Command("nslookup", endpoint).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DNS lookup successfully completed on ", endpoint)
	return string(dns)
}

func dnsLookupPublic(endpoint string) string {
	fmt.Println("Performing Public DNS Lookup on ", endpoint, "using 1.1.1.1......")
	dns, err := exec.Command("nslookup", endpoint, "1.1.1.1").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DNS Lookup successfully completed on", endpoint)
	return string(dns)
}

func pingEndpoint(endpoint string) string {
	fmt.Println("Performing Ping on", endpoint) 
	ping, err := exec.Command("ping", "-c", "5", endpoint).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ping Successfully Completed")
    return string(ping)
}