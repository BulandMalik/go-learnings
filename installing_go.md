## Setting up Environment with Go, Its Tools and IDE

### Install Go
The first step is install in the Go development environment. You don't need a run time to deploy Go applications, but you do need to download the development environment to write Go programs. 

Download and setup in your specific environment.
https://golang.org/dl

## Creating `go` Directory
After the installation is done we need to create a directory to hold our work. One of the things about Go that it's different from many other languages is that Go expect all of your Go code, but the code you write and the third-party libraries that you download, to be located within a single directory called the Go path. By default the Go path is a directory named Go, `all lowercase`, in your home directory. 

## Setting Up Git
download and setup `git` if not exists.

## IDE
We will download Microsoft Visual Studio code as its an excellent text editor that is a Go plugin that's surprisingly good.

Microsoft has actually been very active in the Go community, hiring some of the best and brightest work on Kubernetes for Azure.

Visual Studio code runs on Mac, Windows, and Linux, and you can download it from code dot Visual Studio com.

## Visual Studio Code Go Extension
We get to the Go extension by clicking on the extension icon on the left side and typing in Go into the search box. The first item in the list is the Go extension. We then click install and then click reload. That loads the extension into our current window in Visual Studio code. We have one final step left. The Go extension for VS code depends on several open source Go tools to provide formatting, linting, documentation parsing, and other nice features that you expect in a modern programming environments. Before we can install anything you should Go and check under preferences, settings, and make sure under user settings you have the line `Go DUT infer path` true. That will make sure that Visual Studio code gets your Go path correctly from its environments. 

Once you make sure that's there, press f1 to open up the command palette or type in command `shift P and typing Go` : install update tools. As you will see, you don't the type the whole command in to have it pop up at the list. Once you do so the open-source tools will install. Once it is done your environment is ready. 


