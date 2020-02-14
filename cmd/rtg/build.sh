#!/bin/sh
FOLDER=./bin
rm -r $FOLDER/*
for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        if [ $GOOS == windows ]; then
            go build -v -o "$FOLDER/$1-$GOOS-$GOARCH.exe"
        else
            go build -v -o "$FOLDER/$1-$GOOS-$GOARCH"
        fi
    done
done
