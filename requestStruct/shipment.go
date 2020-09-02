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
	XMLName xml.Name        `xml:"tem:objEnvios" json:"-"`
	Object  *ShipmentObject `xml:"tem:EnviosExterno" json:"Object"`
}
type ShipmentObject struct {
	XMLName                     xml.Name `xml:"tem:EnviosExterno" json:"-"`
	NumGuia                     int      `xml:"tem:Num_Guia" json:"NumGuia"`
	NumSobreporte               int      `xml:"tem:Num_Sobreporte" json:"Num_Sobreporte"`
	NumSobreCajaporte           int      `xml:"tem:Num_SobreCajaPorte" json:"Num_SobreCajaPorte"`
	FecTiempoEntrega            int      `xml:"tem:Fec_TiempoEntrega" json:"Fec_TiempoEntrega"`
	DesTipoTrayecto             int      `xml:"tem:Des_TipoTrayecto" json:"Des_TipoTrayecto"`
	IdeCodFacturacion           string   `xml:"tem:Ide_CodFacturacion" json:"Ide_CodFacturacion"`
	NumPiezas                   int      `xml:"tem:Num_Piezas" json:"Num_Piezas"`
	DesFormaPago                int      `xml:"tem:Des_FormaPago" json:"Des_FormaPago"`
	DesMedioTransporte          int      `xml:"tem:Des_MedioTransporte" json:"Des_MedioTransporte"`
	DesTipoDuracionTrayecto     int      `xml:"tem:Des_TipoDuracionTrayecto" json:"Des_TipoDuracionTrayecto"`
	NomTipoTrayecto             int      `xml:"tem:Nom_TipoTrayecto" json:"Nom_TipoTrayecto"`
	NumAlto                     int      `xml:"tem:Num_Alto" json:"Num_Alto"`
	NumAncho                    int      `xml:"tem:Num_Ancho" json:"Num_Ancho"`
	NumLargo                    int      `xml:"tem:Num_Largo" json:"Num_Largo"`
	NumPesoTotal                int      `xml:"tem:Num_PesoTotal" json:"Num_PesoTotal"`
	DesUnidadLongitud           string   `xml:"tem:Des_UnidadLongitud" json:"Des_UnidadLongitud"`
	DesUnidadPeso               string   `xml:"tem:Des_UnidadPeso" json:"Des_UnidadPeso"`
	NomUnidadEmpaque            string   `xml:"tem:Nom_UnidadEmpaque" json:"Nom_UnidadEmpaque"`
	GenCajaporte                bool     `xml:"tem:Gen_Cajaporte" json:"Gen_Cajaporte"`
	GenSobreporte               bool     `xml:"tem:Gen_Sobreporte" json:"Gen_Sobreporte"`
	DesDiceContenerSobre        string   `xml:"tem:Des_DiceContenerSobre" json:"Des_DiceContenerSobre"`
	DocRelacionado              string   `xml:"tem:Doc_Relacionado" json:"Doc_Relacionado"`
	Ideproducto                 int      `xml:"tem:Ide_Producto" json:"Ide_Producto"`
	IdeDestinatarios            string   `xml:"tem:Ide_Destinatarios" json:"Ide_Destinatarios"`
	IdeManifiesto               string   `xml:"tem:Ide_Manifiesto" json:"Ide_Manifiesto"`
	NumBolsaSeguridad           int      `xml:"tem:Num_BolsaSeguridad" json:"Num_BolsaSeguridad"`
	NumPrecinto                 int      `xml:"tem:Num_Precinto" json:"Num_Precinto"`
	NumVolumenTotal             int      `xml:"tem:Num_VolumenTotal" json:"Num_VolumenTotal"`
	DesDireccionRecogida        string   `xml:"tem:Des_DireccionRecogida" json:"Des_DireccionRecogida"`
	DesTelefonoRecogida         string   `xml:"tem:Des_TelefonoRecogida" json:"Des_TelefonoRecogida"`
	DesCiudadRecogida           string   `xml:"tem:Des_CiudadRecogida" json:"Des_CiudadRecogida"`
	NumPesoFacturado            int      `xml:"tem:Num_PesoFacturado" json:"Num_PesoFacturado"`
	DesTipoGuia                 int      `xml:"tem:Des_TipoGuia" json:"Des_TipoGuia"`
	IDArchivoCargar             string   `xml:"tem:Id_ArchivoCargar" json:"Id_ArchivoCargar"`
	DesCiudadOrigen             int      `xml:"tem:Des_CiudadOrigen" json:"Des_CiudadOrigen"`
	NumValorDeclaradoTotal      float32  `xml:"tem:Num_ValorDeclaradoTotal" json:"Num_ValorDeclaradoTotal"`
	NumValorLiquidado           int      `xml:"tem:Num_ValorLiquidado" json:"Num_ValorLiquidado"`
	NumVlrSobreflete            int      `xml:"tem:Num_VlrSobreflete" json:"Num_VlrSobreflete"`
	NumVlrflete                 int      `xml:"tem:Num_Vlrflete" json:"Num_Vlrflete"`
	NumDescuento                int      `xml:"tem:Num_Descuento" json:"Num_Descuento"`
	NumValorDeclaradoSobreTotal int      `xml:"tem:Num_ValorDeclaradoSobreTotal" json:"Num_ValorDeclaradoSobreTotal"`
	DesTelefono                 int      `xml:"tem:Des_Telefono" json:"Des_Telefono"`
	DesCiudad                   string   `xml:"tem:Des_Ciudad" json:"Des_Ciudad"`
	DesDepartamentoDestino      string   `xml:"tem:Des_DepartamentoDestino" json:"Des_DepartamentoDestino"`
	DesDireccion                string   `xml:"tem:Des_Direccion" json:"Des_Direccion"`
	NomContacto                 string   `xml:"tem:Nom_Contacto" json:"Nom_Contacto"`
	DesDiceContener             string   `xml:"tem:Des_DiceContener" json:"Des_DiceContener"`
	IdeNumIdentificDest         int      `xml:"tem:Ide_Num_Identific_Dest" json:"Ide_Num_Identific_Dest"`
	NumCelular                  int      `xml:"tem:Num_Celular" json:"Num_Celular"`
	DesCorreoElectronico        string   `xml:"tem:Des_CorreoElectronico" json:"Des_CorreoElectronico"`
	DesCiudadRemitente          string   `xml:"tem:Des_CiudadRemitente" json:"Des_CiudadRemitente"`
	DesDireccionRemitente       string   `xml:"tem:Des_DireccionRemitente" json:"Des_DireccionRemitente"`
	DesDepartamentoOrigen       string   `xml:"tem:Des_DepartamentoOrigen" json:"Des_DepartamentoOrigen"`
	NumTelefonoRemitente        string   `xml:"tem:Num_TelefonoRemitente" json:"Num_TelefonoRemitente"`
	NumIdentiRemitente          string   `xml:"tem:Num_IdentiRemitente" json:"Num_IdentiRemitente"`
	NomRemitente                string   `xml:"tem:Nom_Remitente" json:"Nom_Remitente"`
	EstCanalMayorista           bool     `xml:"tem:Est_CanalMayorista" json:"Est_CanalMayorista"`
	NomRemitenteCanal           string   `xml:"tem:Nom_RemitenteCanal" json:"Nom_RemitenteCanal"`
}

