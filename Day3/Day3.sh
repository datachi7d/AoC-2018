#!/bin/bash

cat Day3.input | awk '{ print "\42" $0 "\42" }' | tr "\n" " " | head -c -1 | xargs go run Day3.go
