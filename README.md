# Money Savior - Telegram Bot

<p align="center">
  <img src="Icons/Logo.png" alt="Money Savior Logo" width="350px" style="border-radius: 24px; box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);" />
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/AWS-232F3E?style=for-the-badge&logo=amazon-aws&logoColor=FF9900" alt="AWS" />
  <img src="https://img.shields.io/badge/DynamoDB-527FFF?style=for-the-badge&logo=amazon-dynamodb&logoColor=white" alt="DynamoDB" />
  <img src="https://img.shields.io/badge/Telegram-26A5E4?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram" />
  <img src="https://img.shields.io/badge/AWS%20Lambda-FF9900?style=for-the-badge&logo=aws-lambda&logoColor=white" alt="AWS Lambda" />
</p>

Um assistente pessoal de controle de gastos integrado ao Telegram com suporte a IDs sequenciais, navegacao inteligente por botoes inline e gerenciamento completo de despesas financeiras em nuvem.

---

## Estrutura do Projeto

```
the-money-savior-telegram-bot/
├── 📁 Icons/                  # 🆕 Imagens de exibição do repositório
│   └── 📄 Logo.png            # Logotipo do bot no README
├── 📁 doc/                    # 🆕 Documentação modularizada
│   ├── arquitetura.md         # Estrutura de dados e arquitetura AWS
│   ├── desenvolvimento.md     # Guia de instalação, deploy e variáveis
│   ├── funcionalidades.md     # Camadas de proteção e comandos
│   └── readme_standards.md    # Padrões de documentação e estilo do repositório
├── 📁 cmd/                    # Entry points da aplicação
│   ├── 📁 bot/
│   │   └── main.go            # Entry point do bot (execução local)
│   └── 📁 lambda/
│       ├── main.go            # Executável do Handler AWS Lambda
│       └── deploy.sh          # Script bash de deploy serverless
├── 📁 internal/               # Lógica de negócio privada
│   ├── 📁 bot/
│   │   └── telegram.go        # Roteamento e listeners de mensagens
│   ├── 📁 database/
│   │   └── dynamodb.go        # SDK AWS de integração com DynamoDB
│   ├── 📁 handlers/
│   │   ├── start.go           # Comando /start
│   │   ├── help.go            # Comando /help
│   │   ├── expense.go         # Comando /gastei
│   │   ├── query.go           # Comando /consulta (com navegação)
│   │   ├── delete.go          # Comandos /deletar e /deletartudo
│   │   └── invalid.go         # Handler de comando inválido
│   └── 📁 models/
│       └── expense.go         # Definição do modelo Go da despesa
├── 📄 go.mod                  # Declaração de dependências Go
├── 📄 go.sum                  # Verificação de somas das dependências
├── 📄 README.md               # Este arquivo (Hub)
└── 📄 .gitignore              # Arquivos ignorados pelo Git
```

---

## Hub de Documentacao

A documentacao tecnica detalhada sobre o bot esta dividida e organizada na pasta [doc/](doc/):

* **[Funcionalidades e Comandos](doc/funcionalidades.md)**: Detalhes de execucao, comandos de interacao do Telegram, fluxos do usuario (gastei, consulta, delete) e roadmap futuro.
* **[Arquitetura e Banco de Dados](doc/arquitetura.md)**: Diagrama arquitetural (AWS Lambda + DynamoDB), mapeamento de models Go, schema de tabelas no DynamoDB e padrao de logs.
* **[Guia de Desenvolvimento e Deploy](doc/desenvolvimento.md)**: Instrucoes de pre-requisitos locais, configuracao de credenciais, variaveis de ambiente e roteiro de deploy serverless.
* **[Padroes de Documentacao e Estilo](doc/readme_standards.md)**: Manual de estilo corporativo sobre estrutura de arquivos, regras de logo, badges e autor.

---

## Inclusao Rapida (Quick Start)

### Execucao Local
Defina as configuracoes essenciais no console e inicie o executavel Go:

No Windows (PowerShell):
```powershell
$env:TELEGRAM_BOT_TOKEN="token_do_bot"
$env:TABLE_NAME="expenses"
go run cmd/bot/main.go
```

No Linux / macOS:
```bash
export TELEGRAM_BOT_TOKEN="token_do_bot"
export TABLE_NAME="expenses"
go run cmd/bot/main.go
```

---

## Autores

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/AmrmDev">
        <img src="https://github.com/AmrmDev.png" width="100px;" alt="AmrmDev"/>
        <br />
        <sub><b>AmrmDev</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/ThyagoToledo">
        <img src="https://github.com/ThyagoToledo.png" width="100px;" alt="Thyago Toledo"/>
        <br />
        <sub><b>Thyago Toledo</b></sub>
      </a>
    </td>
  </tr>
</table>

---

Este projeto e disponibilizado sob os termos da licenca MIT. Para mais detalhes consulte o arquivo de licenca do repositorio.
