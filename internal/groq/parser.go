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
