# build the apps
.PHONY: build-all
build-all: proto-gen api-gateway auth kanji wanikani web

.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/auth/proto/auth.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/kanji/proto/kanji.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/wanikani/proto/wanikani.proto

.PHONY: api-gateway
api-gateway:
	docker build -t aeriqu/kanikaki/api-gateway:latest --file ./services/api-gateway/Dockerfile .

.PHONY: auth
auth:
	docker build -t aeriqu/kanikaki/auth:latest --file ./services/auth/Dockerfile .

.PHONY: kanji
kanji:
	docker build -t aeriqu/kanikaki/kanji:latest --file ./services/kanji/Dockerfile .

.PHONY: wanikani
wanikani:
	docker build -t aeriqu/kanikaki/wanikani:latest --file ./services/wanikani/Dockerfile .

.PHONY: web
web:
	docker build -t aeriqu/kanikaki/web:latest --file ./services/web/Dockerfile .


# environment set up
.PHONY: docker-k8s-ingress
docker-k8s-ingress:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.2/deploy/static/provider/cloud/deploy.yaml

# deploy
.PHONY: deploy
deploy:
	terraform -chdir="./terraform" apply -auto-approve

.PHONY: destroy
destroy:
	terraform -chdir="./terraform" destroy -auto-approve

.PHONY: redeploy
redeploy: destroy build-all deploy