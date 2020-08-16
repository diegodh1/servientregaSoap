package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	requestStruct "servientrega/requestStruct"
	responseRequest "servientrega/responseRequest"
)

func main() {
	var data requestStruct.Data
	httpMethod := "POST"
	payload := data.ConvertToJson()
	url := "http://web.servientrega.com:8081/GeneracionGuias.asmx"
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}
	req.Header.Set("Content-type", "text/xml")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}
	response := new(responseRequest.CargueMasivoExternoResponse)
	err = xml.NewDecoder(res.Body).Decode(response)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}
	fmt.Println(response.Body.Guia)
}
