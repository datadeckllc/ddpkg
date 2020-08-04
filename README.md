# ddpkg - Data Deck Go packages
# Installation
1.  Install Golang (https://golang.org/doc/install)
1.  Run `bash build.sh` to build the Go packages
1.  Click the Code button and copy the URL
1.  Create a directory to contain all of the datadeck code. e.g., `mkdir -p ~/projects/datadeckllc`
1.  `cd ~/projects/datadeckllc`
1.  Run `git clone <URL>`.
1.  Run `bash build.sh`


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
