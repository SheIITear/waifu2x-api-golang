# waifu2x-api-golang
This is a bit edited version from https://github.com/nothink/docker-waifu2x-converter-cpp and https://github.com/gladkikhartem/waifurun that runs on ubuntu and doesn't require gpu.

To build the driver run:
> sudo apt install beignet-opencl-icd opencl-headers 

> cd waifu2x && sudo ./build-opencv.sh && sudo ./build.sh

To run the api run:
> go get github.com/gookit/color

> go get github.com/gorilla/mux

> go run main.go
