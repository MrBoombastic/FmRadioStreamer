echo "Updating repositories..."
sudo apt update
echo "Trying to install apt dependencies..."
sudo apt install ffmpeg libsndfile1-dev libsoxr-dev git build-essentials python3-pip opus-tools && echo "Successfully installed apt dependencies!" || echo "Failed to install apt dependencies!"
echo "Trying to install/upgrade youtube-dl..."
sudo sudo pip install --upgrade youtube_dl && echo "Successfully installed youtube-dl!" || echo "Failed to install youtube-dl!"
echo "Creating directories..."
mkdir "core"
mkdir "music"
echo "Cloning PiFmAdv..."
git clone https://github.com/miegl/PiFmAdv temp/ && echo "Successfully cloned PiFmAdv repo!" || echo "Failed to clone PiFmAdv repo!"
mv temp/src/* core/ && echo "Moved it to core directory!" || echo "Failed to move it to core directory!"
echo "Removing unnecessary files..."
sudo rm -r temp/
cd core && make clean && echo "Core cleaned!" || echo "Failed to clean core!"
echo "Compiling PiFmAdv..."
make && echo "Successfully compiled!" || echo "Failed to compile!"
cd ../
echo "Done. Please set YT API key in config.json and DO NOT FORGET to add 'gpu_freq=250' in /boot/config.txt. Enjoy! :)"
