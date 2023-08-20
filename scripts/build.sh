#!/bin/bash
rm build/checker build/push-swap
cd ../checker/
go build .
mv checker ../scripts/build
cd ../push_swap/
go build .
mv push-swap ../scripts/build