func (d *Data) ConvertToJson(orderID int) ([]byte, string) {
	temp := ShipmentObject{}
	correo, _ := d.Prueba(&temp, orderID)
	temp2 := Shipment{Object: &temp}

	rawXMLData := `<tem:objEnvios>
	<!--Zero or more repetitions:-->
	<tem:EnviosExterno>
	   <!--PRODUCTO DOCUMENTO VALORES DESDE 1 K Y HASTA 50 K:-->
	   <!--CARACTETISTICAS DEL ENVIO:-->
	   <tem:Num_Guia>0</tem:Num_Guia>
	   <tem:Num_Sobreporte>0</tem:Num_Sobreporte>
	   <tem:Num_SobreCajaPorte>0</tem:Num_SobreCajaPorte>
	   <tem:Fec_TiempoEntrega>1</tem:Fec_TiempoEntrega>
	   <tem:Des_TipoTrayecto></tem:Des_TipoTrayecto>
	   <tem:Ide_CodFacturacion>SER408</tem:Ide_CodFacturacion>
	   <tem:Num_Piezas></tem:Num_Piezas>
	   <!--CUANDO SEA MERCANCIA INDUSTRIAL (Producto 6) COLOCAR LA CANTIDAD DE PIEZAS FISICAS DEL ENVIO:-->
	   <tem:Des_FormaPago></tem:Des_FormaPago>
	   <tem:Des_MedioTransporte></tem:Des_MedioTransporte>
	   <tem:Des_TipoDuracionTrayecto></tem:Des_TipoDuracionTrayecto>
	   <tem:Nom_TipoTrayecto>1</tem:Nom_TipoTrayecto>
	   <tem:Num_Alto></tem:Num_Alto>
	   <tem:Id_ArchivoCargar></tem:Id_ArchivoCargar>
	   <tem:Num_Ancho></tem:Num_Ancho>
	   <tem:Num_Largo></tem:Num_Largo>
	   <tem:Num_PesoTotal></tem:Num_PesoTotal>
	   <tem:Des_UnidadLongitud></tem:Des_UnidadLongitud>
	   <tem:Des_UnidadPeso></tem:Des_UnidadPeso>
	   <tem:Nom_UnidadEmpaque></tem:Nom_UnidadEmpaque>
	   <!--ingresar el nombre de la unidad de empaque que se encuentre registrada en sisclinet:-->
	   <tem:Gen_Cajaporte>false</tem:Gen_Cajaporte>
	   <tem:Gen_Sobreporte>false</tem:Gen_Sobreporte>
	   <tem:Des_DiceContenerSobre>?</tem:Des_DiceContenerSobre>
	   <!--dice contener solon para producto sobreporte:-->
	   <!--CAMPOS OPCIONALES CON CAIDA IMPRESA EN LA GUIA:-->
	   <tem:Doc_Relacionado></tem:Doc_Relacionado>
	   <!--campo remision en la guia:-->
	   <!--TIPO DE PRODUCTO A UTILIZAR DE SERVIENTREGA:-->
	   <tem:Ide_Producto>2</tem:Ide_Producto>
	   <!--codigos del tipo de producto negociado con Servientrega:-->
	   <!--valor a recaudar, solo aplica para logistica para cobro:-->
	   <!--OTRA INFORMACION, DEJAR SIEMPRE FIJA:-->
	   <tem:Ide_Destinatarios>00000000-0000-0000-0000-000000000000</tem:Ide_Destinatarios>
	   <tem:Ide_Manifiesto>00000000-0000-0000-0000-000000000000</tem:Ide_Manifiesto>
	   <tem:Num_BolsaSeguridad>0</tem:Num_BolsaSeguridad>
	   <tem:Num_Precinto>0</tem:Num_Precinto>
	   <tem:Num_VolumenTotal>0</tem:Num_VolumenTotal>
	   <tem:Des_DireccionRecogida />
	   <!--no aplica, enviar 0:-->
	   <tem:Des_TelefonoRecogida />
	   <!--no aplica, enviar 0:-->
	   <tem:Des_CiudadRecogida />
	   <!--no aplica, enviar 0:-->
	   <tem:Num_PesoFacturado>0</tem:Num_PesoFacturado>
	   <tem:Des_TipoGuia>2</tem:Des_TipoGuia>
	   <tem:Id_ArchivoCargar />
	   <tem:Des_CiudadOrigen>0</tem:Des_CiudadOrigen>
	   <!--dejar en 0:-->
	   <!--VALORES:-->
	   <tem:Num_ValorDeclaradoTotal></tem:Num_ValorDeclaradoTotal>
	   <!--valor declarado del envio total:-->
	   <tem:Num_ValorLiquidado>0</tem:Num_ValorLiquidado>
	   <tem:Num_VlrSobreflete>0</tem:Num_VlrSobreflete>
	   <tem:Num_VlrFlete>0</tem:Num_VlrFlete>
	   <tem:Num_Descuento>0</tem:Num_Descuento>
	   <tem:Num_ValorDeclaradoSobreTotal>0</tem:Num_ValorDeclaradoSobreTotal>
	   <!--INFORMACION DEL DESTINATARIO:-->
	   <tem:Des_Telefono></tem:Des_Telefono>
	   <!--Telefono Destinatario-->
	   <tem:Des_Ciudad></tem:Des_Ciudad>
	   <!--Ciudad de Destino:-->
	   <tem:Des_DepartamentoDestino></tem:Des_DepartamentoDestino>
	   <!--nombre o codigo DANE de departamento de destino:-->
	   <tem:Des_Direccion></tem:Des_Direccion>
	   <!--Direccion de Destino:-->
	   <tem:Nom_Contacto></tem:Nom_Contacto>
	   <!--Nombre de Destinatario:-->
	   <tem:Des_DiceContener></tem:Des_DiceContener>
	   <!--conteneido del envio hasta 20 caracteres:-->
	   <tem:Ide_Num_Identific_Dest></tem:Ide_Num_Identific_Dest>
	   <!--identificacion de destinario :-->
	   <tem:Num_Celular />
	   <!--numero de celular de destinatario:-->
	   <tem:Des_CorreoElectronico></tem:Des_CorreoElectronico>
	   <!--INFORMACION DEL REMITENTE:-->
	   <tem:Des_CiudadRemitente />
	   <!--ciudad de remitente:-->
	   <tem:Des_DireccionRemitente />
	   <!--direccion del remitente:-->
	   <tem:Des_DepartamentoOrigen />
	   <!--nombre o codigo DANE de departamento de Origen:-->
	   <tem:Num_TelefonoRemitente />
	   <!--telefono del remitente:-->
	   <tem:Num_IdentiRemitente />
	   <!--identificacion del remitente:-->
	   <tem:Nom_Remitente />
	   <!--no aplica, dejar vacio:-->
	   --------------------
	   <!--MAYORISTA , NO APICA PARA CLIENTES:-->
	   <tem:Est_CanalMayorista>false</tem:Est_CanalMayorista>
	   <tem:Nom_RemitenteCanal />
	</tem:EnviosExterno>
 </tem:objEnvios>`
	rawXMLHeader := `<?xml version="1.0" encoding="UTF-8"?>
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:tem="http://tempuri.org/">
	   <soap:Header>
		  <tem:AuthHeader>
			 <!--Optional:-->
			 <tem:login>Luis1937</tem:login>
			 <!--Optional:-->
			 <tem:pwd>MZR0zNqnI/KplFlYXiFk7m8/G/Iqxb3O</tem:pwd>
			 <!--Optional:-->
			 <tem:Id_CodFacturacion>SER408</tem:Id_CodFacturacion>
			 <!--Optional:-->
			 <tem:Nombre_Cargue>AJA_GROUP</tem:Nombre_Cargue>
		  </tem:AuthHeader>
	   </soap:Header>
	   <soap:Body>
		  <tem:CargueMasivoExterno>
			 <tem:envios>
				<tem:CargueMasivoExternoDTO>`
	rawXMLFooter := `</tem:CargueMasivoExternoDTO>
	</tem:envios>
	<!--Optional:-->
	<tem:arrayGuias>
	   <!--Zero or more repetitions:-->
	   <tem:string>?</tem:string>
	</tem:arrayGuias>
 </tem:CargueMasivoExterno>
</soap:Body>
</soap:Envelope>`
	xml.Unmarshal([]byte(rawXMLData), &d.ShipmentObject)
	d.ShipmentObject = &temp2
	xmlData, err := xml.Marshal(&d.ShipmentObject)
	if err != nil {
		fmt.Println(err.Error())
	}
	regex := regexp.MustCompile("\\s+")
	rawXMLHeader = regex.ReplaceAllString(rawXMLHeader, " ")
	rawXMLFooter = regex.ReplaceAllString(rawXMLFooter, " ")
	fmt.Println(rawXMLHeader + string(xmlData) + rawXMLFooter)
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
		object.DesCiudad = "73449000"
		object.DesDireccion = direccion
		object.NomContacto = nombre + " " + apellido
		object.DesDiceContener = "CALZADO"
		object.DesTipoGuia = 2
		object.NumAlto = alto
		object.NumAncho = ancho
		object.NumLargo = largo
		object.NumAlto = alto
		object.DesDepartamentoDestino = "47268000"
		object.NomUnidadEmpaque = "GENERICO"
		object.DesUnidadLongitud = unidadLongitud
		object.DesUnidadPeso = unidadPeso
		object.IdeNumIdentificDest = cedula
		object.DesCorreoElectronico = correoPersona
		object.NumCelular = i
		object.IdeCodFacturacion = "SER408"
		object.NomTipoTrayecto = 1
		object.GenCajaporte = false
		object.GenSobreporte = false
		object.DesDiceContenerSobre = "?"
		object.IdeDestinatarios = "00000000-0000-0000-0000-000000000000"
		object.IdeManifiesto = "00000000-0000-0000-0000-000000000000"
		object.FecTiempoEntrega = 1

	}
	return correoPersona, nil
}
