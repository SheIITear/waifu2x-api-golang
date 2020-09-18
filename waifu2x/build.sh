#!/bin/sh

# Get files
set -eux && git clone https://github.com/DeadSix27/waifu2x-converter-cpp /opt/waifu2x-cpp

# Build
cd /opt/waifu2x-cpp
cmake .
make
