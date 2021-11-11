FROM ubuntu:18.04

WORKDIR /students
COPY student-api /students/
EXPOSE 8080

CMD ["./student-api"]