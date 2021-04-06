# Go_PyProj
Creates a simple folder with a starter file, .gitignore and README.md

## Installation

### Windows
1. If you're on Windows, a `.exe` file can be found in [Releases](https://github.com/oliverchen415/go_pyproj/releases). Download it to any folder, and copy the path, e.g. `C:\Users\<username>\Documents\go_pyproj`.
2. Open the Start menu and search for "environment variables". Click on "Edit the system environment variables", then "Environment Variables".
3. In "User Variables for <username>", scroll down to path and click "Edit...".
4. Click "New" and paste the path. Click "OK" until everything's closed.

Restart any terminals you have open. Now you can simply call `go_pyproj` from any location on your computer.

### MacOS/Linux
I don't own a Mac or a Linux machine (I have some issues testing out the created binaries), so you'll need to build the binaries yourself.

You'll only need `Go 1.13+`. Run `go build` on the source code to produce the correct binary for your OS. Just like the Windows instructions above, add the binary to your $PATH variable. You can look up the specifics, but in either case you'll need to edit a config file, e.g. `~/bashrc`.

### go get
Alternatively, if you already have Go, you can run
```
go get github.com/oliverchen415/go_pyproj
```
This will download the files to your $GOPATH. On my system, I can find the files at
```
$GOPATH\pkg\mod\github.com\oliverchen415\...
```

## Usage
From a terminal, `cd` into an approprate location. Then enter
```
go_pyproj -name=<project_name> -type=<file_type>
```

You can leave either one blank or both blank. By default, `<project_name>` is "dummyProj" and `<file_type>` is "py". This can also be seen if you enter
```
go_pyproj -h
```

**Do not add an extra dot before the file type.**

This will create a folder with the project name, a file of the appropriate type with the same project name, a README.md file, and a .gitignore. The tool will not create a folder if you already have a folder of the same name; it will exit out immediately.

The file structure looks like this:
```
<project_name>
├── .gitignore
├── <project_name>.<file_type>
└── README.md
```
