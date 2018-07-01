#!/bin/bash

for f in $(ls); do
    if [ -d $f ]; then
        cd $f
        echo $f
        go test
        cd ..
    fi
done