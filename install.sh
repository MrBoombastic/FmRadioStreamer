echo "FMRADSTR: Updating repositories..."
sudo apt update
echo "FMRADSTR: Trying to install apt dependencies..."
sudo apt install ffmpeg libsndfile1-dev libsoxr-dev git build-essential libsox-fmt-all && echo "FMRADSTR: Successfully installed apt dependencies!" || echo "FMRADSTR: Failed to install apt dependencies!"
echo "FMRADSTR: Creating directories..."
mkdir "core"
mkdir "music"
echo "FMRADSTR: Downloading PiFmAdv..."
git submodule update --init --recursive && echo "FMRADSTR: Done!" || echo "FMRADSTR: Not a git repository! Downloading directly..." && git clone https://github.com/miegl/PiFmAdv.git && git checkout 7562fe79789ec4cc0d209a7b14b07d79cdf6e310
echo "FMRADSTR: Compiling PiFmAdv..."
mv PiFmAdv/src/* core/ && echo "FMRADSTR: Moved it to core directory!" || echo "FMRADSTR: Failed to move it to core directory!"
cd core && make clean && echo "FMRADSTR: Core cleaned!" || echo "FMRADSTR: Failed to clean core!"
make && echo "FMRADSTR: Successfully compiled!" || echo "FMRADSTR: Failed to compile!"
cd ../
echo "FMRADSTR: Done. Please set YT API key in config.json and DO NOT FORGET to add 'gpu_freq=250' in /boot/firmware/config.txt!"
echo "FMRADSTR WARNING: Make sure to always use up-to-date version of youtube_dl!"
echo "FMRADSTR: Enjoy! :)"
