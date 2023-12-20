# Auxie

A collection of auxiliary types, methods, functions and other small useful 
things.  

The name "**Auxie**" comes from the short name "**AUX**" which in its turn comes 
from a full word "**Auxiliary**". For historical reasons, the word "**AUX**" 
is forbidden for usage as a folder name in operating systems based on _MS-DOS_ 
and _Windows_.  

## List of included libraries

1. **Array Search** – [as](as/ReadMe.md) – Various methods for searching in arrays.
2. **Basic Types** – [BasicTypes](BasicTypes/ReadMe.md) – A collection of common sense basic types.
3. **Bit** – [Bit](bit/ReadMe.md) – A simple library to work with bits in _Go_ programming language.
4. **BOM** – [BOM](BOM/ReadMe.md) – Functions and methods to work with _Unicode_ byte order mark.
5. **Boolean** – [Boolean](boolean/ReadMe.md) – A boolean-to-string parser.
6. **Env** – [Env](env/ReadMe.md) – A simple library to work with environment variables.
7. **Errors** – [Errors](errors/ReadMe.md) – A simple library to work with built-in errors.
8. **File** – [File](file/ReadMe.md) – A simple library to work with files and folders.
9. **Header** – [Header](header/ReadMe.md) – _IANA_ registered message header names.
10. **HTTP Helper** – [HTTP Helper](http-helper/ReadMe.md) – A package that helps in processing and testing _HTTP_ requests.
11. **IPA** – [IPA](IPA/ReadMe.md) – Methods to work with Internet Protocol Addresses.
12. **NTS** – [NTS](NTS/ReadMe.md) – Methods for processing null-terminated strings.
13. **Number** – [Number](number/ReadMe.md) – A simple library to work with numbers in _Go_ programming language.
14. **Random** – [Random](random/ReadMe.md) – A simple library to work with random integer numbers.
15. **Range** – [Range](range/ReadMe.md) – Methods to work with ranges of floating point numbers.
16. **Reader** – [Reader](reader/ReadMe.md) – Auxiliary methods and functions for reading.
17. **ReaderSeeker** – [ReaderSeeker](ReaderSeeker/ReadMe.md) – An interface which combines _io.Reader_ and _io.Seeker_ interfaces.
18. **RPoFS** – [RPoFS](rpofs/ReadMe.md) – Random Password of Fixed Size.
19. **RS** – [RS](rs/ReadMe.md) – A Seeker extension to the reader package.
20. **SAAC** – [SAAC](SAAC/ReadMe.md) – Simple Arithmetic Average Calculator.
21. **SLReader** – [SLReader](SLReader/ReadMe.md) – A reader based on io.Reader and using speed limits.
22. **SMA** – [SMA](SMA/ReadMe.md) – Simple Moving Average.
23. **SMS** – [SMS](SMS/ReadMe.md) – Simple Merge Sorter.
24. **Tester** – [Tester](tester/ReadMe.md) – Various useful functions to perform tests.
25. **Time** – [Time](time/ReadMe.md) – A simple library to work with time in _Go_ programming language.
26. **Time Metadata** – [Time Metadata](time-metadata/ReadMe.md) – An auxiliary type for storing time of creation and time of update.
27. **TSB** – [TSB](TSB/ReadMe.md) – Tri-State Boolean.
28. **Unicode** – [Unicode](unicode/ReadMe.md) – _Unicode_ symbols processing package.
29. **VCS** – [VCS](VCS/ReadMe.md) – A collection of helpers for version control systems.
30. **Versioneer** – [Versioneer](Versioneer/ReadMe.md) – Product and dependency version helper.
31. **Zipper** – [Zipper](zipper/ReadMe.md) – Functions to work with compressed files.

## Go language version

Code in this repository requires _Go_ language version 1.20 or newer.

### Microsoft Windows OS support

_Go_ language 1.20 is the last version supported on old _Windows_ operating 
systems.

As stated in release notes of _Go_ 1.20 located at 
https://go.dev/doc/go1.20#windows, 
> Go 1.20 is the last release that will run on any release of Windows 7, 8, 
> Server 2008 and Server 2012.

As stated in release notes of _Go_ 1.21 located at 
https://go.dev/doc/go1.21#windows,
> Go 1.21 requires at least Windows 10 or Windows Server 2016; support for 
> previous versions has been discontinued.

_Go_ language 1.21 provides some breaking changes, so do not rush into upgrading 
to it at any cost.
