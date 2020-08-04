#!/bin/bash

set -eax

here="$(dirname ${BASH_SOURCE[0]})"

cd "${here}"/sosCa

go build

go get -v github.com/PuerkitoBio/goquery
go get -v github.com/360EntSecGroup-Skylar/excelize

cd ..




