# Utils

## sysinfo

sysinfo prints basic system information. I have it run every time I open a new terminal.

The output looks like this:
```
Uptime:	 4 days,  0:21
Loadavg: 0.31, 0.32, 0.38
Root:	 11G / 59G (20%)
```
Install with `go get github.com/jcelliott/utils/sysinfo`

## duration-fmt

duration-fmt takes a duration in milliseconds and prints it in a human-friendly format

Usage looks like this:
```
> duration-fmt 12345
12.3s
> duration-fmt 123456
2m 3s
> duration-fmt 12345678
3h 26m
> duration-fmt 123456789
1d 10h 18m
```

Install with `go get github.com/jcelliott/utils/duration-fmt`

## testserver

testserver starts an http server listening on port 8000 that prints debugging
information about the requests it receives.

Install with `go get github.com/jcelliott/utils/testserver`
