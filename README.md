# IconGen
IconGen can be used to Generate icons of different sizes or reduce image size.

## Compatibility
go1.11

## Getting Started

### Download Release

- Download the release [here](https://github.com/xuzhuoxi/IconGen/releases).

- Download the repository:

	```sh
	go get -u github.com/xuzhuoxi/IconGen
	```
	
	This will retrieve the library.

### Build

Execution the construction file([build.sh](/build/build.sh)) to get the releases if you have already downloaded the repository.

You can modify the construction file([build.sh](/build/build.sh)) to achieve what you want if necessary. The command line description is [here](https://github.com/laher/goxc).

## Run

### Demo

[Here](/demo/win) is a running demo for windows 64bit platform.

[Here](/demo/macOS) is a running demo for MacOS platform.

The running command is consistent of all platforms.

Goto <a href="#command-line">Command Line Description</a>.

### Command Line

Supportted command line parameters as follow:

| -       | -            | -                                                            |
| :------ | :----------- | ------------------------------------------------------------ |
| -size   | **required** | The size of the generated image.                             |
| -base   | optional     | Custom base running folder for each path in the command. Use Execution file directory if no setting. |
| -in     | optional     | Custom source folder or file. Use -base value if no setting. |
| -inFolder     | optional     | true: -in is folder path. false: -in is file path. |
| -out    | optional     | Custom output folder. Use -base value if no setting.         |
| -outFolder     | optional     | true: -out is folder path. false: -out is file path. |
| -format | optional     | The format of the generated image. Supported as follows: png, jpg, jpeg, jps |
| -ratio  | optional     | The quality of the generated image. Supported for jpg,jpeg,jps. |

E.g.:

-size=128,256

-size=128,256x256

-base=/

-in=./source

-in=./source/icon.png

-inFolder=true

-out=c:/image/ouput

-out=c:/image/ouput/b.jpeg

-outFolder=false

-format=jpeg

-format=jpg,png

-ratio=70

## Related Library

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## Contact

xuzhuoxi 

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## License
IconGen source code is available under the MIT [License](/LICENSE).


