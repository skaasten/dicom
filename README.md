
# Running 

```bash
go run cmd/main.go
```

## Add a dcm 

```bash
curl -X POST -F "file=@<path to dcm file>" localhost:8080/dicom

```

The response will show the `ID` of the added file, with http status 201

## Get header attributes

To get one or more header attributes, specify the tags as query params
in the form `group:element`.

```bash
curl -i -v "localhost:8080/dicom/8d25efb3-3331-49ed-ab23-b3f44edb6f01?tag=0010:1010"
```
The response will show the specified tags and values.

## Get images as PNG

To get the images encoded as png, use the GET api with the `Accept-type` set to `image/png`.

```bash
curl "localhost:8080/dicom/a538ed64-3ab3-47a5-b71d-238a349291a3"  -H "Accept-type:image/png" -o xray.png
```

# API

## Add a file
POST /dicom content-type:multipart/form-data
=> 201 url with resource id

## Get tags for a file
GET /dicom/{id}?tags[]=%280008%2C0025%29

# Get png files 
GET /dicom/{id} -H accept-type=image/png


# Structure

## cmd
The command runner to start the server

```
go run cmd/main.go 
```


## service

The service that provides the functionality

## processor

Wraps the dicom 3rd party library 

## repository

Simple in memory data store

## handlers

The http handlers that call out to the service

contains:
new(repository)
AddHandler(file))
GetByIdHandler(id)



