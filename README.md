

# API

POST /dicom content-type:multipart/form-data
=> 201 url resource id

GET /dicom/{id}?tags[]=%280008%2C0025%29

GET /dicom/{id} -H accept-type=image/png

## URL encode dicom tags
(0008,0025) => %280008%2C0025%29

(0008,0025),(0008,0026) => %280008%2C0025%29%2C%280008%2C0026%29

# Structure

## cmd

```
go run cmd/main.go 
```
## handlers

contains:
new(repository, processor)
uploadHandler(file))
getTagsHandler(id, tags)
getPngHandler(id)

```
file := repo.get(id)
return processor.HeaderAttrs(file, tags)
```

getPng(id)

## processor
HeaderAttrs(file, tags) -> []HeaderAttribute
AsPng(file)-> [][]byte

## repository

add()
get(id)



