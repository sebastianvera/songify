# Songify

I wanted to display spotify's current song while streaming
(using OBS) by using a web page and the *CLR Browser* plugin
or *BrowserSource*, but I didn't found either a plugin or software
to do this. This is **under development**, since I built it in a couple of hours.

The server uses websockets to keep the current song updated on
the web page.

> It currently work for **mac only**.

## Usage

You need to have go installed.

- Clone the repo.
- `cd` into it.
- Run `go get`.
- Run `make` to start the program.

This will mount a web server on `localhost:1616`, where you'll be able
to see the current song.

> To change the port you can use the `-port` flag.

----

### TODO

- [ ] Prettify the web page (**I didn't add any css**).
- [ ] Add Angular 2 UI version (just for fun).
- [ ] Add React UI version (just for fun).
- [ ] Maybe add support to other platforms?
- [ ] Refactor server code.

