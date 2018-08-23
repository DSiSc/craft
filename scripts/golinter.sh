#!/bin/bash
#
# Copyright Greg Haskins All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#


set -e

# Add .go directory
declare -a arr=(
"./core"
)

for i in "${arr[@]}"
do
    echo ">>>Checking code under $i/"

    echo "Checking with gofmt"
    OUTPUT="$(gofmt -l -s ./$i)"
    if [[ $OUTPUT ]]; then
        echo "The following files contain gofmt errors"
        echo "$OUTPUT"
        echo "The gofmt command 'gofmt -l -s -w' must be run for these files"
        exit 1
    fi

    echo "Checking with goimports"
    OUTPUT="$(goimports -srcdir $GOPATH/src/github.com/hyperledger/fabric -l $i)"
    if [[ $OUTPUT ]]; then
        echo "The following files contain goimports errors"
        echo $OUTPUT
        echo "The goimports command 'goimports -l -w' must be run for these files"
        exit 1
    fi

    echo "Checking with go vet"
    OUTPUT="$(go vet $i/...)"
    if [[ $OUTPUT ]]; then
        echo "The following files contain go vet errors"
        echo $OUTPUT
        exit 1
    fi
done
