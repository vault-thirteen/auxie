# Auxie

A collection of auxiliary types, methods, functions and other small useful 
things.  

The name "**Auxie**" comes from the short name "**AUX**" which in its turn comes from a 
full word "**Auxiliary**". For historical reasons, the word "**AUX**" is forbidden to 
be used as a folder name in operating systems based on MS-DOS and Windows.  

## List of included libraries

1. **Array Search** – [as](as/ReadMe.md) – Various methods for searching in arrays.
2. **Basic Types** – [BasicTypes](BasicTypes/ReadMe.md) – A collection of common sense basic types.
3. **Bit** – [Bit](bit/ReadMe.md) – A simple library to work with bits in Go programming language.
4. **BOM** – [BOM](BOM/ReadMe.md) – Functions and methods to work with Unicode byte order mark.
5. **Boolean** – [Boolean](boolean/ReadMe.md) – A boolean-to-string parser.
6. **Env** – [Env](env/ReadMe.md) – A simple library to work with environment variables.
7. **File** – [File](file/ReadMe.md) – A simple library to work with files and folders.
8. **HTTP Helper** – [HTTP Helper](http-helper/ReadMe.md) – A package that helps in processing and testing HTTP requests.
9. **IPA** – [IPA](IPA/ReadMe.md) – Methods to work with Internet Protocol Addresses.
10. **NTS** – [NTS](NTS/ReadMe.md) – Methods for processing null-terminated strings.
11. **Number** – [Number](number/ReadMe.md) – A simple library to work with numbers in Go programming language.
12. **Random** – [Random](random/ReadMe.md) – A simple library to work with random integer numbers.
13. **Range** – [Range](range/ReadMe.md) – Methods to work with ranges of floating point numbers.
14. **Reader** – [Reader](reader/ReadMe.md) – Auxiliary methods and functions for reading.
15. **ReaderSeeker** – [ReaderSeeker](ReaderSeeker/ReadMe.md) – An interface which combines io.Reader and io.Seeker interfaces.
16. **RPoFS** – [RPoFS](rpofs/ReadMe.md) – Random Password of Fixed Size.
17. **RS** – [RS](rs/ReadMe.md) – A Seeker extension to the reader package.
18. **SAAC** – [SAAC](SAAC/ReadMe.md) – Simple Arithmetic Average Calculator.
19. **SLReader** – [SLReader](SLReader/ReadMe.md) – A reader based on io.Reader and using speed limits.
20. **SMA** – [SMA](SMA/ReadMe.md) – Simple Moving Average.
21. **SMS** – [SMS](SMS/ReadMe.md) – Simple Merge Sorter.
22. **Time** – [Time](time/ReadMe.md) – A simple library to work with time in Go programming language.
23. **Time Metadata** – [Time Metadata](time-metadata/ReadMe.md) – An auxiliary type for storing time of creation and time of update.
24. **TSB** – [TSB](TSB/ReadMe.md) – Tri-State Boolean.
25. **Unicode** – [Unicode](unicode/ReadMe.md) – Unicode symbols processing package.

## Go language version

Code in this repository requires Go language version 1.20 or newer.

### Microsoft Windows OS support

Go language 1.20 is the last version supported on old Windows operating systems.

As stated in release notes of Go 1.20 located at 
https://go.dev/doc/go1.20#windows, 
> Go 1.20 is the last release that will run on any release of Windows 7, 8, 
> Server 2008 and Server 2012.

As stated in release notes of Go 1.21 located at 
https://go.dev/doc/go1.21#windows,
> Go 1.21 requires at least Windows 10 or Windows Server 2016; support for 
> previous versions has been discontinued.

Go language 1.21 provides some breaking changes, so do not rush into upgrading 
to it at any cost.
