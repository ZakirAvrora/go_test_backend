package app

import (
	"ZakirAvrora/go_test_backend/service2/internals/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type Generator interface {
	Generate(n int) string
}

type app struct {
	gen Generator
}

func New(gen Generator) *app {
	return &app{gen: gen}
}

func (a *app) Handle(w http.ResponseWriter, r *http.Request) {
	//str := a.gen.Generate(14)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	msg, err := json.Marshal(model.MsgModel{Salt: a.gen.Generate(14)})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(msg))
}
