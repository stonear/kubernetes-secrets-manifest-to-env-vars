package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	// Define command-line flags
	filenameInput := flag.String("input", "example.yaml", "Your yaml file input")
	filenameOutput := flag.String("output", "example.env.sh", "Your sh file output")
	flag.Parse()

	viper.SetConfigFile(*filenameInput)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error reading file:", err)
		return
	}

	datas := viper.GetStringMapString("data")

	keys := make([]string, 0, len(datas))
	for k := range datas {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	file, err := os.Create(*filenameOutput)
	if err != nil {
		log.Fatalln("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, k := range keys {
		decodedValue, err := base64.StdEncoding.DecodeString(datas[k])
		if err != nil {
			log.Fatalln("Error decoding value:", err)
			return
		}
		_, err = file.WriteString(fmt.Sprintf("export %s='%s'\n", strings.ToUpper(k), escapeSpecialChars(string(decodedValue))))
		if err != nil {
			log.Fatalln("Error writing to file:", err)
			return
		}
	}

	fmt.Println("File converted successfully:", *filenameOutput)
}

func escapeSpecialChars(input string) string {
	escaped := strings.ReplaceAll(input, "\\", "\\\\")  // Escape backslashes
	escaped = strings.ReplaceAll(escaped, "\n", "\\n")  // Newline
	escaped = strings.ReplaceAll(escaped, "\t", "\\t")  // Tab
	escaped = strings.ReplaceAll(escaped, "\r", "\\r")  // Carriage return
	escaped = strings.ReplaceAll(escaped, "'", "\\'")   // Single quote
	escaped = strings.ReplaceAll(escaped, "\"", "\\\"") // Double quote
	escaped = strings.ReplaceAll(escaped, "$", "\\$")   // Dollar sign
	escaped = strings.ReplaceAll(escaped, "`", "\\`")   // Backtick
	// Add more replacements as needed
	return escaped
}
