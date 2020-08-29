#!/bin/bash

chmod u+x install.sh

cd src
go build -o okra
mv okra ../bin
cd ..