#!/bin/bash
cd cases/
go test -json | go-test-report -o ../test_report.html