package main

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	connect "servientrega/connection"
	requestStruct "servientrega/requestStruct"
	responseRequest "servientrega/responseRequest"
	"strconv"
	"sync"
)

type ApiConsume struct {
	mu         sync.Mutex
	data       requestStruct.Data
	httpMethod string
	url        string
}

func (a *ApiConsume) enviarInfo(pedido int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	a.httpMethod = "POST"
	payload, correo := a.data.ConvertToJson(pedido)
	a.url = "http://190.131.194.159:8059/GeneracionGuias.asmx"
	req, err := http.NewRequest(a.httpMethod, a.url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		insertarError(-1, err)
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
		insertarError(-1, err)
	}
	response := new(responseRequest.CargueMasivoExternoResponse)
	err = xml.NewDecoder(res.Body).Decode(response)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		insertarError(-1, err)
	}
	insertarGuia(response.Body.EnvioObj.NumGuia, pedido)
	enviarCorreo(correo, response.Body.EnvioObj.NumGuia)
	a.mu.Unlock()
}
func enviarCorreo(correo string, guia int) {
	fmt.Println(correo)
	from := "diegodiazh1994@gmail.com"
	pass := "cristiano1994"
	to := correo
	v := strconv.Itoa(guia)
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Pedido generado Rómulo\n\n" +
		"Su número de guía es: " + v + "\n\n" +
		"Muchas gracias por su compra."

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
func ObtenerPedidosPendietes() {
	var wbEnvio ApiConsume
	var wg sync.WaitGroup
	conn := new(connect.Connection)
	listaPendientes := []int{}
	_, err := conn.Connect()
	defer conn.Disconnect()
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx := conn.GetContext()
	db := conn.GetDBConnection()
	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	tsql := fmt.Sprintf("SELECT order_id FROM VistaDespachoServ")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		switch {
		case err == nil:
			listaPendientes = append(listaPendientes, id)
		case err != nil:
			fmt.Println(err.Error())
		default:
			fmt.Println("desconocido")
		}
	}
	wg.Add(len(listaPendientes))
	for _, v := range listaPendientes {
		go wbEnvio.enviarInfo(v, &wg)
	}
	wg.Wait()
}

func insertarGuia(guia int, order int) (int64, error) {
	conn := new(connect.Connection)
	_, err := conn.Connect()
	defer conn.Disconnect()
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx := conn.GetContext()
	db := conn.GetDBConnection()
	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	tsql := fmt.Sprintf("INSERT INTO RegistroGuia(guia_nro,order_id) VALUES(@Guia,@Order)")
	result, err := db.ExecContext(ctx, tsql, sql.Named("Guia", guia), sql.Named("Order", order))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
func insertarError(order int, errProduce error) {
	conn := new(connect.Connection)
	_, err := conn.Connect()
	defer conn.Disconnect()
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx := conn.GetContext()
	db := conn.GetDBConnection()
	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	tsql := fmt.Sprintf("INSERT INTO RegistroError(order_id,error) VALUES(@Order,@Error)")
	result, err := db.ExecContext(ctx, tsql, sql.Named("Order", order), sql.Named("Error", errProduce.Error()))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func main() {
	ObtenerPedidosPendietes()
}
