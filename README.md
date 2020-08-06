# ddpkg - Data Deck Go packages
# Installation
1.  Install Golang (https://golang.org/doc/install)
1.  Create a directory to contain all of the datadeck code. e.g., `mkdir -p ~/projects/datadeckllc`
1.  `cd ~/projects/datadeckllc`
1. On this GitHub page, click the `Code` button and copy the URL for this repository.
1.  Run `git clone <URL>`.
1.  Run `cd ddpkg` 
1.  Run `bash build.sh` to build the Go packages


# Using ddpkkg/sosCA as a CLI tool

1.  Run the following command in Terminal from the ddpkg directory
`bash sosCA/run.sh sosCA/<input_file>.csv`

e.g.: `bash sosCA/run.sh sosCA/test.csv`

NOTE: THe `test.csv` already exists.

1.  Look for the .xlsx file created in the root directory.  Open the file and click on the SOSCA tab.


## Using ddpkg as a library
```
package main
import "github.com/datadeckllc/ddpkg/<pkgname>"
```
