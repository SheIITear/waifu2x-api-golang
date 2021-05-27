# waifu2x-api-golang
This is an site/api made using waifu2x that runs on ubuntu and doesn't require gpu. (heavily cpu depended)

~It can be tested on: https://upgradewaifu.gq~ Down

Requirements:

> cmake +3.8 and 
> golang +1.14.1

Before you can run the .sh files, you need to make them runable by running:

> chmod +x waifu2x/build.sh

To run this, first build it:

> sudo apt install ocl-icd-opencl-dev

> sudo waifu2x/build.sh

Lastly to run the api itself:

> sudo go run main.go
