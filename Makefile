generate:
	protoc -I proto/ proto/student/student.proto --go-grpc_out=proto/student --go_out=proto/student --go-grpc_opt=require_unimplemented_servers=false
	protoc -I proto/ --grpc-gateway_out proto/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true proto/student/student.proto

clean:
	rm proto/student/*.go
