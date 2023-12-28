package httphelper

import "log"

// Unfortunately, people on this planet change their rules very often.
// The new document titled as 'RFC 9110 HTTP Semantics' and published in June
// 2022 changed the HTTP rules again. More information can be found here:
// https://www.rfc-editor.org/rfc/rfc9110.html. Some of the important changes
// of this document are following:

/*
	This document updates RFC 3864 and obsoletes RFCs 2818, 7231, 7232, 7233,
	7235, 7538, 7615, 7694, and portions of 7230.
*/

// Pay attention to the section '5.3' named 'Field Order'. It states that
// multiple headers are not allowed unless it is stated in the header's
// definition.

/*
	This means that, aside from the well-known exception noted below, a sender
	MUST NOT generate multiple field lines with the same name in a message
	(whether in the headers or trailers) or append a field line when a field
	line of the same name already exists in the message, unless that field's
	definition allows multiple field line values to be recombined as a
	comma-separated list (i.e., at least one alternative of the field's
	definition allows a comma-separated list, such as an ABNF rule of #(values)
	defined in Section 5.6.1).
*/

// Taking into consideration that an official document now states examples with
// such words as "Foo", "Bar" and "Baz", as in the example of the section 5.2.,
// all the document looks like a comedy parody with sick clowns.

/*
	Example-Field: Foo, Bar
	Example-Field: Baz
*/

// The Hyper-Text Transfer Protocol (HTTP) became a popular self-evolving,
// self-deprecating and self-obsoleting freak show with unreadable and
// controversial fairy rules which ruin security in the World Wide Web and
// provoke usage of D.D.o.S. by accident. We are sure that this anarchy will
// come to an end at some time in the future, but until then the freak show
// continues ...

func logErrorIfSet(err error) {
	if err != nil {
		log.Println(err)
	}
}
