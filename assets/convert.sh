#!/bin/bash

for f in *
do
  short=$(echo $f | cut -d '.' -f1)
  shortNoDash=$(echo $short | sed "s/-//g")
  echo $short
  ~/Go/bin/file2byteslice -input $f -output $short.go -package assets -var $shortNoDash
done
