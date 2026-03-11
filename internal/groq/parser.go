package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ParsedExpense struct {
	Intent   string  `json:"intent"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Method   string  `json:"method"`
}

type groqRequest struct {
	Model       string        `json:"model"`
	Messages    []groqMessage `json:"messages"`
	Temperature float64       `json:"temperature"`
}

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

const systemPrompt = `Você é um extrator de dados financeiros.
O usuário vai mandar mensagens em linguagem natural sobre gastos ou receitas.
Você deve analisar a mensagem e retornar SOMENTE um JSON válido, sem explicações, sem markdown, sem blocos de código.

O JSON deve seguir exatamente este formato:
{"intent":"expense","amount":0.00,"category":"string","method":"string"}

Regras:
- intent: "expense" para gastos, "income" para receitas/entradas
- amount: valor numérico (ex: 45.50). Se não mencionar valor, use 0
- category: categoria do gasto em português e minúsculas (ex: "uber", "supermercado", "restaurante", "moletom", "farmácia")
- method: método de pagamento — use exatamente um destes valores: "pix", "debito", "credito", "dinheiro". Se não mencionado, use "outros"

Exemplos:
mensagem: "mano gastei 45 de uber no pix"
resposta: {"intent":"expense","amount":45.00,"category":"uber","method":"pix"}

mensagem: "paguei 200 num moletom no debito"
resposta: {"intent":"expense","amount":200.00,"category":"vestuario","method":"debito"}

mensagem: "recebi 1500 de freela hoje"
resposta: {"intent":"income","amount":1500.00,"category":"freela","method":"pix"}

mensagem: "almocei por 32 reais no credito"
resposta: {"intent":"expense","amount":32.00,"category":"restaurante","method":"credito"}`
