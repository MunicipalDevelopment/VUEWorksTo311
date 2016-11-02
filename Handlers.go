package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ------------------------------------------------------------------
//        Structure to get the ID from VUEWorks from the XML sent to server at / calling IncomingVUEWorksServiceReqClosed

type vwID struct {
	XMLName  xml.Name `xml:"Envelope"`
	Bodydata []Body   `xml:"Body"`
}

type Body struct {
	Closed []SCClosed `xml:"ServiceRequestClosed"`
}

type SCClosed struct {
	ID string `xml:"ID"`
}

//------------------------------------------------------------------------

//----------------------------------------------
//    Parse response from VUEWorks

type Response struct {
	XMLName  xml.Name `xml:"string"`
	Response string   `xml:",chardata"`
}

//---------------------------

func isAlive(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<html><head><title>Alive</title></head><body><img src='http://ibrandolab.com/lab/wp-content/uploads/2016/06/its-alive-1200x475.jpg'></body></html>")

}

func IncomingVUEWorksServiceReqClosed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//r.ParseForm()

	// VUEWORKS calls the server dmdview:8311/ and posts the XML data with ID in it.
	//I parse the ID
	body, _ := ioutil.ReadAll(r.Body)
	id := vwID{}
	xml.Unmarshal(body, &id)

	//To get the resolution and reference id for 311, call VUEWorks and pass the ID
	resolution := getResolutionFromVUEWorks(id.Bodydata[0].Closed[0].ID)

	//Parse out the ref ID and resoultion
	res := Response{}

	xml.Unmarshal([]byte(resolution), &res)

	//XML contains string inside that isnt xml. Thought about parsing it as xml but this works.
	DeleteFirstPart := strings.Split(res.Response, "<ActionDesc>")
	DeleteSecondPart := strings.Split(DeleteFirstPart[1], "</ActionDesc>")
	ResolutionText := DeleteSecondPart[0]

	GetRefID := strings.Split(res.Response, "<Ref_ID>")
	GetRefIDSecond := strings.Split(GetRefID[1], "</Ref_ID>")
	RefID := GetRefIDSecond[0]

	if strings.Contains(ResolutionText, "wd") {
		WrongDept(RefID, ResolutionText)
		log.Println("************** Wrong Department ***************" + RefID + "   " + ResolutionText)
	} else {
		update311(RefID)
	}

}

func total(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	body.WriteString(`<?xml version="1.0"?>
<soapenv:Envelope xmlns:v1="urn:messages.ws.rightnow.com/v1_3" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
<soapenv:Header>
<h:ClientInfoHeader xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="urn:messages.ws.rightnow.com/v1_3" xmlns:h="urn:messages.ws.rightnow.com/v1_3">
<AppID>Basic Update</AppID>
</h:ClientInfoHeader>
<o:Security xmlns:o="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" soapenv:mustUnderstand="1">
<o:UsernameToken>
<o:Username>USERNAME</o:Username>
<o:Password>PASSWORD</o:Password>
</o:UsernameToken>
</o:Security>
</soapenv:Header>
<soapenv:Body>
<v1:QueryCSV>
<v1:Query>select COUNT(*) from Incident where Incident.CustomFields.c.additional_notes LIKE '%VUEWorks%';</v1:Query>
<v1:Delimiter>|</v1:Delimiter>
</v1:QueryCSV>
</soapenv:Body>
</soapenv:Envelope>`)

	client := &http.Client{}
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", "https://YourCity.311.com/cgi-bin/City.cfg/services/soap", bytes.NewBuffer([]byte(body.String())))
	if err != nil {
		//	fmt.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("soapAction", "queryCSV; charset=utf-8")
	// now POST it
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
	}
	//fmt.Println(resp)
	//fmt.Printf("%+v\n", resp)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		//	fmt.Printf("%s", err)
		//	os.Exit(1)
	}
	//fmt.Printf("%s\n", string(contents))
	c := string(contents)
	first := strings.Split(c, "<n0:Row>")
	second := strings.Split(first[1], "</n0:Row>")
	theCount := second[0]
	fmt.Fprintln(w, "<html><head><title>Total Closed by VUEWorks</title></head><body><h1>Total 311 Incidents Closed by VUEWorks</h1><h1>"+theCount+"</h1>")
}
