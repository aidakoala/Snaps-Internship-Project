#!/bin/bash
# I have placed the script in the src dir of go

go build -o /tmp/snap github.com/snapcore/snapd/cmd/snap
go install github.com/snapcore/snapd/cmd/snap/...
go build -o /tmp/snapd github.com/snapcore/snapd/cmd/snapd