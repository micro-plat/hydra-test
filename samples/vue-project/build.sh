

r=`npm run build`

sleep 1 
cd dist

rm -f dist.zip

zip -q -r dist.zip ./*

mv dist.zip  ../../$1

sleep 1

pwd
sh ../../$1/build.sh 