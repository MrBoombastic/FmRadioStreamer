cd core/ || exit
make clean
make
cd ../
sudo npm i
sudo apt install ffmpeg
sudo apt install libsndfile1-dev
echo Now it\'s time to configure RDS in and type YouTube v3 key into config.json. Do it manually. 