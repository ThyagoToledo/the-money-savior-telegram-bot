# Desenvolvimento, Instalacao e Deploy

Este documento contem as instrucoes para configurar o ambiente de desenvolvimento local, guias de deploy serverless na AWS, variaveis de ambiente requeridas, politicas de seguranca e diretrizes para contribuicao de codigo.

---

## Pre-requisitos do Ambiente

Para executar o bot Money Savior localmente ou compilar os arquivos para producao, certifique-se de possuir:

- **Go**: Versao igual ou superior a 1.23.0.
- **AWS CLI**: Credenciais de acesso configuradas na maquina local (com permissao para gravar no DynamoDB).
- **Token de Bot do Telegram**: Gerado atraves do bot oficial do Telegram [@BotFather](https://t.me/botfather).
- **Tabela DynamoDB**: Uma tabela ativa na AWS com o nome `expenses` (configurada conforme o manual de arquitetura).

---

## Instalacao e Execucao Local

### 1. Clonar o Repositorio e Baixar Dependencias
Clone o repositorio do GitHub e acesse a pasta raiz para baixar os pacotes do Go:
```bash
git clone https://github.com/ThyagoToledo/the-money-savior-telegram-bot.git
cd the-money-savior-telegram-bot
go mod download
```

### 2. Configurar Variaveis de Ambiente
Defina as configuracoes necessarias no terminal do seu sistema operacional:

No Linux / macOS:
```bash
export TELEGRAM_BOT_TOKEN="seu_token_aqui"
export AWS_REGION="us-east-1"
export TABLE_NAME="expenses"
```

No Windows (PowerShell):
```powershell
$env:TELEGRAM_BOT_TOKEN="seu_token_aqui"
$env:AWS_REGION="us-east-1"
$env:TABLE_NAME="expenses"
```

### 3. Iniciar o Bot
Execute o entrypoint da aplicacao:
```bash
go run cmd/bot/main.go
```

---

## Deploy Serverless (AWS Lambda)

A extensao foi estruturada para permitir execucao serverless barata e escalavel. Para realizar o empacotamento e deploy na AWS:

1. Acesse o diretorio do Lambda:
   ```bash
   cd cmd/lambda
   ```
2. Execute o script de deploy automatizado:
   ```bash
   ./deploy.sh
   ```
   *Nota: O script compila o executavel Go focado na arquitetura Linux do AWS Lambda, gera o arquivo compactado zip e faz a atualizacao da funcao cadastrada na AWS via AWS CLI.*

---

## Referencia de Variaveis de Ambiente

| Variavel | Descricao | Obrigatoria | Padrao |
|----------|-----------|-------------|--------|
| `TELEGRAM_BOT_TOKEN` | Token secreto do bot do Telegram | Sim | - |
| `TABLE_NAME` | Nome da tabela criada no DynamoDB | Sim | `expenses` |
| `AWS_REGION` | Regiao AWS para conexao de servicos | Nao | `us-east-1` |

---

## Dependencias de Bibliotecas Go

As dependencias principais exigidas pelo projeto estao descritas no arquivo `go.mod`:

- `github.com/aws/aws-lambda-go`: SDK de execucao de funcoes Lambda.
- `github.com/aws/aws-sdk-go-v2`: SDK principal de integracao aos servicos AWS.
- `github.com/aws/aws-sdk-go-v2/config`: Leitura de credenciais locais.
- `github.com/aws/aws-sdk-go-v2/service/dynamodb`: Modulo especifico para persistencia.
- `github.com/go-telegram-bot-api/telegram-bot-api/v5`: Wrapper de comunicacao com a API de bots do Telegram.

---

## Diretrizes de Seguranca do Projeto

Para manter a integridade dos dados e seguranca da aplicacao:
- **Validacao de Entradas**: Todos os comandos recebidos sao limpos e validados antes de qualquer execucao ou persistencia.
- **Isolamento de Contas**: Cada usuario do Telegram so consegue interagir com as despesas associadas ao seu proprio ID.
- **Confirmacoes Obrigatorias**: Processos destrutivos (como remocoes) contam com botoes inline de confirmacao de intencao.
- **Credenciais Seguras**: Chaves de API e segredos AWS sao passados exclusivamente por variaveis de ambiente do sistema, nunca salvos diretamente no codigo.

---

## Guia de Contribuicao

1. Faca um fork do repositorio do projeto no GitHub.
2. Crie uma branch de desenvolvimento local para sua alteracao:
   ```bash
   git checkout -b feature/AmazingFeature
   ```
3. Realize os commits com mensagens limpas e objetivas:
   ```bash
   git commit -m "feat: Add some AmazingFeature"
   ```
4. Envie a branch criada para o seu repositorio remoto:
   ```bash
   git push origin feature/AmazingFeature
   ```
5. Abra uma solicitacao de Pull Request detalhando as modificacoes propostas para revisao.
