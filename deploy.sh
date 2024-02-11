#!/bin/bash

file=".version" #the file where you keep your string name

current_version=$(cat "$file")        #the output of 'cat $file' is assigned to the $name variable

go build -o grinha-de-backend-2024-q1-demode

new_version=$(echo $current_version | awk -F. -v OFS=. 'NF==1{print ++$NF}; NF>1{if(length($NF+1)>length($NF))$(NF-1)++; $NF=sprintf("%0*d", length($NF), ($NF+1)%(10^length($NF))); print}')

echo $new_version > $file
new_image="grinha-http2:v$new_version"


echo "Deploying $new_image"

docker build -t $new_image .
docker tag $new_image registry.geff.ws/rinha2024q1/$new_image
docker push registry.geff.ws/rinha2024q1/$new_image

echo "Nova versão: $new_version"
