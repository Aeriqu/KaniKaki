

# generate certs for k8s
.PHONY: generate-ssl
generate-ssl:
	generate-ssl-ca
	generate-ssl-server

.PHONY: generate-ssl-ca
generate-ssl-ca:
	cfssl gencert -initca ./certs/ca/ca-csr.json | cfssljson -bare ./certs/ca/ca

.PHONY: generate-ssl-server
generate-ssl-server:
	cfssl gencert -ca=./certs/ca/ca.pem -ca-key=./certs/ca/ca-key.pem --config=./certs/ca/ca-config.json -profile=kubernetes ./certs/server/server-csr.json | cfssljson -bare ./certs/server/server


# build the apps
.PHONY: build-all
build-all: proto-gen hello-world api-gateway auth web

.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/hello-world/proto/hello_world.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/auth/proto/auth.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/wanikani/proto/wanikani.proto

.PHONY: hello-world
hello-world:
	docker build -t aeriqu/kanikaki/hello-world:latest --file ./services/hello-world/Dockerfile .

.PHONY: api-gateway
api-gateway:
	docker build -t aeriqu/kanikaki/api-gateway:latest --file ./services/api-gateway/Dockerfile .

.PHONY: auth
auth:
	docker build -t aeriqu/kanikaki/auth:latest --file ./services/auth/Dockerfile .

.PHONY: web
web:
	docker build -t aeriqu/kanikaki/web:latest --file ./services/web/Dockerfile .

.PHONY: wanikani
wanikani:
	docker build -t aeriqu/kanikaki/wanikani:latest --file ./services/wanikani/Dockerfile .


# environment set up
.PHONY: docker-k8s-ingress
docker-k8s-ingress:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.7.0/deploy/static/provider/cloud/deploy.yaml

# deploy
.PHONY: deploy
deploy:
	terraform -chdir="./terraform" apply -auto-approve

.PHONY: destroy
destroy:
	terraform -chdir="./terraform" destroy -auto-approve

.PHONY: redeploy
redeploy: destroy build-all deploy