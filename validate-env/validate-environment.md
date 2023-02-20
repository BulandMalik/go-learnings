## Validating Environment

Run the sample program.

Visual Studio code has a terminal built into it which we can launch with typing ctrl shift back tick, or by opening the command palette and typing in terminal, create new integrated terminal. 

Once the terminal is opened you can type in `Go run hello Buland`, and you should see HelloWorld printed on the next line. The Go command has lots of helpful sub commands, including fmt and format your code, run that runs a specified file, and build which creates a binary. 

If you want to see Go build a binary for you type and `Go build hello.go`, hit return. You'll then see a file created in your directory called hello. You can run this file by typing in `./hello`. 

> Congratulations, you've made your first Go binary, you know have a completely working Go installation on your computer. 

## Go Playground

http://play.golang.org

It can generate a unique URL that allows you to share your code snippet with other people. Click the share button and a text field will pop up next to it with a URL. You can copy this URL and save it for later, and you can share it with other people, or if you just want to come back to your idea later you can enter in that URL again and there's your program.

You might have to set `go env -w GO111MODULE=auto`

> GO111MODULE=off forces Go to behave the GOPATH way, even outside of GOPATH. 
