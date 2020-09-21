#!/bin/sh

# get path 
ALOTUS=$(pwd)
OPENCV_DIR="$ALOTUS/src"
WAIFU2X_DIR="/opt/waifu2x-cpp"
FIX_DIR="/usr/local/share/waifu2x-converter-cpp/"
FILE_DIR="/data"

################################   OPENCV   ################################

if [ -d "$OPENCV_DIR" ]; then
        echo "Opencv already built"
    else
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
        make -j4
        make install
fi

################################   WAIFU2X   ################################

if [ -d "$WAIFU2X_DIR" ]; then
        echo "Waifu2x already built"
    else
        # get files
        set -eux && git clone https://github.com/DeadSix27/waifu2x-converter-cpp /opt/waifu2x-cpp

        # build
        cd /opt/waifu2x-cpp
        cmake .
        make -j4
fi

################################     FIX    ################################

# this part fixes the missing noise files error and creates needed directiories for the api to function. This script needs sudo privs.

# go back to starting dir
cd $ALOTUS

# creates dir for noise files, if doesn't exist
if [ -d "$FIX_DIR" ]; then
        echo "Noise files already moved"
    else
        mkdir /usr/local/share/waifu2x-converter-cpp/
        
        # move the noise files
        cd noise_files
        cp *.* /usr/local/share/waifu2x-converter-cpp/
fi

# creates the dir where photos are saved, if not already made
if [ -d "$FIX_DIR" ]; then
        echo "Already exists"
    else 
        mkdir /data
fi

echo "Finished"
