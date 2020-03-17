sudo apt-get update
sudo apt-get install ffmpeg libsndfile1-dev
cd core/ || exit
make clean
make
cd ../
sudo npm i
echo Now it\'s time to configure RDS in and type YouTube v3 key into config.json. Do it manually. 
