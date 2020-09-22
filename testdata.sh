#!/bin/zsh

rm -r testdata

mkdir testdata
mkdir testdata/img

curl http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.jpg > ./testdata/azarashi.jpg
curl http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.jpg > ./testdata/img/azarashi.jpg

curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/tanuki.jpg
curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/img/tanuki.jpg

curl http://flat-icon-design.com/f/f_object_157/s512_f_object_157_0bg.png > ./testdata/osaru.png

curl https://1.bp.blogspot.com/-aE1d5IM9FuY/XwkxlUXaVYI/AAAAAAABaC4/5Y9ueDFkL-ktyOHXczaj5dvlCXieT-pigCNcBGAsYHQ/s1600/syougatsu_mark_mochi.png > ./testdata/mochi.png
