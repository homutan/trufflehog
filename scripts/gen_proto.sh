#!/bin/bash

set -eux

protoc -I proto/ \
    -I ${GOPATH}/src \
    -I /usr/local/include \
    -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
    --go_out=./pkg/pb/credentialspb --go_opt=paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./pkg/pb/credentialspb" \
    proto/credentials.proto
protoc -I proto/ \
    -I ${GOPATH}/src \
    -I /usr/local/include \
    -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
    --go_out=./pkg/pb/sourcespb --go_opt=paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./pkg/pb/sourcespb" \
    proto/sources.proto
protoc -I proto/ \
    -I ${GOPATH}/src \
    -I /usr/local/include \
    -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
    --go_out=./pkg/pb/detectorspb --go_opt=paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./pkg/pb/detectorspb" \
    proto/detectors.proto
protoc -I proto/ \
    -I ${GOPATH}/src \
    -I /usr/local/include \
    -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
    --go_out=./pkg/pb/source_metadatapb --go_opt=paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./pkg/pb/source_metadatapb" \
    proto/source_metadata.proto
protoc -I proto/ \
    -I ${GOPATH}/src \
    -I /usr/local/include \
    -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
    --go_out=./pkg/pb/custom_detectorspb --go_opt=paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./pkg/pb/custom_detectorspb" \
    proto/custom_detectors.proto
