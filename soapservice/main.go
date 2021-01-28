package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"servientrega/connection"
	requestStruct "servientrega/requeststruct"
	responseRequest "servientrega/responserequest"
	"strconv"

	gomail "gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

//APIConsume func
type APIConsume struct {
	data       requestStruct.Data
	httpMethod string
	url        string
}

func (a *APIConsume) enviarInfo(pedido int, despacho *requestStruct.DespachoServ, db *gorm.DB) {
	a.httpMethod = "POST"
	payload, correo := a.data.ConvertToJSON(pedido, despacho)
	a.url = "http://web.servientrega.com:8081/GeneracionGuias.asmx"
	req, err := http.NewRequest(a.httpMethod, a.url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
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
	}
	response := new(responseRequest.CargueMasivoExternoResponse)
	err = xml.NewDecoder(res.Body).Decode(response)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
	}
	fmt.Println(response.Body.EnvioObj.NumGuia)
	insertarGuia(response.Body.EnvioObj.NumGuia, pedido, db)
	enviarCorreo(correo, response.Body.EnvioObj.NumGuia)
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
		</body>
	</html>`, v))

	// Send the email to Bob
	d := gomail.NewPlainDialer("smtpout.secureserver.net", 80, from, pass)
	d.DialAndSend(m)
}

//ObtenerPedidosPendietes metodo para tener los pedidos pendientes por enviar
func ObtenerPedidosPendietes(db *gorm.DB) {
	var wbEnvio APIConsume
	listaPendientes := []requestStruct.DespachoServ{}
	db.Find(&listaPendientes)
	for _, v := range listaPendientes {
		wbEnvio.enviarInfo(v.OrderID, &v, db)
	}
}

//RegistroGuia struct
type RegistroGuia struct {
	GuiaNro int `gorm:"type:int;"`
	OrderID int `gorm:"type:int;"`
}

func insertarGuia(guia int, order int, db *gorm.DB) {
	data := RegistroGuia{
		GuiaNro: guia,
		OrderID: order,
	}
	db.Create(&data)
	db.Exec("UPDATE despacho SET despacho_estado = ? WHERE order_id = ?", true, order)
}
func insertarError(order int, errProduce error, db *sql.DB, ctx *context.Context) {
	tsql := fmt.Sprintf("INSERT INTO RegistroError(order_id,error) VALUES(@Order,@Error)")
	db.ExecContext(*ctx, tsql, sql.Named("Order", order), sql.Named("Error", errProduce.Error()))
}

func main() {
	var connection = new(connection.Config)
	connection.Init()
	db, err := connection.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	ObtenerPedidosPendietes(db)
}
