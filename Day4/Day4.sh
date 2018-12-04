#!/bin/bash

cat $1 | sort -h | awk '{ print "\42" $0 "\42" }' | tr "\n" " " | head -c -1 | xargs go run Day4.go
