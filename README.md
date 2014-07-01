Burner API
----------

Host a file for one download, then burn it.

### Installation

- Install [go](http://golang.org) on your server
- Install [git](http://git-scm.org) on your server
- Clone this repo down onto your server
- Install [goop](https://github.com/nitrous-io/goop) with `go get github.com/nitrous-io/goop`
- Run `goop install` to install all dependencies
- Run the app on port 80 with `goop go run server.go -p 80` (`p` flag is whatever port you want)
- ???
- Profit!

### Usage

- Make a `POST` to `/new` with a file, get back a unique id
- Make any request to `/that-id-you-got-back`, and get the file back
- Try it again and you get a 404 because the file got [baleeted](http://cl.ly/QR8M/baleete.gif), what up

### Some Notes

This is a small simple tool written for a specific purpose, but if you'd like to use it for anything you are welcome to (yay open source!). At the moment, this is API-only, meaning there is no app or web interface anywhere for it. It's like a super super nerdy version of snapchat.

This is my first project using go, and my first time writing any go code, so there are certainly some things that could be improved with the codebase, although I have enjoyed the learning process. There also are almost certainly some security issues with the implementation. Things that would make it more stable and secure include:

- Auto deleting the file if it isn't downloaded after X hours to prevent wasted space
- Not allowing uploads larger than X megabytes
- Throttling requests so that it can't be DDOS'ed

Perhaps someday I will implement these. Or you can do it and send a pull request if you want. That would be sweet.

### License & Contributing

- Licensed under [MIT](LICENSE.md)
- If you want to contribute, feel free to send in a PR homie
