### TEST BRANTCH! ###

IF YOU INSTALL THIS:
- There is a limit for the input values to be 255 in length, it's a hard limit at the moment.
- Working on the Markov chains, but it still is fun to see what it is capable of doing.
- Matrix handling is still a work in progress.
- Inputs _ARE_ encrypted, eventually I plan on making sure that logged data is only the outputs, and not the input values. Course, this is the dev branch, so it's ideal to get an idea.
- Prediction values are basically based on the number / length of the input value. That means you only get a number at this point in time for it.
- The number of matches the current algorithm can see is about ~1 thousand if I recall.

Note:
- Still working on the GRU's, and it's going to take a while to truly make a fully context heavy piece in the algorithm. Probably going to have to use mapping eventually. I am using switch-case, which uses converted arrays into strings, and than it matches them with the input (not too effective, but it's early alpha)
- Still working on the outputs
- Still working on the matching/prediction

At the moment, it's not doing too badly. This is 4 months of work later, considering the entire algorithm is FOSS, and is designed to run entirely locally on the machine – I do have counter measures to ensure idiot devs don't have an unsecure way to gather input/output information, particularly from the algorithm (as the algorithm encodes into ASCII, and Hex. In fact, the Hex is separated, but used for logging. The ASCII is being used for the GRU's for encoded inputs)

This algorithm is around 3k lines by now, but I keep iterating.

- "Why is it taking so long?"
```
I am using the standard library in Go, while Fyne is the UI.
```

- "What is the purpose of the project?"
```
This project is more or less a passion project; to see if I can do it, and to unravel a lot of the intricate details of the way algorithms work.
```

- "Why choose Go?"
```
I previously used C, and before that, Python. But Python relied too heavily on thirdparty libraries, and the syntax of Python is...I'll just be honest, dreadful. I might write pieces in C, but if that _is_ going to be the case, it's to ensure that the algorithm can run a bit more smoothly where it needs to. Go has the benefit of garbage collection; so it's genuinely for that benefit. That and it's technically more modern, meaning that Go has the benefit of more modern technologies (ie: Generics, and a lack of semi-colons for marking).
```

- "Why not use (x) library?"
```
Even if it speeds up development; I prefer to understand my algorithm and applications; the core of each application I have written (most are not published, but still) is usually in the standard library of the language I use. This is intentional; because ease != simplicity.
```

That should clear up the questions.

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
