#!/bin/bash

windres -o main.syso main.rc
go build -o KafkaGTK.exe  -ldflags "-w -s -H=windowsgui"

