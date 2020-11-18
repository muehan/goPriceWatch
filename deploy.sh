npm install
npm run build
~/go/bin/bee pack -be "-tags prod" -exr="^(?:node_modules)$"
cp goPriceWatch.tar.gz /var/www/keramic.ch/goPriceWatch.tar.gz
cd /var/www/keramic.ch/ 
tar -xvzf *.tar.gz
rm -rf *.tar.gz
