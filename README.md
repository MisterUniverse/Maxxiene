# Maxxiene
Maxx is a personal cli tool I use for productivity on Windows machines

to use maxx:
```POWERSHELL
git clone https://github.com/MisterUniverse/Maxxiene.git

```

Change directories into Maxxiene where the main.go file is located.

```POWERSHELL
cd ./Maxxiene
```

Build the exe
```GO
go build
```

Now you can place that exe where you want to be able to access and use maxx or create you an environment variable.

### Here is an example of how to use Maxxiene:

This command will generate a list of todo items that you have created. If you have none then the list will be empty or have a default value depending on which version you have.
```POWERSHELL
.\maxx list
```

This command will ping an endpoint at the specified and returns a status code.
```POWERSHELL
.\maxx net ping -u 192.168.0.1
```


## Commands for Maxx
- `help`: Help command basically gives you this list of commands
- `browser`: Opens your default browser
- `code`: Command used for generating code templates
- `info`: Information about the current system
- `list`: Displays you todo list and items
- `net`: The net package is a pallette of network commands
- `notes`: Displays a list of files that are your notes
- `server`: Spools up a server of the available type
- `todo`: for adding items to your todo list
- `version`: Prints the version number of Maxxiene



### browser cmd
The browser cmd uses `tabs.txt` located in the ./resources directory. `tabs.txt` is a list of url's you wish the browser command to open every time you call it. If you want to add url endpoints to this list just add them on a new line. Basically your `tabs.txt` file should looks something like this:

```
https://github.com/
https://google.com/
https://example.com/
```

You can add as many as you would like (I think).


### server

***Quickly spool up a local server***

#### Http Usage

With no arguments, `maxx server --http` starts serving files under ./ over port 8080.

```POWERSHELL
$ .\maxx server --http
INFO[0000] Serving ./ over 0.0.0.0:8080... Stop with ^C
```

With `-d` arg, specify the directory to be served.

```POWERSHELL
$ .\maxx --http -d ./templates/
INFO[0000] Serving ./templates/ over 0.0.0.0:8080... Stop with ^C 
```

With `-p` arg, specify the port to serve on.

```POWERSHELL
$ .\maxx --http -p 31415
INFO[0000] Serving ./ over 0.0.0.0:31415... Stop with ^C 
```


## TODO:
- [/] finish boilerplate code sections
    - [X] html section
    - [X] css section
    - [X] js section
    - [X] go section
    - [X] python section
    - [] Clean and Optimize
- [] add command that set's up a work environment set execution policy and install winget first
    git clone https://github.com/MisterUniverse/env_scripts.git
- [] Server command
    - [] http
       - [] need to open server on go routine
    - [] tcp
    - [] udp
    - [] rtp
