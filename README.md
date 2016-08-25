## Remote Access

# Build it as your convenience

Binary name, default region and default security group can be set at build time with environment variables `REMOTE_ACCESS_BASENAME`, `REMOTE_ACCESS_REGION` and `REMOTE_ACCESS_SG`


``` sh
$ REMOTE_ACCESS_BASENAME=company-remote-access REMOTE_ACCESS_VERSION=0.1  REMOTE_ACCESS_REGION=us-west-1 REMOTE_ACCESS_SG=sg-xxxxxxxx make build -j 3
GOOS=linux GOARCH=amd64 go build -o company-remote-access-linux-amd64-0.1 -ldflags "-X main.version=0.1 -X main.securitygroup=sg-xxxxxxxx -X main.region=us-west-1"
GOOS=windows GOARCH=amd64 go build -o company-remote-access-windows-amd64-0.1 -ldflags "-X main.version=0.1 -X main.securitygroup=sg-xxxxxxxx -X main.region=us-west-1"
GOOS=darwin GOARCH=amd64 go build -o company-remote-access-darwin-amd64-0.1 -ldflags "-X main.version=0.1 -X main.securitygroup=sg-xxxxxxxx -X main.region=us-west-1"
$
$ ls -1 company-*
company-remote-access-darwin-amd64-0.1
company-remote-access-linux-amd64-0.1
company-remote-access-windows-amd64-0.1
$
$ ./company-remote-access-linux-amd64-0.1 -help
Usage of ./company-remote-access-linux-amd64-0.1:
  -dry-run
    	Test if the action is possible but do not actually do it
  -ip string
    	Give the IP address to add. If not given current IP is used
  -region string
    	Override region (default "us-west-1")
  -security-group string
    	Override security group (default "sg-xxxxxxxx")
  -version
    	Display the version
```

Binary is built for 3 os/arch couples that are linux/amd64, windows/amd64 and darwin/amd64.

The version can be overriden using the environment variable `REMOTE_ACCESS_VERSION`. Version can be retrieved with the flag `-version`

``` sh
$ ./company-remote-access-linux-amd64-0.1 -version
0.1
```
