# golsb

Since I wanted to look into steganography for some time,
this is a small project dedicated to steganography.

It provides functionality to **encode** a message in an image and
to **decode** the message from the image.

## Status

[![Build Status](https://travis-ci.com/lzutao/golsb.svg?branch=master)](https://travis-ci.com/lzutao/golsb)

## Build

Clone [this repo](https://github.com/lzutao/godhchat) and go to cloned directory.
Run `go build` to build the program.

If you want Windows Cross Compiling, run on Linux:

```bash
GOOS=windows GOARCH=386 go build
```

## Usage

Help menu:

```bash
% ./golsb
usage: ./golsb [-h] [--encode | --decode] [--message=message] <image.jpg>

Steganography image file

optional arguments:
  -decode
      decode a hidden message from an image
  -encode
      encode a message into an image
  -message string
      message to be encoded to the image

[+] Written by lzutao
```

### Example

#### Encode

```bash
% ./golsb --encode --message=message foo.jpg
```

#### Decode

```bash
% ./golsb --decode foo.jpg

```

## About Steganography

Steganography is the art of hiding a message inside another message.
In this case we will hide a text message inside an image.
An image will most propably go unnotified, not a bunch of people will
suspect a message hidden inside an image.
Steganography is **no means of encryption**, just a way of hiding data inside an image.

If you want to learn about Steganography in detail head over to
[the Wikipedia article](http://en.wikipedia.org/wiki/Steganography).

## Implementation Details

The User chooses an image, the image data is then normalized,
meaning that each RGB value is decremented by one if it is not even.
This is done for every pixel in the image.

Next the message is converted to a binary representation,
8 Bits per character of the message.
This binary representation is then applied to the normalized image, 3 Bit per pixel.
This concludes, that the maximal length of a message hidden in an image is:

    Image Width * Image Height * 3
    ------------------------------
                  8

Since the image was normalized, we now know that an **even** r,
g or b value is **0** and an **uneven** is a **1**.
And this is how the message is decoded back from the image.

## Additional layers of security

As mentioned before, steganography is no means of encryption,
just a way to hide data from plain sight.
But one could, for example, hide a pgp encrypted message inside an image.
So even if the image did not go unnoticed, the message would still only
be readable by the person it was addressed to.

## License

Released under [MIT License](LICENSE)

## THANKS TO

- https://github.com/stylesuxx/steganography