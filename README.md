# alto
A speedy, light package manager that's ready for the masses

## Installing
There are 3 recommended ways of installing alto.
### Using the included shellfile (recommended)
Simply run the following command.
```sh
curl -sL https://raw.githubusercontent.com/altopm/alto/master/install.sh | sudo sh
```
### Using the prebuilt binary
Go to the [installs page](github.com/altopm/alto/releases) and download the latest binary.
What you need to do next depends on your distro.
### Building the binary yourself
You can build the binary yourself. You **must** have [go](https://golang.org) installed. Then clone and change into the `alto` directory:
```sh
git clone https://github.com/altopm/alto.git && cd alto
```
Then run the following command:
```sh
go build
```
We strongly recommend that you then add this to your $PATH environment variable so that you can run `alto` from anywhere.
## Usage
### Installing packages
To install a package, run the following command:
```sh
sudo alto install <package>
```
