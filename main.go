package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	inputPath  string
	outputPath string
)

const (
	chromeExtensionLevelDBKey = "_chrome-extension://chphlpgkkbolifaimnlloiipkdnihall\x00\x01state"
)

func init() {
	const (
		defaultInputPath  = ""
		inputPathUsage    = "LevelDB database path"
		defaultOutputPath = "-"
		outputPathUsage   = "Output file path (\"-\" to print to standard output)"
	)
	flag.StringVar(&inputPath, "input", defaultInputPath, inputPathUsage)
	flag.StringVar(&inputPath, "i", defaultInputPath, inputPathUsage+" (shorthand)")
	flag.StringVar(&outputPath, "output", defaultOutputPath, outputPathUsage)
	flag.StringVar(&outputPath, "o", defaultOutputPath, outputPathUsage)
}

func main() {
	flag.Parse()

	if inputPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	db, err := leveldb.OpenFile(inputPath, &opt.Options{
		ReadOnly: true,
	})
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer db.Close()

	rawData, err := db.Get([]byte(chromeExtensionLevelDBKey), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	t := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	r := transform.NewReader(bytes.NewReader(rawData[1:]), t)

	var data interface{}
	dec := json.NewDecoder(r)
	err = dec.Decode(&data)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	var w io.Writer
	if outputPath == "-" {
		w = os.Stdout
	} else {
		f, err := os.Create(outputPath)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		defer f.Close()
		w = f
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}
