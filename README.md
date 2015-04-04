# tree
A simple tree command clone written in Go

## Installation

To install the $(GOPATH)/bin/tree executable:

    go get github.com/boynton/tree

## Usage

    usage: tree [-a] [-F] dir ...

Assuming $(GOPATH)/bin is in your PATH:

    $ tree
    .
    ├── LICENSE
    ├── README.md
    └── tree.go

	$ tree /tmp/treetest/
	treetest
	├── file1.txt
	└── subdir
	    ├── file2.txt
	    └── src
	$ tree -F /tmp/treetest/
	treetest
	├── file1.txt
	└── subdir/
	    ├── file2.txt
	    └── src@
	$ tree -a /tmp/treetest/
	treetest
	├── .hidden1.txt
	├── file1.txt
	└── subdir
	    ├── file2.txt
	    └── src

