package main

import (
	"context"
	"fmt"
	"hugornda/picapuento/cmd/picapuento/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Action string `json:"action"`
}

type Picagem struct {
	codigo          string
	senha           string
	tipo_movimento  string
	pontocontrolo   string
	pontocontroloid string
	empresa_id      string
	empresa_url     string
	token           string
	modulo_pausas   string
}

const (
	CHECKIN           = "checkin"
	CHECKOUT          = "checkout"
	VALIDATE_USER_URL = "https://picaponto.pt/ponto/ajax_valdidate"
	CHANGE_STATUS_URL = "https://picaponto.pt/ponto/ajax_add"
	CONTENT_TYPE      = "application/x-www-form-urlencoded; charset=UTF-8"
	CHECKOUT_MOVEMENT = "Saída"
	CHECKIN_MOVEMENT  = "Entrada"
	USERID            = "PICA_PONTO_USER_ID"
	USER_PASSWORD     = "PICA_PONTO_USER_PASSWORD"
)

var c http.Client = http.Client{Timeout: time.Duration(1) * time.Minute}

var picagem Picagem = Picagem{
	codigo:          os.Getenv(USERID),
	senha:           os.Getenv(USER_PASSWORD),
	pontocontrolo:   "false",
	pontocontroloid: "0",
	empresa_id:      "1228",
}

func Picaponto(ctx context.Context, event *Request) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	message := fmt.Sprintf("Action pretended is %s!", event.Action)
	token, err := getPicaPontoToken()
	if err != nil {
		fmt.Printf("error: %e", err)
		return nil, err
	}
	switch event.Action {
	case CHECKIN:
		message = checkIn(token)
	case CHECKOUT:
		message = checkOut(token)
	default:
		panic("Not a valid action")
	}

	return &message, nil
}

func createRequestReader(token, tipoMovimento string) io.Reader {
	body := url.Values{}
	body.Set("codigo", picagem.codigo)
	body.Set("senha", picagem.senha)
	body.Set("tipo_movimento", tipoMovimento)
	body.Set("pontocontrolo", picagem.pontocontrolo)
	body.Set("pontocontroloid", picagem.pontocontroloid)
	body.Set("empresa_id", picagem.empresa_id)
	body.Set("empresa_url", picagem.empresa_url)
	body.Set("token", token)
	body.Set("modulo_pausas", picagem.modulo_pausas)
	return strings.NewReader(body.Encode())

}

func checkIn(token string) string {
	return ChangeStatus(token, CHECKIN_MOVEMENT)
}

func checkOut(token string) string {
	return ChangeStatus(token, CHECKOUT_MOVEMENT)
}

func ChangeStatus(token, tipoMovimento string) string {
	reader := createRequestReader(token, tipoMovimento)
	res, err := c.Post(CHANGE_STATUS_URL, CONTENT_TYPE, reader)

	if err != nil {
		return "Error"
	}
	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	return string(response)

}

func getPicaPontoToken() (string, error) {
	resp, err := c.Get("https://picaponto.pt/ponto/1228/Crossjoin")
	if err != nil {
		fmt.Printf("Error %s", err)
		return "", fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return utils.ExtractToken(string(body))
}

func main() {
	utils.DisplayArt()
	lambda.Start(Picaponto)
}
