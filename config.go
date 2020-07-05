package main

import (
	"os"
	"path"
)

var (
 	WORKDIR, _ = os.Getwd()
 	TEMPDIR = path.Join(WORKDIR, "TEMP")
 	DATADIR = path.Join(WORKDIR, "DATA")
)