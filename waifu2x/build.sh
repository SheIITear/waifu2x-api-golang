#!/bin/sh

# get path 
ALOTUS=$(pwd)

################################   OPENCV   ################################

# set url
OpenCV_URL="https://github.com/opencv/opencv/archive/3.4.3.tar.gz"

# create dir
mkdir src

# get opencv files
wget -q -O - $OpenCV_URL | tar zxvf - -C src --strip-components 1

# create cmake dir and cmake
mkdir cmake_tmp
cd cmake_tmp
cmake -DCMAKE_INSTALL_PREFIX=/usr ../src

# make and make install
make
make install

################################   WAIFU2X   ################################

# get files
set -eux && git clone https://github.com/DeadSix27/waifu2x-converter-cpp /opt/waifu2x-cpp

# build
cd /opt/waifu2x-cpp
cmake .
make

################################     FIX    ################################

# this part fixes the missing noise files error and creates needed directiories for the api to function. This script needs sudo privs.

# go back to starting dir
cd $ALOTUS

# creates dir for noise files
mkdir /usr/local/share/waifu2x-converter-cpp/

# creates the dir where photos are saved
mkdir /data

# move the noise files
cd noise_files
cp *.* /usr/local/share/waifu2x-converter-cpp/

echo "Done"
