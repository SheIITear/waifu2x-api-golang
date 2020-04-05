#!/bin/sh

OpenCV_URL="https://github.com/opencv/opencv/archive/3.4.3.tar.gz"

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
