#!/bin/bash

# This script fixes the missing noise files error and creates needed directiories for the api to function. This script needs sudo privs.

# Creates dir for noise files
mkdir /usr/local/share/waifu2x-converter-cpp/

# Creates the dir where photos are saved
mkdir /data

# Move the noise files
cd noise_files
cp *.* /usr/local/share/waifu2x-converter-cpp/

echo "Done"
