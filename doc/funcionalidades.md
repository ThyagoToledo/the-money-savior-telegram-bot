# Funcionalidades e Comandos do Bot

Este documento descreve detalhadamente as features do bot Money Savior, os comandos disponiveis no Telegram, os fluxos internos de execucao e os planos para versoes futuras.

---

## Funcionalidades Principais

O bot Money Savior oferece os seguintes recursos para o controle financeiro pessoal:

- **Registro de Gastos**: Permite salvar rapidamente despesas especificando o valor, a categoria e, opcionalmente, o metodo de pagamento.
- **IDs Sequenciais**: Cada gasto registrado recebe um identificador sequencial (1, 2, 3...) exclusivo para o usuario, simplificando consultas e exclusoes.
- **Consulta Inteligente**: Listagem completa de despesas ou busca detalhada por ID, com suporte a navegacao interativa.
- **Navegacao por Botoes**: Botoes inline no Telegram que permitem avancar ou retroceder entre os registros cadastrados de forma fluida.
- **Exclusao Segura**: Opcao para deletar um gasto especifico ou apagar todos os registros, ambos exigindo confirmacao inline previa para evitar exclusoes acidentais.

---

## Comandos Disponiveis no Telegram

### 1. Registrar Gasto
```text
/gastei <valor> <categoria> [metodo]
```
- **Uso**: Registra uma nova despesa. O valor deve ser numerico (ex: 45.50), a categoria descreve o gasto e o metodo e opcional.
- **Exemplo**: `/gastei 45.50 supermercado debito`

### 2. Consultar Gastos
```text
/consulta                  # Lista todos os gastos cadastrados
/consulta <ID>             # Exibe detalhes de um gasto especifico
```
- **Recurso**: Ao consultar por ID, botoes inline de "Anterior" e "Proximo" sao exibidos no card para navegar sequencialmente.

### 3. Deletar Gasto
```text
/deletar <ID>              # Exclui o gasto correspondente ao ID informado
/deletartudo               # Remove todos os gastos cadastrados do usuario
```
- **Seguranca**: Ambos os comandos exibem uma caixa de confirmacao na tela com os botoes [Confirmar] e [Cancelar].

### 4. Suporte e Ajuda
```text
/help                      # Exibe a listagem e descricao de todos os comandos
/start                     # Apresenta a mensagem de boas-vindas do assistente
```

---

## Fluxo de Operacoes

### Registro de Gasto
1. O usuario digita o comando `/gastei 50.00 uber pix`.
2. O bot valida o formato e a corretude do valor numerico.
3. O banco de dados DynamoDB e consultado para encontrar o proximo ID sequencial livre para o usuario.
4. Os dados sao persistidos na tabela do DynamoDB.
5. O bot envia uma mensagem de confirmacao ao usuario.

### Consulta com Navegacao Inline
1. O usuario digita o comando `/consulta 3`.
2. O bot busca a despesa correspondente ao ID sequencial 3 do usuario logado.
3. O bot renderiza o card com as informacoes detalhadas.
4. O usuario interage clicando nos botoes inline de navegacao ("Anterior" ou "Proximo").
5. O bot atualiza o conteudo do card inline diretamente, sem poluir o chat com novas mensagens.

### Exclusao com Confirmacao
1. O usuario executa `/deletar 3`.
2. O bot exibe o resumo da despesa juntamente com os botoes de confirmacao.
3. O usuario clica em confirmar.
4. O registro e removido fisicamente do DynamoDB.
5. O bot confirma a exclusao e reorganiza os indices sequenciais de forma automatica na consulta seguinte do usuario.

---

## Roadmap Futuro (Proximas Implementacoes)

Recursos planejados para inclusao no bot Money Savior:
- Geracao de resumo mensal consolidado de gastos.
- Visualizacao de graficos estatisticos de consumo por categoria.
- Exportacao completa das despesas nos formatos CSV e PDF.
- Configuracao de limite maximo de gastos diarios.
- Alertas e notificacoes automaticas de limites excedidos (overspending).
- Suporte a integracao com outros bancos de dados.
- Modo de conta compartilhada para controle financeiro em grupo.
