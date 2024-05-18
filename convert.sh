#!/bin/sh

for s in 384 192 152 144 128 96 72; do
	convert cj-512x512.png -resize $s cj-${s}x${s}.png
done
