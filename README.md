# student-api

An API for processing student details. The database used is postgres
## Environment variables
- DB_HOST : Host of the postgres database. Default value = postgres
## How to run 

start the service by using docker-compose command
```
cd student-api
docker-compose up -d
```

## Requests
### POST
Add Student details to DB <br>
**Url** : http://localhost:8080/students
<br> **body**: 
```json
    {
	"admission_no": 1,
	"name": "Anu",
	"address": "Kerala",
	"class": "CSE",
	"age": 28
    }
```
### GET
Get Student details from DB for given admission number <br>
**Url** : http://localhost:8080/students/{admission_no}
### PUT
Can update students address, class and name using admission number. Cannot update admission number and name <br>
**Url** : http://localhost:8080/students/{admission_no}
<br> **body**: 
```json
    {
	"address": "Trivandrum",
	"class": "CSE 1",
	"age": 21
    }
```
### DELETE
Remove Student details from DB for given admission number<br>
**Url** : http://localhost:8080/students/{admission_no}