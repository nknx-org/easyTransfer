env GOOS=linux GOARCH=arm GOARM=5 go build -o build/easyTransfer_armV6
env GOOS=linux GOARCH=arm GOARM=6 go build -o build/easyTransfer_armV7
env GOOS=linux GOARCH=arm GOARM=7 go build -o build/easyTransfer_armV8
env GOOS=linux GOARCH=amd64 go build -o build/easyTransfer_amd64
