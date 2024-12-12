#!/bin/sh
/usr/bin/s3fs hiholive-stream-storage hls_output -o passwd_file=${HOME}/.passwd-s3fs
./app
