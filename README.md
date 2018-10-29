## Set the Environment Variable

### For Windows

    > set GOROOT="C:\go"
    > set GOPATH="%USERPROFILE%\go"
    > set GOBIN="%USERPROFILE%\go\bin"
    > set PATH="%PATH%;%GOROOT%\bin;%GOBIN%"

### For Linux

    $ vim ~/.bash_profile

    # set the golang environment
    export GOROOT="/usr/local/go"
    export GOPATH="$HOME/go/path"
    export GOBIN="$HOME/go/path/bin"
    export PATH="$PATH:$GOROOT/bin:$GOBIN"

## Install Go Package for Visual Studio Code
    go get -u -v github.com/mdempsky/gocode
    go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
    go get -u -v github.com/ramya-rao-a/go-outline
    go get -u -v github.com/acroca/go-symbols
    go get -u -v golang.org/x/tools/cmd/guru
    go get -u -v golang.org/x/tools/cmd/gorename
    go get -u -v github.com/derekparker/delve/cmd/dlv
    go get -u -v github.com/rogpeppe/godef
    go get -u -v golang.org/x/tools/cmd/godoc
    go get -u -v github.com/sqs/goreturns
    go get -u -v github.com/golang/lint/golint
    go get -u -v github.com/cweill/gotests/...
    go get -u -v github.com/fatih/gomodifytags
    go get -u -v github.com/josharian/impl
    go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
    go get -u -v github.com/haya14busa/goplay/cmd/goplay
    go get -u -v golang.org/x/net/context

## Install Go Package for Max Exchange API

    go get -u -v github.com/maicoin/max-exchange-api-go

## Warning

- The CoinmarketCap_*_k.csv files are created by [cryptoCMD](https://github.com/guptarohit/cryptoCMD) Project

## LICENSE

    <one line to give the program's name and a brief idea of what it does.>
    Copyright (C) <year>  <name of author>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
