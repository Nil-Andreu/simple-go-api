docker:
	docker build -t simple-api:v1 .
	docker run -p 8080:8080 simple-api:v1