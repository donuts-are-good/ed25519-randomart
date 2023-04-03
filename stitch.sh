ffmpeg -r 60 -f image2 -s 512x512 -i img/ed25519_key_%02d.png -vcodec libx264 -crf 25 -pix_fmt yuv420p output.mp4
