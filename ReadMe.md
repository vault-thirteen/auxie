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
6. **CSV** – [CSV](CSV/ReadMe.md) – A writer for _CSV_ format.
7. **dev-inf** – [Dev Inf](dev-inf/ReadMe.md) – Various development infrastructure scripts.
8. **endianness** – [endianness](endianness/ReadMe.md) – Detection of "endianness" of a machine.
9. **Env** – [Env](env/ReadMe.md) – A simple library to work with environment variables.
10. **Errors** – [Errors](errors/ReadMe.md) – A simple library to work with built-in errors.
11. **File** – [File](file/ReadMe.md) – A simple library to work with files and folders.
12. **hash** – [hash](hash/ReadMe.md) – Helper functions for hash sums.
13. **Header** – [Header](header/ReadMe.md) – _IANA_ registered message header names.
14. **HTML-DOM** – [HTML-DOM](HTML-DOM/ReadMe.md) – A library to work with _HTML_ document object model.
15. **HTTP Helper** – [HTTP Helper](http-helper/ReadMe.md) – A package that helps in processing and testing _HTTP_ requests.
16. **IPA** – [IPA](IPA/ReadMe.md) – Methods to work with Internet Protocol Addresses.
17. **MemUC** – [MemUC](MemUC/ReadMe.md) – A library to cope with memory usage in _Go_ programming language.
18. **MIME** – [MIME](MIME/ReadMe.md) – MIME types and basic functions to work with them.
19. **MySQL-IT** – [MySQL-IT](MySQL-IT/ReadMe.md) – A template for an application importing data into _MySQL_ database.
20. **NTS** – [NTS](NTS/ReadMe.md) – Methods for processing null-terminated strings.
21. **Number** – [Number](number/ReadMe.md) – A simple library to work with numbers in _Go_ programming language.
22. **Random** – [Random](random/ReadMe.md) – A simple library to work with random integer numbers.
23. **Range** – [Range](range/ReadMe.md) – Methods to work with ranges of floating point numbers.
24. **Reader** – [Reader](reader/ReadMe.md) – Auxiliary methods and functions for reading.
25. **ReaderSeeker** – [ReaderSeeker](ReaderSeeker/ReadMe.md) – An interface which combines _io.Reader_ and _io.Seeker_ interfaces.
26. **RPoFS** – [RPoFS](rpofs/ReadMe.md) – Random Password of Fixed Size.
27. **RS** – [RS](rs/ReadMe.md) – A Seeker extension to the reader package.
28. **SAAC** – [SAAC](SAAC/ReadMe.md) – Simple Arithmetic Average Calculator.
29. **SLReader** – [SLReader](SLReader/ReadMe.md) – A reader based on io.Reader and using speed limits.
30. **SMA** – [SMA](SMA/ReadMe.md) – Simple Moving Average.
31. **SMS** – [SMS](SMS/ReadMe.md) – Simple Merge Sorter.
32. **SQL** – [SQL](SQL/ReadMe.md) – Various _SQL_ helpers.
33. **SSC** – [SSC](SSC/ReadMe.md) – Simple Stupid Cache.
34. **Tester** – [Tester](tester/ReadMe.md) – Various useful functions to perform tests.
35. **Time** – [Time](time/ReadMe.md) – A simple library to work with time in _Go_ programming language.
36. **Time Metadata** – [Time Metadata](time-metadata/ReadMe.md) – An auxiliary type for storing time of creation and time of update.
37. **TSB** – [TSB](TSB/ReadMe.md) – Tri-State Boolean.
38. **Unicode** – [Unicode](unicode/ReadMe.md) – _Unicode_ symbols processing package.
39. **VCS** – [VCS](VCS/ReadMe.md) – A collection of helpers for version control systems.
40. **Versioneer** – [Versioneer](Versioneer/ReadMe.md) – Product and dependency version helper.
41. **Zipper** – [Zipper](zipper/ReadMe.md) – Functions to work with compressed files.

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
