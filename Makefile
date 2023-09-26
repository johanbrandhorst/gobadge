prepare-gceu:
	go run cmd/main.go -conf=gceu22

flash-gceu: prepare-gceu flash

prepare-gcuk:
	go run cmd/main.go -conf=gcuk22

flash-gcuk: prepare-gcuk flash

prepare-gcus22:
	go run cmd/main.go -conf=gcus22

flash-gcus22: prepare-gcus22 flash

prepare-gcus:
	go run cmd/main.go -conf=gcus23

flash-gcus: prepare-gcus flash

prepare-tinygo:
	go run cmd/main.go -conf=tinygo

flash-tinygo: prepare-tinygo flash

prepare-fosdem:
	go run cmd/main.go -conf=fosdem23

flash-fosdem: prepare-fosdem flash

prepare-boundary:
	go run cmd/main.go -conf=boundary

flash:
	sudo udisksctl mount -b /dev/sdb
	sudo tinygo build -o /run/media/root/PYBADGEBOOT/firmware.uf2 -target gobadge.json -ldflags="-X main.YourName='Johan' -X main.YourTitleA1='Boundary' -X main.YourTitleA2='at HashiCorp' -X main.YourQRText='https://jbrandhorst.com' -X main.YourTitleB1='I made this' -X main.YourTitleB2='using TinyGo'" .
