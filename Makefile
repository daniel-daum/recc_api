test:
	go test
	
build:
	@docker build -t recc .

run:
	@docker run -it --name recc recc
