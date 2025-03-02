### TEST BRANTCH! ###

IF YOU INSTALL THIS:
- There is a limit for the input values to be 255 in length, it's a hard limit at the moment.
- Working on the Markov chains, but it still is fun to see what it is capable of doing.
- Matrix handling is still a work in progress.
- Inputs _ARE_ encrypted, eventually I plan on making sure that logged data is only the outputs, and not the input values. Course, this is the dev branch, so it's ideal to get an idea.
- Prediction values are basically based on the number / length of the input value. That means you only get a number at this point in time for it.
- The number of matches the current algorithm can see is about ~1 thousand if I recall.

How would you install and use it?
1. Clone the repository:
```
git clone https://github.com/PixEthos/PixAI.git
```
2. Install Golang:
Fedora:
```
dnf install golang
```
Debian:
```
apt install golang
```
Arch:
```
pacman -s golang
```
3. Open the location of where you installed it:
```
cd ~/Documents/PixAI/
```
4. Run: "go build ." and it will build the binary.

Or:
You can run the binary already inside it, usually it's about a week older than the current - but it should work. 

YOU WILL NEED FYNE, you may have issues with development libraries; so it's intended for you to be able to get those development pieces; x11-devel particularly.

NOTE:
This branch is just the 'development branch', there's no generative for the AI as of now.

This is an early, early, early alpha release for those curious about what I had been working on.
