# docker-waifu2x-converter-cpp

## usage

```
docker run --rm -it -v `pwd`:/srv/waifu2x nothink/waifu2x -- -i /srv/waifu2x/src.jpg -o /srv/waifu2x/dst_2x.png
```


## --help

```
USAGE:

   /opt/waifu2x-cpp/waifu2x-converter-cpp  [--list-opencv-formats]
                                        [--list-processor] [--block_size
                                        <integer>] [--disable-gpu]
                                        [--force-OpenCL] [-p <integer>] [-j
                                        <integer>] [--model_dir <string>]
                                        [--scale_ratio <double>]
                                        [--noise_level <1|2|3>] [-m <noise
                                        |scale|noise_scale>] [-q] [-r
                                        <bool>] [-o <string>] -i <string>
                                        [--] [--version] [-h]


Where:

   --list-opencv-formats
     dump opencv supported format list

   --list-processor
     dump processor list

   --block_size <integer>
     block size

   --disable-gpu
     disable GPU

   --force-OpenCL
     force to use OpenCL on Intel Platform

   -p <integer>,  --processor <integer>
     set target processor

   -j <integer>,  --jobs <integer>
     number of threads launching at the same time

   --model_dir <string>
     path to custom model directory (don't append last / )

   --scale_ratio <double>
     custom scale ratio

   --noise_level <1|2|3>
     noise reduction level

   -m <noise|scale|noise_scale>,  --mode <noise|scale|noise_scale>
     image processing mode

   -q,  --quiet
     Enable quiet mode.

   -r <bool>,  --recursive_directory <bool>
     Search recursively through directories to find more images to process.


     If this is set to 0 it will only check in the directory specified if
     the input is a directory instead of an image.

     You mustn't supply this argument with something other than 0 or 1.

   -o <string>,  --output <string>
     path to output image file or directory  (you should use the full path)

   -i <string>,  --input <string>
     (required)  path to input image file or directory (you should use the
     full path)

   --,  --ignore_rest
     Ignores the rest of the labeled arguments following this flag.

   --version
     Displays version information and exits.

   -h,  --help
     Displays usage information and exits.


   waifu2x OpenCV Fork - https://github.com/DeadSix27/waifu2x-converter-cpp
``
