<h2>Hello, I'm Berik <i>Alanapapa</i> Bazarov!</h2>
<img align='right' src="gopher-rainbow.gif" width="230">
<p><em>Freelancer and student at <a href="https://alem.school/" target="_blank">Alem Coding School</a><img src="https://media.giphy.com/media/WUlplcMpOCEmTGBtBW/giphy.gif" width="30"> 
</em></p>

[![Github: alanapapa](https://img.icons8.com/material-outlined/48/000000/github.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.github.com/alanapapa/)](https://www.github.com/alanapapa/)
[![Instagram: berikbazarov](https://img.icons8.com/fluent/48/000000/instagram-new.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.instagram.com/berikbazarov/)](https://www.instagram.com/berikbazarov/)
[![Facebook: berikbazarov](https://img.icons8.com/color/48/000000/facebook.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.facebook.com/bazarovberik/)](https://www.facebook.com/bazarovberik/)

![GitHub followers](https://img.shields.io/github/followers/alanapapa?label=Follow&style=social)


## ***ascii-art***
### This is my solution to the *ascii-art* educational task on [01 Edu System](https://www.01-edu.org/)

<br>


### Objectives

Ascii-art consists on receiving a `string` as an argument and outputting the `string` in a graphic representation of ASCII.

- This project should handle numbers, letters, spaces, special characters and `\n`.
- Take a look at the ASCII manual.

This project will help you learn about :

- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Manipulation of structures.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code present a **test file**.

- It will be given some [**banner**](https://github.com/01-edu/public/blob/master/subjects/ascii-art) files with a specific graphical template representation of ASCII. The files are formatted in a way that it is not necessary to change them.

### Banner Format

- Each character has an height of 8 lines.
- Characters are separated by a new line `\n`.
- Here is an example of ' ', '!' and '"'(one dot represents one space) :

```console

......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......

etc
```

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello"
  _              _   _
 | |            | | | |
 | |__     ___  | | | |   ___
 |  _ \   / _ \ | | | |  / _ \
 | | | | |  __/ | | | | | (_) |
 |_| |_|  \___| |_| |_|  \___/


student@ubuntu:~/ascii-art$ ./ascii-art "HeLlO"
  _    _          _        _    ____
 | |  | |        | |      | |  / __ \
 | |__| |   ___  | |      | | | |  | |
 |  __  |  / _ \ | |      | | | |  | |
 | |  | | |  __/ | |____  | | | |__| |
 |_|  |_|  \___| |______| |_|  \____/


student@ubuntu:~/ascii-art$
```
