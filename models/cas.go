package models

import "encoding/xml"

// CASResponse adalah struktur XML standar untuk CAS Service Validation
type CASResponse struct {
	XMLName xml.Name `xml:"http://www.yale.edu/tp/cas serviceResponse"`

	// Jika login sukses
	AuthenticationSuccess struct {
		User       string        `xml:"user"`
		Attributes CASAttributes `xml:"attributes"`
	} `xml:"authenticationSuccess"`

	// Jika login gagal
	AuthenticationFailure struct {
		Code    string `xml:"code,attr"`
		Message string `xml:",chardata"`
	} `xml:"authenticationFailure"`
}

// CASAttributes menangkap data tambahan seperti Nama, NIP, atau Role jika disediakan oleh Undiksha
type CASAttributes struct {
	// Tambahkan field ini sesuai dengan atribut yang dikirim oleh CAS Undiksha
	// Biasanya nama fieldnya case-sensitive tergantung server CAS
	Fullname string `xml:"fullname"`
	NIP      string `xml:"nip"`
	Mail     string `xml:"mail"`
	Group    string `xml:"group"`
}
