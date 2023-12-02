package versions

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Version2() {
	targetFilePath := ".env"

	targetFile, err := os.OpenFile(targetFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	targetFileData, err := os.ReadFile(targetFilePath)

	if err != nil {
		log.Fatal(err)
	}

	baseFilePath := "files/.env"

	readFile(targetFile, targetFileData, baseFilePath)

}

func readFile(targetFile *os.File, targetFileData []byte, filePath string) {
	fileData, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	lineBreak := []byte("\n")

	var variable []byte
	var variableName []byte

	var variablesToAdd []byte

	for {
		indexEndVariableName := bytes.Index(fileData, []byte("="))

		if indexEndVariableName == -1 {
			break
		}

		variableName = bytes.TrimRight(fileData[:indexEndVariableName], " ")

		indexEndVariable := bytes.Index(fileData, lineBreak)

		variable = fileData[:indexEndVariable]
		fileData = fileData[indexEndVariable+1:]

		if bytes.Equal(targetFileData, variableName) {
			fmt.Printf("Variável %s já esta presente.\n", string(variableName))
		} else {
			variablesToAdd = append(variablesToAdd, lineBreak...)
			variablesToAdd = append(variablesToAdd, variable...)
		}
	}

	if len(targetFileData) == 0 {
		variablesToAdd = bytes.TrimLeft(variablesToAdd, "\n")
	}

	code := append(targetFileData, variablesToAdd...)

	targetFile.Truncate(0)

	targetFile.Write(code)
}
