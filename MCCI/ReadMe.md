# MCCI

Multi-Channel Colour Image

## Description

*Multi-Channel Colour Image*, or shortly – *MCCI*, is a format for storing 
two-dimensional (2D) coloured images. It is similar to the so-called 
*Planar RGB* format which stores colour channels separately.

> The *Planar RGB* format stores an image's red, green, and blue colour 
> channels as separate two-dimensional planes (arrays) in memory rather than 
> interleaving them pixel by pixel.

According to encyclopedias, such way of storing information is used in various 
important industries: it is used in filters, colour-space pipelines and various 
libraries that process channels independently or share logic with planar `YUV` 
buffers. Also, similar formats are used in so-called Machine Vision systems and 
some image capturing hardware.

In common sense, storing each colour channel separately allows to simulate 
vision of non-human beings and mechanisms. For example, some animals on planet 
Earth use 4-channel vision, some of them see the world only in 2-colour mode. 

The main purpose of this format is storing each colour channel of an image 
separately from other channels. The idea of such format came from old analogue 
hardware, i.e. analogue devices, and other formats popular in cinematography, 
such as the `YCbCr` format.

The main principle used in *MCCI* image format is to give users much more 
freedom than in ordinary formats having the strict `RGB` pixel order. 

Unfortunately, the humanity on planet Earth in its current state still has not 
yet developed a unified standard for storing multichannel 2D images. Some old 
systems use `BGR` pixel order, newer software and hardware use `RGB` pixel 
order, some industries use formats other than `RGB`, such as `YCbCr`. So, a 
decision was made to create a single one unified format.

> And one ring to rule them all. 
*The Lord of the Rings. John Ronald Reuel Tolkien, 1937-1949.* 

## Functionality

Some of the supported functionality is described below.  
While this format is an experiment, it may be changed in future revisions. 

Format provides access to pixels as a 2D canvas, i.e. a two-dimensional array 
of pixels. Numbering of rows starts from the bottom, as opposed to Windows API 
and Golang's `image.Image` interface. So, the first row of pixels is the 
bottom one, and the last row of pixels is the top one. Numbering of pixels in 
a row increases from left to right as in many well known formats. Such 
behaviour was selected specifically to support such industries as science, game 
development and cinematography, where zero point is often located in the bottom 
left corner.

The class constructor, even though *Go* language does not have any classes, 
accepts several important settings of the resulting image:

- Channel pixel format,
- Alpha channel mode.

The *MCCI* format allows you to store pixels with either straight colours or 
with so-called pre-multiplied colours, when colour values are pre-multiplied by 
alpha channel. The first mode allows for better precision, the second mode 
makes calculations faster.

Colour values of each pixel in *MCCI* format can be stored using 8-bit 
integers, 16-bit integers or IEEE 754 32-bit floating point numbers. 8-bit 
colour channels are used to save space on disk while FP numbers enable access 
to advanced features, such as high fidelity, high dynamic range (HDR), 
increased precision and more.

## Future Usage

In the distant future, when photo camera sensors and other detecting hardware 
will become more advanced, this format will allow users to store additional 
information in 2D images, such as colour values in infra-red (IR) range, colour 
values in ultra-violet (UV) range and even much more. 

The format allows users to set textual name and wave length in nanometers (nm) 
for each colour channel. In theory, the format can store information in 
non-optical spectres, such as emission in radio band ranges or radioactive 
radiation in ranges beyond optical.
