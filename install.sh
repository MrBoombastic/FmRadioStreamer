sudo apt update
sudo apt install ffmpeg && echo "Successfully installed ffmpeg!" || echo "Failed to install ffmpeg!"
sudo apt install python3-pip && sudo sudo pip install --upgrade youtube_dl && echo "Successfully installed youtube-dl!" || echo "Failed to install youtube-dl!"
rm core/placeholder.txt
git clone https://github.com/miegl/PiFmAdv temp/ && echo "Successfully cloned PiFmAdv repo!" || echo "Failed to clone PiFmAdv repo!"
mv temp/src/* core/ && echo "Moved it to core" || echo "Failed to move it to core"
sudo rm -r temp/
cd core && make clean && echo "Core cleaned." || echo "Failed to clean core."
make && echo "Successfully compiled!" || echo "Failed to compile!"
cd ../
yarn && echo "Successfully installed npm modules!" || echo "Failed to install npm modules!"
echo "Now it's time to configure RDS and type YouTube v3 key into config.json. Please do it manually."
