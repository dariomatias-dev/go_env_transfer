package versions

import (
	"bytes"
	"log"
	"os"
)

func Version1() {
	filePath := "files/.env"

	baseCodeVariableNames, baseCodeVariables, err := getVariableNames(filePath)
	if err != nil {
		log.Fatal(err)
	}

	filePath = ".env"
	targetCodeVariableNames, _, err := getVariableNames(filePath)
	if err != nil {
		log.Fatal(err)
	}

	variablesToAdd := []string{}

	isPresent := false
	for baseCodeVariableIndex, baseCodeVariableName := range *baseCodeVariableNames {
		for _, tarbaseCodeVariableName := range *targetCodeVariableNames {
			if baseCodeVariableName == tarbaseCodeVariableName {
				isPresent = true
			}
		}

		if !isPresent {
			variablesToAdd = append(variablesToAdd, (*baseCodeVariables)[baseCodeVariableIndex])
		}

		isPresent = false
	}

	addVariables(filePath, variablesToAdd)
}

func getVariableNames(filePath string) (*[]string, *[]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	variableNames := []string{}
	variables := []string{}

	for {
		endIndex := bytes.Index(data, []byte("\n"))
		if endIndex == -1 {
			break
		}

		variable := data[:endIndex]
		variables = append(variables, string(variable))
		data = data[endIndex+1:]

		endVariableNameIndex := bytes.Index(variable, []byte(" "))
		variableName := variable[:endVariableNameIndex]
		variableNames = append(variableNames, string(variableName))
	}

	return &variableNames, &variables, nil
}

func addVariables(filePath string, variableIndicesAdd []string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	code, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	codeBuffer := code

	for _, variable := range variableIndicesAdd {
		codeBuffer = append(codeBuffer, []byte(variable)...)
		codeBuffer = append(codeBuffer, []byte("\n")...)
	}

	file.Write(codeBuffer)

	if err := file.Truncate(0); err != nil {
		return err
	}

	_, err = file.Write(codeBuffer)
	if err != nil {
		return err
	}

	return nil
}
