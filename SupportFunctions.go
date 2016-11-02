package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Get the resolution from VUEWorks by sending the ID
func getResolutionFromVUEWorks(s string) string {
	// http://stackoverflow.com/questions/29564032/xml-http-post-in-go
	var body bytes.Buffer
	body.WriteString("AppAuthKey=123456&ProjectName=City State&ID=")
	body.WriteString(s)

	client := &http.Client{}
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", "http://YourServer/vueworks/webservices/servicecallws.asmx/GetRequestDataByID", bytes.NewBuffer([]byte(body.String())))
	if err != nil {
		//	fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	return (string(contents))

}

func update311(id string) string {

	var body bytes.Buffer
	body.WriteString(`<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"><s:Header>
<h:ClientInfoHeader xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="urn:messages.ws.rightnow.com/v1_3" xmlns:h="urn:messages.ws.rightnow.com/v1_3">
<AppID>Basic Update</AppID>
</h:ClientInfoHeader>
<o:Security xmlns:o="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" s:mustUnderstand="1">
<o:UsernameToken>
<o:Username>USERNAME</o:Username>
<o:Password>PASSWORD</o:Password>
</o:UsernameToken>
</o:Security>
</s:Header>
<s:Body xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<Update xmlns="urn:messages.ws.rightnow.com/v1_3">
<RNObjects xmlns:q1="urn:objects.ws.rightnow.com/v1_3" xsi:type="q1:Incident">
<ID xmlns="urn:base.ws.rightnow.com/v1_3" id="`)

	body.WriteString(id)
	body.WriteString(`"/>
<q1:CustomFields>
   <ObjectType xmlns="urn:generic.ws.rightnow.com/v1_3">
      <Namespace xsi:nil="true"/>
      <TypeName>IncidentCustomFields</TypeName>
   </ObjectType>
  <GenericFields dataType="OBJECT" name="c" xmlns="urn:generic.ws.rightnow.com/v1_3">
      <DataValue>
         <ObjectValue>
            <ObjectType>
                <Namespace xsi:nil="true"/>
                <TypeName>IncidentCustomFieldsc</TypeName>
            </ObjectType>
               <GenericFields dataType="STRING" name="additional_notes">
                  <DataValue>
                  <StringValue>Resolved. Closed in VUEWorks</StringValue>
                  </DataValue>
               </GenericFields>
            </ObjectValue>
      </DataValue>
   </GenericFields>
</q1:CustomFields>
<q1:StatusWithType><q1:Status><ID xmlns="urn:base.ws.rightnow.com/v1_3" id="2"/>
</q1:Status>
</q1:StatusWithType>
</RNObjects>
<ProcessingOptions>
<SuppressExternalEvents>false</SuppressExternalEvents>
<SuppressRules>false</SuppressRules>
</ProcessingOptions>
</Update>
</s:Body>
</s:Envelope>`)

	//log.Println(body.String())
	client := &http.Client{}
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", "https://YourCity.311.com/cgi-bin/city.cfg/services/soap", bytes.NewBuffer([]byte(body.String())))
	if err != nil {
		//	fmt.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("soapAction", "basicUpdate; charset=utf-8")
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
	//st := strings.Split(c, ">")
	//s2 := st[14]
	//s3 := strings.Split(s2, "<")
	//id := s3[0]

	//log to file

	fileHandle, _ := os.Create("C:\\Users\\Me\\Desktop\\logs\\" + id + ".txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, c)
	writer.Flush()

	//to file

	return (c)

}

func WrongDept(rid string, dept string) {

	var body bytes.Buffer
	body.WriteString(`<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
	<s:Header>
		<h:ClientInfoHeader xmlns:h="urn:messages.ws.rightnow.com/v1_3" xmlns="urn:messages.ws.rightnow.com/v1_3" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
			<AppID>Basic Update</AppID>
		</h:ClientInfoHeader>
		<o:Security s:mustUnderstand="1" xmlns:o="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
			<o:UsernameToken >
				<o:Username>USERNAME</o:Username>
				<o:Password>PASSWORD</o:Password>
			</o:UsernameToken>
		</o:Security>
	</s:Header>
	<s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
		<Update xmlns="urn:messages.ws.rightnow.com/v1_3">
			<RNObjects xsi:type="q1:Incident" xmlns:q1="urn:objects.ws.rightnow.com/v1_3">
				<ID id="`)
	body.WriteString(rid)
	body.WriteString(`" xmlns="urn:base.ws.rightnow.com/v1_3"/>
				<q1:CustomFields>
					<ObjectType xmlns="urn:generic.ws.rightnow.com/v1_3">
						<Namespace xsi:nil="true"/>
						<TypeName>IncidentCustomFields</TypeName>
					</ObjectType>
					<GenericFields dataType="OBJECT" name="c" xmlns="urn:generic.ws.rightnow.com/v1_3">
						<DataValue>
							<ObjectValue>
								<ObjectType>
									<Namespace xsi:nil="true"/>
									<TypeName>IncidentCustomFieldsc</TypeName>
								</ObjectType>
								<GenericFields dataType="BOOLEAN" name="decision_maker">
									<DataValue>
											<BooleanValue>false</BooleanValue>
									</DataValue>
								</GenericFields>
								</ObjectValue>
								</DataValue>
								</GenericFields>
								</q1:CustomFields>
				<q1:StatusWithType>
					<q1:Status>
						<ID id="8" xmlns="urn:base.ws.rightnow.com/v1_3"/>
					</q1:Status>
				</q1:StatusWithType>
				<q1:Threads>
					<q1:ThreadList action="add">
						<q1:EntryType>
							<ID id="4" xmlns="urn:base.ws.rightnow.com/v1_3"/>
						</q1:EntryType>
						<q1:MailHeader xsi:nil="true"/>
						<q1:Text>`)
	body.WriteString(dept)
	body.WriteString(`</q1:Text>
						<q1:ValidNullFields xsi:nil="true"/>
					</q1:ThreadList>
				</q1:Threads>
			</RNObjects>
			<ProcessingOptions>
				<SuppressExternalEvents>false</SuppressExternalEvents>
				<SuppressRules>false</SuppressRules>
			</ProcessingOptions>
		</Update>
	</s:Body>
</s:Envelope>`)
	//log.Println(body.String())

	client := &http.Client{}
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", "https://YourCity.311.com/cgi-bin/city.cfg/services/soap", bytes.NewBuffer([]byte(body.String())))
	if err != nil {
		//	fmt.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("soapAction", "basicUpdate; charset=utf-8")
	// now POST it

	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
	}
	//fmt.Println(resp)
	//fmt.Printf("%+v\n", resp)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	c := string(data)
	//log to file

	fileHandle, _ := os.Create("C:\\Users\\Me\\Desktop\\logs\\wrongdept\\" + rid + ".txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, c)
	writer.Flush()

	//to file

	//log.Println(contents)

}
