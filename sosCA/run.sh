#!/bin/bash

here="$(dirname ${BASH_SOURCE[0]})"

go run "${here}"/sosCA.go $@
