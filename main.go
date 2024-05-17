package main

import (
	"log"
	"magnetflow/torrentfile"
	"os"
)

func main() {
	inPath, outPath := os.Args[1], os.Args[2]
	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)

	if err != nil {
		log.Fatal(err)
	}

}
