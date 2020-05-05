sudo apt-get update
sudo apt-get install ffmpeg libsndfile1-dev && echo "Successfully installed modules!" || echo "Failed to install modules!"
cd core && make clean && echo "Core cleaned." || echo "Failed to clean core."
make && echo "Successfully compiled!" || echo "Failed to compile!"
cd ../
npm i && echo "Successfully installed npm modules!" || echo "Failed to install npm modules!"
echo "Now it's time to configure RDS and type YouTube v3 key into config.json. Please do it manually."