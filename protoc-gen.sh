#!/usr/bin/env bash

protoc -I. --gocosmos_out=paths=source_relative:. a/a.proto
protoc -I. --gocosmos_out=paths=source_relative:. b/b.proto
protoc -I. --gocosmos_out=paths=source_relative:. c/c.proto
protoc -I. --gocosmos_out=paths=source_relative:. d/d.proto
protoc -I. --gocosmos_out=paths=source_relative:. e/e.proto
protoc -I. --gocosmos_out=plugins=interfacetype,paths=source_relative:. root.proto
