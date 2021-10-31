#Dummy Server
___

Http test server written in Golang.

Can get any request method and log headers, body, path and remote address.
Useful if you need to know request details from undocumented clients.
Logs like

`time=2021-10-31T17:49:08Z level=info msg=Got request Body={"test": 1} Headers=Cache-Control=[no-cache] Postman-Token=[dc7d20fd-73c3-44fd-8094-ff3d73371d7c] Accept-Encoding=[gzip, deflate, br] Connection=[keep-alive] Content-Length=[11] Content-Type=[application/json] User-Agent=[PostmanRuntime/7.26.8] Accept=[*/*] Authorization=[Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoiNTJlZTYyN2UtZWVjYi00NDFiLTlhNTItYzg4NzI4NTBmOWEyIiwidXNlcm5hbWUiOiJjb21wYW55IiwiZXhwIjoxNTQ5OTM4OTE2LCJlbWFpbCI6IiJ9.DZQCGZqWA2m6i9DpioE1R5tQX_u9YPuTG-oPQIirAlg]  Method=POST Path=/api/v1/licence/ RemoteAddr=172.17.0.1:65324`

###Running

Most easy

docker run -p 8080:8080 shadrus/dummyserver:latest

Can also be deployed to Kubernetes:

kubectl apply -f https://github.com/shadrus/DummyServer/blob/main/deployments/dummyserver.yaml

###Configuration
Can be configured by optional environment variables:
1. port - what port to listen to. Default port is 8080.
2. logPath - where to store log. If omitted, then log to stderr.