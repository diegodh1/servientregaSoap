package requestStruct

import (
	"encoding/xml"
	"regexp"
)

type Data struct {
	ShipmentObject *Shipment
}
type Shipment struct {
	XMLName xml.Name        `xml:"ns1:objEnvios" json:"-"`
	Object  *ShipmentObject `xml:"ns1:EnviosExterno" json:"Object"`
}
type ShipmentObject struct {
	XMLName                 xml.Name `xml:"ns1:EnviosExterno" json:"-"`
	NumGuia                 int      `xml:"ns1:Num_Guia" json:"NumGuia"`
	NumSobreporte           int      `xml:"ns1:Num_Sobreporte" json:"Num_Sobreporte"`
	DocRelacionado          string   `xml:"ns1:Doc_Relacionado" json:"Doc_Relacionado"`
	NumPiezas               int      `xml:"ns1:Num_Piezas" json:"Num_Piezas"`
	DesTipoTrayecto         int      `xml:"ns1:Des_TipoTrayecto" json:"Des_TipoTrayecto"`
	Ideproducto             int      `xml:"ns1:Ide_producto" json:"Ide_producto"`
	DesFormaPago            int      `xml:"ns1:Des_FormaPago" json:"Des_FormaPago"`
	DesMedioTransporte      int      `xml:"ns1:Des_MedioTransporte" json:"Des_MedioTransporte"`
	NumPesoTotal            int      `xml:"ns1:Num_PesoTotal" json:"Num_PesoTotal"`
	NumValorDeclaradoTotal  float32  `xml:"ns1:Num_ValorDeclaradoTotal" json:"Num_ValorDeclaradoTotal"`
	NumVolumenTotal         int      `xml:"ns1:Num_VolumenTotal" json:"Num_VolumenTotal"`
	NumBolsaSeguridad       int      `xml:"ns1:Num_BolsaSeguridad" json:"Num_BolsaSeguridad"`
	DesTipoDuracionTrayecto int      `xml:"ns1:Des_TipoDuracionTrayecto" json:"Des_TipoDuracionTrayecto"`
	DesTelefono             int      `xml:"ns1:Des_Telefono" json:"Des_Telefono"`
	DesCiudad               string   `xml:"ns1:Des_Ciudad" json:"Des_Ciudad"`
	DesDireccion            string   `xml:"ns1:Des_Direccion" json:"Des_Direccion"`
	NomContacto             string   `xml:"ns1:Nom_Contacto" json:"Nom_Contacto"`
	NumValorLiquidado       int      `xml:"ns1:Num_ValorLiquidado" json:"Num_ValorLiquidado"`
	DesDiceContener         string   `xml:"ns1:Des_DiceContener" json:"Des_DiceContener"`
	DesTipoGuia             int      `xml:"ns1:Des_TipoGuia" json:"Des_TipoGuia"`
	NumVlrSobreflete        int      `xml:"ns1:Num_VlrSobreflete" json:"Num_VlrSobreflete"`
	NumVlrflete             int      `xml:"ns1:Num_Vlrflete" json:"Num_Vlrflete"`
	NumDescuento            int      `xml:"ns1:Num_Descuento" json:"Num_Descuento"`
	NumPesoFacturado        int      `xml:"ns1:Num_PesoFacturado" json:"Num_PesoFacturado"`
	IdePaisOrigen           int      `xml:"ns1:idePaisOrigen" json:"idePaisOrigen"`
	IdePaisDestino          int      `xml:"ns1:idePaisDestino" json:"idePaisDestino"`
	NumAlto                 int      `xml:"ns1:Num_Alto" json:"Num_Alto"`
	NumAncho                int      `xml:"ns1:Num_Ancho" json:"Num_Ancho"`
	NumLargo                int      `xml:"ns1:Num_Largo" json:"Num_Largo"`
	DesDepartamentoDestino  string   `xml:"ns1:Des_DepartamentoDestino" json:"Des_DepartamentoDestino"`
	DesDepartamentoOrigen   string   `xml:"ns1:Des_DepartamentoOrigen" json:"Des_DepartamentoOrigen"`
	NomUnidadEmpaque        string   `xml:"ns1:Nom_UnidadEmpaque" json:"Nom_UnidadEmpaque"`
	DesUnidadLongitud       string   `xml:"ns1:Des_UnidadLongitud" json:"Des_UnidadLongitud"`
	DesUnidadPeso           string   `xml:"ns1:Des_UnidadPeso" json:"Des_UnidadPeso"`
}

