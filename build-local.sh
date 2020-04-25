#!/bin/bash

for pofile in `find ./local -name *.po`
do
    mofile=`echo "$pofile" | sed 's/\.po/.mo/g'`
    msgfmt --statistics -v -o $mofile $pofile
done
zip -q -r local.zip  ./local
go-bindata -o=tools/local_pack.go -pkg=tools local.zip
