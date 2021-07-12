#!/bin/bash
# This file builds the application backend from the sources.
set -euo pipefail
printf "Building MaX Pondus backend\n"

# Make Go to install necessary packages globally.
set GO111MODULE=off

# Resolve the absolute path of the project root form the script path.
ROOTPATH=$(dirname $(readlink -f $0))
ROOTPATH=$(dirname $ROOTPATH)

# Resolve and store the version of the Go environment.
GOVERSION=$(go version)

# Resolve the folder where to put all build results.
BUILDPATH="$ROOTPATH/bin/backend"

# Resolve the folder which contains the main of the application.
CMDPATH="$ROOTPATH/cmd/backend"

# Show the information related to compilation environment.
printf "Detected environment information\n"
printf "\tProject root\t $ROOTPATH\n"
printf "\tBuild path\t $BUILDPATH\n"      
printf "\tCmd path\t $CMDPATH\n"        
printf "\tGo version\t $GOVERSION\n"

printf "Running tests...\n"
go test -v ./internal/...

# Remove the old build directory to ensure that we get a clean build.
printf "Removing the old build directory if it already exists\n"
[ -d $BUILDPATH ] && rm -rf $BUILDPATH

# Gather and compile the backend source files into an executable.
printf "Compiling...\n"
mkdir -p $BUILDPATH && cd $BUILDPATH
COMPILATIONSTART=$(date)
go build -race $CMDPATH
COMPILATIONEND=$(date)

# Resolve the path of the created executable.
EXECUTABLEPATH=$BUILDPATH/backend

# Show information related to compilation.
printf "Compiling completed:\n"
printf "\tExecutable\t $EXECUTABLEPATH\n"
printf "\tStart time\t $COMPILATIONSTART\n"
printf "\tEnd time\t $COMPILATIONEND\n"
printf "Build completed.\n"
