#!/bin/bash

# Collect build data
BUILDMACHINE=`uname -n`
BUILDUSER=`whoami`
GOVER=`go version | awk -F' ' '{print $3}'`
BUILDARCH="`go env GOHOSTOS`/`go env GOHOSTARCH`"
TARGETARCH="`go env GOOS`/`go env GOARCH`"
BUILDTS=`date +"%s"`
GITINFO=`git describe --long --tags`
DEVBUILD="True"

# Build
echo "Building Pravaah ..."
echo "  Build go version    : $GOVER"
echo "  Build machine       : $BUILDMACHINE"
echo "  Build user          : $BUILDUSER"
echo "  Build host arch     : $BUILDARCH"
echo "  Build target arch   : $TARGETARCH"
echo "  Build timestamp     : $BUILDTS"
echo "  Build git info      : $GITINFO"

go get
go install  -ldflags "-X pravaah/version.BuildMachine=$BUILDMACHINE \
                      -X pravaah/version.BuildUser=$BUILDUSER \
                      -X pravaah/version.GOVersion=$GOVER \
                      -X pravaah/version.BuildArch=$BUILDARCH \
                      -X pravaah/version.TargetArch=$TARGETARCH \
                      -X pravaah/version.BuildTS=$BUILDTS \
                      -X pravaah/version.GITInfo=$GITINFO \
                      -X pravaah/version.DevBuild=$DEVBUILD"