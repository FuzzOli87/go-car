#!/bin/bash
GOARM=7 GOARCH=arm GOOS=linux go build && scp go-car pi@192.168.86.34:Development/go/
