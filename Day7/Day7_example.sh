#!/bin/bash

cat input.txt.example | sort | awk '{ print "\42" $0 "\42" }' | tr "\n" " " | head -c -1 | xargs go run Day7.go
