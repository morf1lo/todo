build:
	docker build -t todo .

run:
	docker run --name todo-app -p 70:8080 todo