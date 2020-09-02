package requestStruct

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"regexp"
	connect "servientrega/connection"
	"strconv"
)

type Data struct {
	ShipmentObject *Shipment
}
type Shipment struct {
	XMLName xml.Name        `xml:"objEnvios" json:"-"`
	Object  *ShipmentObject `xml:"EnviosExterno" json:"Object"`
}
type ShipmentObject struct {
	XMLName                 xml.Name `xml:"EnviosExterno" json:"-"`
	NumGuia                 int      `xml:"Num_Guia" json:"NumGuia"`
	NumSobreporte           int      `xml:"Num_Sobreporte" json:"Num_Sobreporte"`
	DocRelacionado          string   `xml:"Doc_Relacionado" json:"Doc_Relacionado"`
	NumPiezas               int      `xml:"Num_Piezas" json:"Num_Piezas"`
	DesTipoTrayecto         int      `xml:"Des_TipoTrayecto" json:"Des_TipoTrayecto"`
	Ideproducto             int      `xml:"Ide_producto" json:"Ide_producto"`
	DesFormaPago            int      `xml:"Des_FormaPago" json:"Des_FormaPago"`
	DesMedioTransporte      int      `xml:"Des_MedioTransporte" json:"Des_MedioTransporte"`
	NumPesoTotal            int      `xml:"Num_PesoTotal" json:"Num_PesoTotal"`
	NumValorDeclaradoTotal  float32  `xml:"Num_ValorDeclaradoTotal" json:"Num_ValorDeclaradoTotal"`
	NumVolumenTotal         int      `xml:"Num_VolumenTotal" json:"Num_VolumenTotal"`
	NumBolsaSeguridad       int      `xml:"Num_BolsaSeguridad" json:"Num_BolsaSeguridad"`
	DesTipoDuracionTrayecto int      `xml:"Des_TipoDuracionTrayecto" json:"Des_TipoDuracionTrayecto"`
	DesTelefono             int      `xml:"Des_Telefono" json:"Des_Telefono"`
	DesCiudad               string   `xml:"Des_Ciudad" json:"Des_Ciudad"`
	DesDireccion            string   `xml:"Des_Direccion" json:"Des_Direccion"`
	NomContacto             string   `xml:"Nom_Contacto" json:"Nom_Contacto"`
	NumValorLiquidado       int      `xml:"Num_ValorLiquidado" json:"Num_ValorLiquidado"`
	DesDiceContener         string   `xml:"Des_DiceContener" json:"Des_DiceContener"`
	DesTipoGuia             int      `xml:"Des_TipoGuia" json:"Des_TipoGuia"`
	NumVlrSobreflete        int      `xml:"Num_VlrSobreflete" json:"Num_VlrSobreflete"`
	NumVlrflete             int      `xml:"Num_Vlrflete" json:"Num_Vlrflete"`
	NumDescuento            int      `xml:"Num_Descuento" json:"Num_Descuento"`
	NumPesoFacturado        int      `xml:"Num_PesoFacturado" json:"Num_PesoFacturado"`
	IdePaisOrigen           int      `xml:"idePaisOrigen" json:"idePaisOrigen"`
	IdePaisDestino          int      `xml:"idePaisDestino" json:"idePaisDestino"`
	NumAlto                 int      `xml:"Num_Alto" json:"Num_Alto"`
	NumAncho                int      `xml:"Num_Ancho" json:"Num_Ancho"`
	NumLargo                int      `xml:"Num_Largo" json:"Num_Largo"`
	DesDepartamentoDestino  string   `xml:"Des_DepartamentoDestino" json:"Des_DepartamentoDestino"`
	DesDepartamentoOrigen   string   `xml:"Des_DepartamentoOrigen" json:"Des_DepartamentoOrigen"`
	NomUnidadEmpaque        string   `xml:"Nom_UnidadEmpaque" json:"Nom_UnidadEmpaque"`
	DesUnidadLongitud       string   `xml:"Des_UnidadLongitud" json:"Des_UnidadLongitud"`
	DesUnidadPeso           string   `xml:"Des_UnidadPeso" json:"Des_UnidadPeso"`
}

