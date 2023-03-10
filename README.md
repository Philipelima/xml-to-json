# xml-to-json converter

![go](https://img.shields.io/static/v1?label=Golang&labelColor=07a0f8&message=1.19.5&color=000000&logo=go&logoColor=ffffff&style=flat-square)
<p align="center">
<img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/4c14a00f0a14e373b7f4d0639a02477403fb9df8.png" height=300>
</p>

<br>

A package to convert xml files to json. 

#### How you can use it on your project:

First of all, use the  **go get** command to import the converter package from repository



```bash
go get github.com/philipelima/xml-to-json
```

So you can import it on your code

```go
    import "github.com/philipelima/xml-to-json"
```
the xml with which we're going to work:

```xml
    <note>
        <to>Peter Parker</to>
        <from>Tony Stark</from>
        <heading>Reminder</heading>
        <body>Don't forget your costume</body>
    </note>
```

After importing the package, let's define how we want to read the xml that we are going to work with and how we want it to appear as json, we can do this from a single struct.

```go
    type Notes struct {
	    To        string    `xml:"to"      json:"to"`; 
	    From      string    `xml:"from"    json:"from"`
	    Heading   string    `xml:"heading" json:"heading"`
	    Body      string    `xml:"body"    json:"note_body"`
    }
```

Inside your function, let's create a <code>Notes</code> reference

```go
    notes := &Notes{}
```

After that, we can call the function responsible for the conversion, passing the xml file path and the reference we created and finally get the json:

```go
    json := xmltojson.JsonFromXml("path/to/xml/file", notes)
```

The final code is similar to this:

```go

    type Notes struct {
	    To        string    `xml:"to"      json:"to"`; 
	    From      string    `xml:"from"    json:"from"`
	    Heading   string    `xml:"heading" json:"heading"`
	    Body      string    `xml:"body"    json:"note_body"`
    }

    func ConvertFunction() string {

        notes := &Notes{}
        
        json  := xmltojson.JsonFromXml("path/to/xml/file", notes)

        return json
    }
```