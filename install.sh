echo "FMRADSTR: Updating repositories..."
sudo apt update
echo "FMRADSTR: Trying to install apt dependencies..."
sudo apt install ffmpeg libsndfile1-dev libsoxr-dev git build-essential python3-pip opus-tools && echo "FMRADSTR: Successfully installed apt dependencies!" || echo "FMRADSTR: Failed to install apt dependencies!"
echo "FMRADSTR: Trying to install/upgrade youtube-dl..."
sudo sudo pip install --upgrade youtube_dl && echo "FMRADSTR: Successfully installed youtube-dl!" || echo "FMRADSTR: Failed to install youtube-dl!"
echo "FMRADSTR: Creating directories..."
mkdir "core"
echo "FMRADSTR: Downloading PiFmAdv..."
git submodule update --init --recursive
echo "FMRADSTR: Moving and compiling PiFmAdv..."
mv PiFmAdv/src/* core/ && echo "FMRADSTR: Moved it to core directory!" || echo "FMRADSTR: Failed to move it to core directory!"
cd core && make clean && echo "FMRADSTR: Core cleaned!" || echo "FMRADSTR: Failed to clean core!"
make && echo "FMRADSTR: Successfully compiled!" || echo "FMRADSTR: Failed to compile!"
cd ../
echo "Done. Please set YT API key in config.json and DO NOT FORGET to add 'gpu_freq=250' in /boot/config.txt. Enjoy! :)"
