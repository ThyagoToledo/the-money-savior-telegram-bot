# Arquitetura e Estrutura de Dados

Este documento descreve a infraestrutura tecnica do bot Money Savior, o fluxo de comunicacao entre componentes, a definicao das estruturas de dados em Go e o schema de tabelas da AWS.

---

## Diagrama Arquitetural

A comunicacao do bot opera em duas vias de execucao (hospedagem tradicional baseada em servidor ou execucao serverless via AWS Lambda) integrada ao DynamoDB:

```text
┌─────────────────────────────────────────────────────────┐
│                  Telegram Bot (Go)                       │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  Bot Routes / Message Handlers                      │ │
│  │  - /gastei  - /consulta  - /deletar               │ │
│  │  - /deletartudo  - /help  - /start                │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────┐
│           AWS Lambda + API Gateway                       │
│  (Opcional para deploy serverless)                       │
└─────────────────────────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────┐
│              AWS DynamoDB                               │
│  Tabela: expenses                                        │
│  PK: user_id | SK: expense_id                          │
│  Atributos: amount, category, method, seq_id           │
└─────────────────────────────────────────────────────────┘
```

---

## Estrutura de Modelos em Go

Os dados de despesas sao mapeados no codigo Go utilizando a seguinte estrutura de dados:

```go
type Expense struct {
    UserID    int64     // ID do usuario Telegram (Partition Key)
    ChatID    int64     // ID do chat do Telegram
    Username  string    // Username do usuario no Telegram
    Amount    float64   // Valor financeiro do gasto
    Category  string    // Categoria da despesa
    Method    string    // Metodo de pagamento (pix, debito, credito)
    CreatedAt time.Time // Carimbo de data/hora de criacao
    ExpenseID string    // Sort Key DynamoDB (formato: user_id#timestamp)
    SeqID     int       // Identificador sequencial numerico (1, 2, 3...)
}
```

---

## Schema da Tabela DynamoDB

O banco de dados Amazon DynamoDB utiliza uma chave primaria composta para organizar os registros de forma escalavel e isolada:

- **Nome da Tabela**: `expenses`
- **Chave Primaria**:
  - **Chave de Particao (Partition Key / PK)**: `user_id` (Tipo: Number) - Garante o isolamento dos dados por usuario.
  - **Chave de Ordenacao (Sort Key / SK)**: `expense_id` (Tipo: String) - Identificador unico baseado na juncao do ID do usuario e carimbo de tempo.
- **Atributos Cadastrados**:
  - `seq_id` (Number): Indice sequencial para ordenacao e exibicao ao usuario.
  - `amount` (Number): Valor flutuante da despesa.
  - `category` (String): Descritivo de categoria.
  - `method` (String): Metodo de pagamento utilizado.
  - `created_at` (String): Data de criacao formatada sob a norma RFC3339.
  - `username` (String): Nome de usuario do Telegram.
  - `chat_id` (Number): Canal do chat de origem.

---

## Padrao de Logs e Depuracao

O bot mantem a geracao de logs padronizados enviados para a saida padrao (stdout), facilitando o monitoramento de erros de execucao no console local ou na console CloudWatch da AWS:

- `[INFO]`: Mensagens gerais de inicializacao, conexao de API bem sucedida e comandos recebidos.
- `[WARN]`: Avisos nao fatais, como comandos incorretos do usuario ou formatacao invalida de parametros.
- `[ERROR]`: Falhas de conexao, falha na gravacao do DynamoDB ou erros de chamada na API do Telegram.
- `[DEBUG]`: Informacoes de desenvolvimento para depuracao minuciosa de objetos e responses de APIs.
