# Projeto Go ENV Transfer

Este é um projeto pessoal desenvolvido para explorar a linguagem de programação Go (Golang) e colocar em prática uma ideia que surgiu sem um propósito específico. O objetivo principal é ler um arquivo .env e transferir suas variáveis para outro arquivo, como se estivesse utilizando uma linha de comando que ativa esse processo.

## Descrição do Projeto

A motivação por trás deste projeto foi a curiosidade e a vontade de desenvolver uma solução para uma tarefa específica, mesmo que sem uma aplicação prática imediata. A ideia básica consiste em ler um arquivo .env e copiar todas as suas variáveis para outro arquivo, seguindo algumas condições específicas:

- Não deve escrever uma variável no arquivo de destino se ela já existir, exibindo uma mensagem no console para possível verificação.
- A verificação deve ser realizada mesmo se houver comentários e espaços entre as variáveis no arquivo de destino.
- Mesmo se não houver uma quebra de linha no final do arquivo de destino, a verificação não deve ser afetada.

## Status do Desenvolvimento

O projeto evoluiu ao longo do tempo, com diferentes versões representando mudanças significativas na lógica de implementação. Cada versão foi arquivada em diferentes estágios do desenvolvimento para permitir comparação de resultados no final.

Versões:

- v1, v2: Versões iniciais com códigos que podem ser refatorados.
- v3: Última versão completamente desenvolvida e funcional.

## Reflexões sobre o Desenvolvimento

É importante notar que, embora seja possível alcançar resultados semelhantes com um código mais curto e eficiente, a satisfação pessoal foi alcançada ao considerar o contexto de aprendizado da linguagem Go, que foi estudada nas últimas semanas. O projeto serve como um exercício prático para aplicar conhecimentos recém-adquiridos.

Observação: Refatorações e otimizações são possíveis nos códigos das versões iniciais, mas não foram realizadas porque não convém.

## Como Utilizar

Certifique-se de ter o ambiente Go configurado.

Baixe o repositório:

```bash
git clone https://github.com/dariomatias-dev/go_env_transfer
```

Execute a versão desejada (v3 é a mais recente e completamente desenvolvida).

```bash
go run main.go
```