func (d *Data) ConvertToJson() []byte {
	temp := ShipmentObject{
		NumGuia:        0,
		NumSobreporte:  0,
		DocRelacionado: "41566",
	}
	temp2 := Shipment{
		Object: &temp,
	}
	d.ShipmentObject = &temp2
	rawXMLData := `<ns1:objEnvios>
	<ns1:EnviosExterno>
	<ns1:Num_Guia></ns1:Num_Guia>
	<ns1:Num_Sobreporte></ns1:Num_Sobreporte>
	<ns1:Doc_Relacionado></ns1:Doc_Relacionado>
	<ns1:Num_Piezas></ns1:Num_Piezas>
	<ns1:Des_TipoTrayecto></ns1:Des_TipoTrayecto>
	<ns1:Ide_producto></ns1:Ide_producto>
	<ns1:Des_FormaPago></ns1:Des_FormaPago>
	<ns1:Des_MedioTransporte></ns1:Des_MedioTransporte>
	<ns1:Num_PesoTotal></ns1:Num_PesoTotal>
	<ns1:Num_ValorDeclaradoTotal></ns1:Num_ValorDeclaradoTotal>
	<ns1:Num_VolumenTotal></ns1:Num_VolumenTotal>
	<ns1:Num_BolsaSeguridad></ns1:Num_BolsaSeguridad>
	<ns1:Des_TipoDuracionTrayecto></ns1:Des_TipoDuracionTrayecto>
	<ns1:Des_Telefono></ns1:Des_Telefono>
	<ns1:Des_Ciudad></ns1:Des_Ciudad>
	<ns1:Des_Direccion></ns1:Des_Direccion>
	<ns1:Nom_Contacto></ns1:Nom_Contacto>
	<ns1:Num_ValorLiquidado></ns1:Num_ValorLiquidado>
	<ns1:Des_DiceContener></ns1:Des_DiceContener>
	<ns1:Des_TipoGuia></ns1:Des_TipoGuia>
	<ns1:Num_VlrSobreflete></ns1:Num_VlrSobreflete>
	<ns1:Num_Vlrflete></ns1:Num_Vlrflete>
	<ns1:Num_Descuento></ns1:Num_Descuento>
	<ns1:Num_PesoFacturado></ns1:Num_PesoFacturado>
	<ns1:idePaisOrigen></ns1:idePaisOrigen>
	<ns1:idePaisDestino></ns1:idePaisDestino>
	<ns1:Num_Alto></ns1:Num_Alto>
	<ns1:Num_Ancho></ns1:Num_Ancho>
	<ns1:Num_Largo></ns1:Num_Largo>
	<ns1:Des_DepartamentoDestino></ns1:Des_DepartamentoDestino>
	<ns1:Des_DepartamentoOrigen></ns1:Des_DepartamentoOrigen>
	<ns1:Nom_UnidadEmpaque></ns1:Nom_UnidadEmpaque>
	<ns1:Des_UnidadLongitud></ns1:Des_UnidadLongitud>
	<ns1:Des_UnidadPeso></ns1:Des_UnidadPeso>
	</ns1:objEnvios>
	</ns1:EnviosExterno>`
	rawXMLHeader := `<?xml version="1.0" encoding="UTF-8"?>
	<env:Envelope xmlns:env="http://www.w3.org/2003/05/soap-envelope"
	xmlns:ns1="http://tempuri.org/">
	<env:Header>
	<ns1:AuthHeader>
	<ns1:login>luis1937</ns1:login>
	<ns1:pwd>BpSUh12jBIiWdACDozgOaQ==</ns1:pwd>
	<ns1:Id_CodFacturacion>SER408</ns1:Id_CodFacturacion>
	<ns1:Nombre_Cargue>Colombia1</ns1:Nombre_Cargue>
	</ns1:AuthHeader>
	</env:Header>
	<env:Body>
	<ns1:CargueMasivoExterno>
	<ns1:envios>
	<ns1:CargueMasivoExternoDTO>`
	rawXMLFooter := `</ns1:CargueMasivoExternoDTO>
	</ns1:envios>
	<ns1:arrayGuias />
	</ns1:CargueMasivoExterno>
	</env:Body>
	</env:Envelope>`
	xml.Unmarshal([]byte(rawXMLData), &d.ShipmentObject)
	xmlData, _ := xml.Marshal(&d.ShipmentObject)
	regex := regexp.MustCompile("\\s+")
	rawXMLHeader = regex.ReplaceAllString(rawXMLHeader, " ")
	rawXMLFooter = regex.ReplaceAllString(rawXMLFooter, " ")
	return []byte(rawXMLHeader + string(xmlData) + rawXMLFooter)

}
