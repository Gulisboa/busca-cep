package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao tentar fazer a requisição %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler as respostas%v\n", err)
		}
		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta%v\n", err)
		}
		file, err := os.Create("arquivo.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo%v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Cidade: %s, CEP: %s, Rua: %s", data.Localidade, data.Cep, data.Logradouro))
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo%v", err)
		}
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Printf("Logradouro: %s", data.Logradouro)
	}
}
