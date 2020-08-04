# ddpkg - Data Deck Go packages
# Installation
1.  Install Golang (https://golang.org/doc/install)
1.  Run `bash build.sh` to build the Go packages

# Using ddpkkg/sosCA as a CLI tool

1.  Run the following command in Terminal
bash sosCA/run.sh sosCA/<input_file>.csv

e.g.: `bash sosCA/run.sh sosCA/test.csv`

1.  Look for the .xlsx file created in the root directory.  Open the SOSCA tab.


## Using ddpkg as a library
```
package main
import "github.com/datadeckllc/ddpkg/<pkgname>"
```
