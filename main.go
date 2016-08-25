package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	version       = "undefined"
	securitygroup = "undefined"
	region        = "undefined"
)

func main() {

	var err error
	var v bool
	var ip string
	var sg string
	var r string
	var dryRun bool
	flag.BoolVar(&v, "version", false, "Display the version")
	flag.StringVar(&ip, "ip", "", "Give the IP address to add. If not given current IP is used")
	flag.StringVar(&sg, "security-group", securitygroup, "Override security group")
	flag.StringVar(&r, "region", region, "Override region")
	flag.BoolVar(&dryRun, "dry-run", false, "Test if the action is possible but do not actually do it")
	flag.Parse()
	if v {
		fmt.Println(version)
		return
	}
	if dryRun {
		fmt.Println("Runnning in dry run mode")
	}

	if len(ip) == 0 {
		fmt.Print("Getting your IP: ")
		ip, err = myipv4()
		if err != nil {
			fmt.Println("Failed to retrieve your IP,", err)
			return
		}
	} else {
		fmt.Print("Using IP: ")
	}
	fmt.Println(ip)
	fmt.Print("Adding your IP... ")
	sess, err := session.NewSession(&aws.Config{Region: aws.String(r)})
	if err != nil {
		fmt.Println("Failed to create session,", err)
		return
	}
	client := ec2.New(sess)

	err = AddIP(client, sg, ip, dryRun)
	if err != nil {
		fmt.Println("Failed to add IP,", err)
		return
	}
	fmt.Println("done")
}

func myipv4() (string, error) {

	url := "http://ipinfo.io"
	var ipinfo struct {
		Ip       string `json:ip`
		Hostname string `json:hostname,omitempty`
		City     string `json:city,omitempty`
		Region   string `json:region,omitempty`
		Country  string `json:country,omitempty`
		Loc      string `json:loc,omitempty`
		Org      string `json:org,omitempty`
		Postal   string `json:postal,omitempty`
	}
	// Mimic command curl ipinfo.com -v
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "curl/7.43.0")
	req.Header.Add("Accept", "*/*")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ipinfo)
	if err != nil {
		return "", err
	}
	if len(ipinfo.Ip) == 0 {
		return "", errors.New("")
	}
	return ipinfo.Ip, nil
}

func AddIP(c *ec2.EC2, sg, ip string, dry bool) error {

	cidrip := ip + "/32"
	protocol := "-1"
	input := ec2.AuthorizeSecurityGroupIngressInput{
		CidrIp:     &cidrip,
		DryRun:     &dry,
		GroupId:    &sg,
		IpProtocol: &protocol,
	}
	_, err := c.AuthorizeSecurityGroupIngress(&input)
	if err != nil && dry == true && strings.Contains(err.Error(), "DryRunOperation") {
		// Dry run case should not be reported as an error
		return nil
	}
	return err
}
