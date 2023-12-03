package versions

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Version3(destinationFilePath string) {
	referenceFilePath := "files/.env"
	// Lê o conteúdo do arquivo de referência
	referenceFileData, err := os.ReadFile(referenceFilePath)

	if err != nil {
		log.Fatal(err)
	}

	// Abre ou cria o arquivo de destino
	destinationFile, err := os.OpenFile(destinationFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer destinationFile.Close()

	// Lê o conteúdo do arquivo de destino
	destinationFileData, err := os.ReadFile(destinationFilePath)

	if err != nil {
		log.Fatal(err)
	}

	destinationFileData = bytes.Trim(destinationFileData, "\n")

	var variablesToAdd []byte
	var variableNamesToAdd [][]byte

	// Caracteres convertidos para byte
	breakLine := []byte("\n")
	emptySpace := []byte(" ")

	// Caso o arquivo de destino esteja vazio, não é necessário fazer nenhuma vefiricação
	if len(destinationFileData) == 0 {
		variablesToAdd = append(variablesToAdd, referenceFileData...)
	} else {
		referenceFileVariables := bytes.Split(referenceFileData, breakLine)
		destinationFileVariables := bytes.Split(destinationFileData, breakLine)

		// Cria um loop com base nas variáveis do arquivo de referência
		for _, referenceFileVariable := range referenceFileVariables {
			endVariableNameIndex := bytes.Index(referenceFileVariable, emptySpace)
			if endVariableNameIndex == -1 {
				break
			}

			// Obtém o nome da variável de ambiente do arquivo de referência
			referenceFileVariableName := referenceFileVariable[:endVariableNameIndex]

			// Cria um loop com base nas variáveis do arquivo de destino
			// Tem como finalidade verificar se a variável do arquivo de referência existe no arquivo de destino
			for destinationFileVariableIndex, destinationFileVariable := range destinationFileVariables {
				// Caso a linha atual esteja vazia ou seja um comentário, o fluxo de execução passa para a linha seguinte
				if len(destinationFileVariable) == 0 || destinationFileVariable[0] == '#' {
					continue
				}

				endTargetNameIndex := bytes.Index(destinationFileVariable, emptySpace)
				if endTargetNameIndex == -1 {
					break
				}

				// Obtém o nome da variável de ambiente do arquivo de destino
				destinationFileVariableName := destinationFileVariable[:endTargetNameIndex]

				// Verifica se a variável do arquivo de referência é igual a variável atual do loop do arquivo de destino
				if bytes.Equal(referenceFileVariableName, destinationFileVariableName) {
					fmt.Printf("Variável %s já existe no arquivo .env.\n", referenceFileVariableName)
					break
				}

				// Adiciona a variável caso ela não exista no arquivo de destino
				if destinationFileVariableIndex == len(destinationFileVariables)-1 {
					variableNamesToAdd = append(variableNamesToAdd, referenceFileVariableName)
					variablesToAdd = append(variablesToAdd, breakLine...)
					variablesToAdd = append(variablesToAdd, referenceFileVariable...)
				}
			}
		}
	}

	// Caso não haja variáveis para adicionar, o programa é encerrado.
	if len(variablesToAdd) == 0 {
		return
	}

	// Limpa o arquivo de destino
	destinationFile.Truncate(0)

	// Uni o código do arquivo de destino com as variáveis que serão adicionadas
	newCode := append(destinationFileData, variablesToAdd...)
	newCode = append(newCode, breakLine...)

	// Escreve no arquivo de destino
	destinationFile.Write(newCode)

	// Caso o arquivo de destino esteja vazio, significa que todas as variáveis do arquivo de referência serão adicionadas
	if len(variableNamesToAdd) == 0 {
		fmt.Println("Acesse o arquivo .env e preencha a(s) variável(eis) de ambiente.")
		return
	}

	// Variável responsável por informar quais variáveis de ambiente foram adicionadas
	variableNames := string(variableNamesToAdd[0])

	// Caso tenha sido adicionado mais de uma variável do arquivo de referência, a variável variableNames é atualizada
	if len(variableNamesToAdd) > 1 {
		for variableNameToAddIndex := 1; variableNameToAddIndex < len(variableNamesToAdd); variableNameToAddIndex++ {
			if variableNameToAddIndex == len(variableNamesToAdd)-1 {
				variableNames = fmt.Sprintf("%s e %s", string(variableNames), string(variableNamesToAdd[variableNameToAddIndex]))
				break
			}

			variableNames = fmt.Sprintf("%s, %s", string(variableNames), string(variableNamesToAdd[variableNameToAddIndex]))
		}
	}

	fmt.Printf("Acesse o arquivo .env e preencha a(s) variável(eis) de ambiente %s.\n", variableNames)
}
