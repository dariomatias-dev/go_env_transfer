package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	baseFilePath := "files/.env"
	// Lê o código do arquivo base
	baseFileData, err := os.ReadFile(baseFilePath)

	if err != nil {
		log.Fatal(err)
	}

	targetFilePath := ".env"

	// Abre ou cria o arquivo alvo
	targetFile, err := os.OpenFile(targetFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer targetFile.Close()

	// Lê o código do arquivo alvo
	targetFileData, err := os.ReadFile(targetFilePath)

	if err != nil {
		log.Fatal(err)
	}

	targetFileData = bytes.TrimRight(targetFileData, "\n")

	var variablesToAdd []byte

	// Caracteres convertidos para byte
	breakLine := []byte("\n")
	emptySpace := []byte(" ")

	// Caso o arquivo alvo esteja vazio, não é necessário fazer nenhuma vefiricação
	if len(targetFileData) == 0 {
		variablesToAdd = append(variablesToAdd, baseFileData...)
	} else {

		baseFileVariables := bytes.Split(baseFileData, breakLine)
		targetFileVariables := bytes.Split(targetFileData, breakLine)

		// Cria um loop com base nas variáveis do arquivo base
		for _, baseFileVariable := range baseFileVariables {
			endVariableNameIndex := bytes.Index(baseFileVariable, emptySpace)
			if endVariableNameIndex == -1 {
				break
			}

			// Obtém o nome da variável de ambiente do arquivo base
			baseFileVariableName := baseFileVariable[:endVariableNameIndex]

			// Cria um loop com base nas variáveis do arquivo alvo
			// Tem como finalidade verificar se a variável do arquivo base existe no arquivo alvo
			for targetFileVariableIndex, targetFileVariable := range targetFileVariables {
				// Caso a linha atual esteja vazia ou seja um comentário, o fluxo de execução passa para a linha seguinte
				if len(targetFileVariable) == 0 || targetFileVariable[0] == '#' {
					continue
				}

				endTargetNameIndex := bytes.Index(targetFileVariable, emptySpace)
				if endTargetNameIndex == -1 {
					break
				}

				// Obtém o nome da variável de ambiente do arquivo alvo
				targetFileVariableName := targetFileVariable[:endTargetNameIndex]

				// Verifica se a variável do arquivo base é igual a variável atual do loop do arquivo alvo
				if bytes.Equal(baseFileVariableName, targetFileVariableName) {
					fmt.Printf("Variável %s já existe no arquivo .env\n", baseFileVariableName)
					break
				}

				// Adiciona a variável caso ela não exista no arquivo alvo
				if targetFileVariableIndex == len(targetFileVariables)-1 {
					variablesToAdd = append(variablesToAdd, breakLine...)
					variablesToAdd = append(variablesToAdd, baseFileVariable...)
				}
			}
		}
	}

	// Limpa o arquivo alvo.
	targetFile.Truncate(0)

	// Uni o código do arquivo alvo com as variáveis que serão adicionadas
	newCode := append(targetFileData, variablesToAdd...)
	newCode = append(newCode, breakLine...)

	// Escreve no arquivo alvo
	targetFile.Write(newCode)
}
