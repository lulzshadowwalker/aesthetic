# a e s t h e t i c 

## Table of Contents
- [Demonstration](#demonstration)
  - [Images](#images)
  - [Videos](#videos)
- [CLI Flags](#cli-flags)

## Demonstration

### Images
📷 [Source](https://www.pinterest.com/pin/847802698626328951/)

<div style="display: grid; grid-template-columns: 1fr;">
  <img src="https://imgur.com/kUZOwhL.jpg" alt="source image" style="width: 96%;">
</div>

<div style="display: grid; grid-template-columns: 1fr 1fr;">
  <img src="https://imgur.com/TT1aGmX.png" alt="Image 1" style="width: 48%;">
  <img src="https://imgur.com/YOQAxXK.png" alt="Image 2" style="width: 48%;">
</div>

<div style="display: grid; grid-template-columns: 1fr 1fr;">
  <img src="https://imgur.com/6mGSGv8.png" alt="Image 3" style="width: 48%;">
  <img src="https://imgur.com/u21Xg9y.png" alt="Image 4" style="width: 48%;">
</div>

### Videos
🎬 [orignal creator](https://www.tiktok.com/@asethetic.check15)<br>

[**Imgur**](https://imgur.com/a/CmH1nov)


## CLI Flags
- src `string` **mandatory** <br>
*Source image or video to be converted. mandatory<br>
Currently supports most image formats and mp4 for video sources*
- dest `string`<br>
*Destination of the converted image or video <br>
(default desktop)*
- dot-size `int`<br>
*the upper bound for the dot size<br>
(default 4)*
- grayscale `bool`<br>
*Decides whether the result should be grayscale (black & white) or preserve the original colors of the image (default false)*
- sample-interval `int`<br>
*Sampling rate for the given image or video<br>
Dictates how many pixels are going to be read<br>
Essentially corresponds to the gap or noise size (higher sampling rate, less dots or info to represent the image, larger gaps)<br>
(default 5)*
- white-noise `bool` <br>
*uhm.. white noise <br>
the gap or background color<br>
(default false)*
- help <br>
*Lists all commands*