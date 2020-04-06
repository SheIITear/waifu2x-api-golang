# waifu2x-api-golang
This is a bit edited version from https://github.com/nothink/docker-waifu2x-converter-cpp and https://github.com/gladkikhartem/waifurun that runs on ubuntu and doesn't require gpu.

Requirements:

> cmake +3.8,
> golang

Before you can run the .sh files, you need to make them runable by running:

> chmod +x script.sh

To run this, first build the driver:
> sudo apt install ocl-icd-opencl-dev

> cd waifu2x && sudo ./build-opencv.sh && sudo ./build.sh

Then run RUNME.sh to create needed directories and move files:
> sudo RUNME.sh

Lastly to run the api itself:
> sudo go get github.com/gookit/color && sudo go get github.com/gorilla/mux

> sudo go run main.go
