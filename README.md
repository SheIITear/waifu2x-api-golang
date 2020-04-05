# waifu2x-api-golang
This is a bit edited version from https://github.com/nothink/docker-waifu2x-converter-cpp and https://github.com/gladkikhartem/waifurun that runs on ubuntu and doesn't require gpu.

To run this, first build the driver:
> sudo apt install beignet-opencl-icd opencl-headers 

> cd waifu2x && sudo ./build-opencv.sh && sudo ./build.sh

Then run RUNME.sh to create needed directories and move files:
> sudo RUNME.sh

Lastly to run the api itself:
> go get github.com/gookit/color

> go get github.com/gorilla/mux

> sudo go run main.go
