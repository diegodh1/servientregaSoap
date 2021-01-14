package main

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	connect "servientrega/connection"
	requestStruct "servientrega/requestStruct"
	responseRequest "servientrega/responseRequest"
	"strconv"
	"sync"

	gomail "gopkg.in/gomail.v2"
)

//APIConsume func
type APIConsume struct {
	mu         sync.Mutex
	data       requestStruct.Data
	httpMethod string
	url        string
}

func (a *APIConsume) enviarInfo(pedido int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	a.httpMethod = "POST"
	payload, correo := a.data.ConvertToJson(pedido)
	a.url = "http://web.servientrega.com:8081/GeneracionGuias.asmx"
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
	from := "noreply-ventas@calzadoromulo.com.co"
	pass := "Temporal.2021@"
	to := correo
	v := strconv.Itoa(guia)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Número de Guía Rómulo")
	m.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	<style>
		body {
		   font-family: "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif; 
		   font-weight: 300;
		}
	</style>
	<html>
		<body>
		<p>Confirmamos su compra con el siguiente n&uacute;mero de gu&iacute;a:</p>
		<p><span style="color: #ff0000;"><strong>%s</strong></span></p>
		<p>!Gracias por ser parte de la familia de Calzado Romulo!</p>
		<p><a href="https://www.servientrega.com/wps/portal/Colombia/transacciones-personas/rastreo-envios"><img src="https://drive.google.com/uc?export=download&amp;id=1e6Z2HfMa3QyDnAdxg48irwdOGoAWfAjQ" alt="guia" width="756" height="510" /></a></p>
		<p>Quedo atenta,</p>
		<p><img src="https://drive.google.com/uc?export=download&amp;id=1K7BZl2lR6YCeGuCIfB9SV4kdSwqsxQS2" alt="firma" width="589" height="474" /></p>
		</body>
	</html>`, v))

	// Send the email to Bob
	d := gomail.NewPlainDialer("smtpout.secureserver.net", 80, from, pass)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}

//ObtenerPedidosPendietes metodo para tener los pedidos pendientes por enviar
func ObtenerPedidosPendietes() {
	var wbEnvio APIConsume
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
	//wg.Add(len(listaPendientes))
	wg.Add(1)
	//or _, v := range listaPendientes {
	go wbEnvio.enviarInfo(21237, &wg)
	//}
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
