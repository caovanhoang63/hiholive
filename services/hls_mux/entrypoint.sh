#!/bin/sh
/usr/bin/s3fs stream.hiholive.fun hls_output -o nosuid,nonempty,nodev,allow_other,use_path_request_style,passwd_file=.passwd-s3fs
./app
