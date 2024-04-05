# API

POST /dicom content-type:multipart/form-data
=> 201 url resource

GET /dicom/{id}?tags[]=%280008%2C0025%29

GET /dicom/{id} -H accept-type=image/png

## URL encode dicom tags
(0008,0025) => %280008%2C0025%29

(0008,0025),(0008,0026) => %280008%2C0025%29%2C%280008%2C0026%29

