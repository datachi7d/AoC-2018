#!/bin/bash

go run Day1.go -input "`cat Day1.input | tr "\n" "," | head -c -1`"
