# /ascii-art 🇦

## Presentation

/ascii-art is a program that takes a string as an argument and returns a graphical representation of characters in ASCII art. It is written in GO as it follows the /challenge-go and the /go-reloaded modules. It is currently optimized to run on systems running Linux Ubuntu. The project has been worked on collectively using Git.

> Anatole, Florian, William

## Usage

First, start by building the GO files using the `go build` command.

```
student@ubuntu:~/ascii-art$ go build options/classic.go
```

Then, you can freely run the program.

```
student@ubuntu:~/ascii-art$ ./classic "Classic"
```

## Options

There are currently five different options  for /ascii-art : *Classic*, *Fs*, *Color*, *Output* and *Justify*. Feel free to try them out!

### Classic

```
student@ubuntu:~/ascii-art$ ./classic "Classic"
  _____   _                       _
 / ____| | |                     (_)
| |      | |   __ _   ___   ___   _    ___
| |      | |  / _` | / __| / __| | |  / __|
| |____  | | | (_| | \__ \ \__ \ | | | (__
 \_____| |_|  \__,_| |___/ |___/ |_|  \___|


student@ubuntu:~/ascii-art$
```

### Fs

```
student@ubuntu:~/ascii-art$ ./fs "Fs" shadow

_|_|_|_|
_|         _|_|_|
_|_|_|   _|_|
_|           _|_|
_|       _|_|_|


student@ubuntu:~/ascii-art$
```

### Color

```
student@ubuntu:~/ascii-art$ ./color "Color" --color=black
  _____           _
 / ____|         | |
| |        ___   | |   ___    _ __
| |       / _ \  | |  / _ \  | '__|
| |____  | (_) | | | | (_) | | |
 \_____|  \___/  |_|  \___/  |_|


student@ubuntu:~/ascii-art$
```

### Output

```
student@ubuntu:~/ascii-art$ ./output "Output" standard --output=output.txt
student@ubuntu:~/ascii-art$ cat output.txt
  ____            _                     _
 / __ \          | |                   | |
| |  | |  _   _  | |_   _ __    _   _  | |_
| |  | | | | | | | __| | '_ \  | | | | | __|
| |__| | | |_| | \ |_  | |_) | | |_| | \ |_
 \____/   \__,_|  \__| | .__/   \__,_|  \__|
                       | |
                       |_|
student@ubuntu:~/ascii-art$
```

### Justify

```
student@ubuntu:~/ascii-art$ ./justify "Justify" standard --align=right
                                                     _                 _     _    __          
                                                    | |               | |   (_)  / _|         
                                                    | |  _   _   ___  | |_   _  | |_   _   _  
                                                _   | | | | | | / __| | __| | | |  _| | | | | 
                                               | |__| | | |_| | \__ \ \ |_  | | | |   | |_| | 
                                                \____/   \__,_| |___/  \__| |_| |_|    \__, | 
                                                                                       __/ /  
                                                                                      |___/   

student@ubuntu:~/ascii-art$
```