func (d *Data) ConvertToJson(orderID int) ([]byte, string) {
	temp := ShipmentObject{}
	correo, _ := d.Prueba(&temp, orderID)
	temp2 := Shipment{Object: &temp}

	rawXMLData := `<objEnvios>
	<EnviosExterno>
	<Num_Guia></Num_Guia>
	<Num_Sobreporte></Num_Sobreporte>
	<Doc_Relacionado></Doc_Relacionado>
	<Num_Piezas></Num_Piezas>
	<Des_TipoTrayecto></Des_TipoTrayecto>
	<Ide_producto></Ide_producto>
	<Des_FormaPago></Des_FormaPago>
	<Des_MedioTransporte></Des_MedioTransporte>
	<Num_PesoTotal></Num_PesoTotal>
	<Num_ValorDeclaradoTotal></Num_ValorDeclaradoTotal>
	<Num_VolumenTotal></Num_VolumenTotal>
	<Num_BolsaSeguridad></Num_BolsaSeguridad>
	<Des_TipoDuracionTrayecto></Des_TipoDuracionTrayecto>
	<Des_Telefono></Des_Telefono>
	<Des_Ciudad></Des_Ciudad>
	<Des_Direccion></Des_Direccion>
	<Nom_Contacto></Nom_Contacto>
	<Num_ValorLiquidado></Num_ValorLiquidado>
	<Des_DiceContener></Des_DiceContener>
	<Des_TipoGuia></Des_TipoGuia>
	<Num_VlrSobreflete></Num_VlrSobreflete>
	<Num_Vlrflete></Num_Vlrflete>
	<Num_Descuento></Num_Descuento>
	<Num_PesoFacturado></Num_PesoFacturado>
	<idePaisOrigen></idePaisOrigen>
	<idePaisDestino></idePaisDestino>
	<Num_Alto></Num_Alto>
	<Num_Ancho></Num_Ancho>
	<Num_Largo></Num_Largo>
	<Des_DepartamentoDestino></Des_DepartamentoDestino>
	<Des_DepartamentoOrigen></Des_DepartamentoOrigen>
	<Nom_UnidadEmpaque></Nom_UnidadEmpaque>
	<Des_UnidadLongitud></Des_UnidadLongitud>
	<Des_UnidadPeso></Des_UnidadPeso>
	</objEnvios>
	</EnviosExterno>`
	rawXMLHeader := `<?xml version="1.0" encoding="UTF-8"?>
	<env:Envelope xmlns:env="http://www.w3.org/2003/05/soap-envelope"
	xmlns:ns1="http://tempuri.org/">
	<env:Header xmlns="http://tempuri.org/">
	<AuthHeader>
	<login>Luis1937</login>
	<pwd>MZR0zNqnI/KplFlYXiFk7m8/G/Iqxb3O</pwd>
	<Id_CodFacturacion>SER408</Id_CodFacturacion>
	<Nombre_Cargue>Colombia1</Nombre_Cargue>
	</AuthHeader>
	</env:Header>
	<env:Body>
	<CargueMasivoExterno xmlns="http://tempuri.org/">
	<envios>
	<CargueMasivoExternoDTO>`
	rawXMLFooter := `</CargueMasivoExternoDTO>
	</envios>
	<arrayGuias />
	</CargueMasivoExterno>
	</env:Body>
	</env:Envelope>`
	xml.Unmarshal([]byte(rawXMLData), &d.ShipmentObject)
	d.ShipmentObject = &temp2
	xmlData, err := xml.Marshal(&d.ShipmentObject)
	if err != nil {
		fmt.Println(err.Error())
	}
	regex := regexp.MustCompile("\\s+")
	rawXMLHeader = regex.ReplaceAllString(rawXMLHeader, " ")
	rawXMLFooter = regex.ReplaceAllString(rawXMLFooter, " ")
	return []byte(rawXMLHeader + string(xmlData) + rawXMLFooter), correo

}

func (d *Data) Prueba(object *ShipmentObject, orderID int) (string, error) {
	conn := new(connect.Connection)
	var correoPersona string
	_, err := conn.Connect()
	defer conn.Disconnect()
	if err != nil {
		return "", err
	}
	ctx := conn.GetContext()
	db := conn.GetDBConnection()
	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return "", err
	}
	tsql := fmt.Sprintf(`SELECT order_id,
	num_items_sold,
	net_total,
	billing_cedula,
	billing_first_name,
	billing_last_name,
	billing_city,
	billing_state,
	billing_phone,
	billing_email,
	billing_address,
	idProducto,
	tipoDuracionTrayecto,
	medioTransporte,
	tipoTrayecto,
	paisDestino,
	paisOrigen,
	departamentoOrigen,
	unidadEmpaque,
	largo,
	ancho,
	alto,
	peso,
	valorDeclarado,
	unidadLongitud,
	unidadPeso 
	FROM VistaDespachoServ WHERE order_id = @orderID`)
	rows, err := db.QueryContext(ctx, tsql, sql.Named("orderID", orderID))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		var id, nroItems, totalVendido, cedula, idProducto, tipoDuracionTrayecto, medioTransporte, tipoTrayecto, paisOrigen, paisDestino, largo, ancho, alto, peso, valorDeclarado int
		var nombre, apellido, ciudad, departamentoDestino, telefono, direccion, departamentoOrigen, unidadEmpaque, unidadLongitud, unidadPeso string

		// Get values from row.
		err := rows.Scan(&id, &nroItems, &totalVendido, &cedula, &nombre, &apellido, &ciudad, &departamentoDestino, &telefono, &correoPersona, &direccion, &idProducto, &tipoDuracionTrayecto, &medioTransporte, &tipoTrayecto, &paisDestino, &paisOrigen, &departamentoOrigen, &unidadEmpaque, &largo, &ancho, &alto, &peso, &valorDeclarado, &unidadLongitud, &unidadPeso)
		if err != nil {
			return "", err
		}
		object.DocRelacionado = strconv.Itoa(id)
		object.NumPiezas = nroItems
		object.DesTipoTrayecto = tipoTrayecto
		object.Ideproducto = idProducto
		object.DesFormaPago = 2
		object.DesMedioTransporte = medioTransporte
		object.NumPesoTotal = peso
		object.NumValorDeclaradoTotal = float32(valorDeclarado)
		object.DesTipoDuracionTrayecto = tipoDuracionTrayecto
		i, err := strconv.Atoi(telefono)
		if err == nil {
			object.DesTelefono = i
		}
		object.DesCiudad = ciudad
		object.DesDireccion = direccion
		object.NomContacto = nombre + " " + apellido
		object.DesDiceContener = "CALZADO"
		object.DesTipoGuia = 44
		object.IdePaisOrigen = paisOrigen
		object.IdePaisDestino = paisDestino
		object.NumAlto = alto
		object.NumAncho = ancho
		object.NumLargo = largo
		object.NumAlto = alto
		object.DesDepartamentoDestino = departamentoDestino
		object.DesDepartamentoOrigen = departamentoOrigen
		object.NomUnidadEmpaque = unidadEmpaque
		object.DesUnidadLongitud = unidadLongitud
		object.DesUnidadPeso = unidadPeso

	}
	return correoPersona, nil
}
