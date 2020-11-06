build:
	go build -o output/adminbot_output
	cp ./config.json ./output/config.json

run:
	go build -o output/adminbot_output
	cp ./config.json ./output/config.json
	cd ./output && ./adminbot_output


