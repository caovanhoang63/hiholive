#!/bin/sh
/usr/bin/s3fs hiholive-stream-storage hls_output -o nosuid,nonempty,nodev,allow_other,passwd_file=.passwd-s3fs
./app
