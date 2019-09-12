for pofile in `find ./local -name *.po`
do
    mofile=`echo "$pofile" | sed 's/\.po/.mo/g'`
    msgfmt --statistics -v -o $mofile $pofile
done
zip -q -r local-test.zip  ./local
go-bindata -o=local_pack.go -pkg=main local.zip
