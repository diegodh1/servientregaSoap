package responserequest

import "encoding/xml"

//CargueMasivoExternoResponse struct
type CargueMasivoExternoResponse struct {
	XMLName xml.Name
	Body    struct {
		XMLName                   xml.Name
		CargueMasivoExternoResult bool   `xml:"CargueMasivoExternoResponse>CargueMasivoExternoResult" json:"CargueMasivoExternoResult"`
		EnvioObj                  *Envio `xml:"CargueMasivoExternoResponse>envios>CargueMasivoExternoDTO>objEnvios>EnviosExterno" json:"Envio"`
		Guia                      string `xml:"CargueMasivoExternoResponse>arrayGuias>string" json:"Guia"`
	}
}

//Envio struct
type Envio struct {
	XMLName                 xml.Name `xml:"EnviosExterno" json:"-"`
	NumGuia                 int      `xml:"Num_Guia" json:"Num_Guia"`
	NumSobreporte           int      `xml:"Num_Sobreporte" json:"Num_Sobreporte"`
	DocRelacionado          string   `xml:"Doc_Relacionado" json:"Doc_Relacionado"`
	NumPiezas               int      `xml:"Num_Piezas" json:"Num_Piezas"`
	DesTipoTrayecto         int      `xml:"Des_TipoTrayecto" json:"Des_TipoTrayecto"`
	Ideproducto             int      `xml:"Ide_producto" json:"Ide_producto"`
	IdeDestinatarios        string   `xml:"Ide_Destinatarios" json:"Ide_Destinatarios"`
	IdeManifiesto           string   `xml:"Ide_Manifiesto" json:"Ide_Manifiesto"`
	DesFormaPago            int      `xml:"Des_FormaPago" json:"Des_FormaPago"`
	DesMedioTransporte      int      `xml:"Des_MedioTransporte" json:"Des_MedioTransporte"`
	NumPesoTotal            int      `xml:"Num_PesoTotal" json:"Num_PesoTotal"`
	NumValorDeclaradoTotal  int      `xml:"Num_ValorDeclaradoTotal" json:"Num_ValorDeclaradoTotal"`
	NumVolumenTotal         int      `xml:"Num_VolumenTotal" json:"Num_VolumenTotal"`
	NumBolsaSeguridad       int      `xml:"Num_BolsaSeguridad" json:"Num_BolsaSeguridad"`
	NumPrecinto             int      `xml:"Num_Precinto" json:"Num_Precinto"`
	DesTipoDuracionTrayecto int      `xml:"Des_TipoDuracionTrayecto" json:"Des_TipoDuracionTrayecto"`
	DesTelefono             int      `xml:"Des_Telefono" json:"Des_Telefono"`
	DesCiudad               string   `xml:"Des_Ciudad" json:"Des_Ciudad"`
	DesDireccion            string   `xml:"Des_Direccion" json:"Des_Direccion"`
	NomContacto             string   `xml:"Nom_Contacto" json:"Nom_Contacto"`
	NumValorLiquidado       int      `xml:"Num_ValorLiquidado" json:"Num_ValorLiquidado"`
	DesDiceContener         string   `xml:"Des_DiceContener" json:"Des_DiceContener"`
	DesTipoGuia             int      `xml:"Des_TipoGuia" json:"Des_TipoGuia"`
	NumVlrSobreflete        int      `xml:"Num_VlrSobreflete" json:"Num_VlrSobreflete"`
	NumVlrFlete             int      `xml:"Num_VlrFlete" json:"Num_VlrFlete"`
	NumDescuento            int      `xml:"Num_Descuento" json:"Num_Descuento"`
	IdePaisOrigen           int      `xml:"idePaisOrigen" json:"idePaisOrigen"`
	IdePaisDestino          int      `xml:"idePaisDestino" json:"idePaisDestino"`
	NumPesoFacturado        int      `xml:"Num_PesoFacturado" json:"Num_PesoFacturado"`
	EstCanalMayorista       bool     `xml:"Est_CanalMayorista" json:"Est_CanalMayorista"`
	NumAlto                 int      `xml:"Num_Alto" json:"Num_Alto"`
	NumAncho                int      `xml:"Num_Ancho" json:"Num_Ancho"`
	NumLargo                int      `xml:"Num_Largo" json:"Num_Largo"`
	DesDepartamentoDestino  string   `xml:"Des_DepartamentoDestino" json:"Des_DepartamentoDestino"`
	DesDepartamentoOrigen   string   `xml:"Des_DepartamentoOrigen" json:"Des_DepartamentoOrigen"`
	GenCajaporte            bool     `xml:"Gen_Cajaporte" json:"Gen_Cajaporte"`
	GenSobreporte           bool     `xml:"Gen_Sobreporte" json:"Gen_Sobreporte"`
	NomUnidadEmpaque        string   `xml:"Nom_UnidadEmpaque" json:"Nom_UnidadEmpaque"`
	DesUnidadLongitud       string   `xml:"Des_UnidadLongitud" json:"Des_UnidadLongitud"`
	DesUnidadPeso           string   `xml:"Des_UnidadPeso" json:"Des_UnidadPeso"`
	EstEnviarCorreo         bool     `xml:"Est_EnviarCorreo" json:"Est_EnviarCorreo"`
}
