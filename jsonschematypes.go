package dsmnd

// See https://json-schema.org/understanding-json-schema/reference/string.html#built-in-formats
type JsonSchemaBuiltinFormat Datum

var JsonSchemaBuiltinFormats = []JsonSchemaBuiltinFormat{}

/*
"date-time": Date and time together, for example, 2018-11-13T20:20:39+00:00.
"time": Time, for example, 20:20:39+00:00
"date": Date, for example, 2018-11-13.
"duration": A duration as defined by the ISO 8601 ABNF for “duration”.
"email": Internet email address, see RFC 5321, section 4.1.2.
"idn-email": The internationalized form of an Internet email
"hostname": Internet host name, see RFC 1123, section 2.1.
"idn-hostname": An internationalized Internet host name,
"ipv4": IPv4 address, according to dotted-quad ABNF syntax
"ipv6": IPv6 address, as defined in RFC 2373, section 2.2.
"uuid": A Universally Unique Identifier as defined by RFC 4122.
"uri": A universal resource identifier (URI), according to RFC3986.
"uri-reference": A URI Reference (either a URI or a relative-reference),
"iri": The internationalized equivalent of a “uri”, according to RFC3987.
"iri-reference": The internationalized equivalent of a “uri-reference”,
"iri-reference") rather than "uri" (or "iri"). "uri" should
*/

/*
The following is the list of formats specified in the JSON Schema specification.

Dates and times

Dates and times are represented in RFC 3339, section 5.6. This is
a subset of the date format also commonly known as ISO8601 format.

"date-time": Date and time together, for example, 2018-11-13T20:20:39+00:00.
"time": Time, for example, 20:20:39+00:00
"date": Date, for example, 2018-11-13.
"duration": A duration as defined by the ISO 8601 ABNF for “duration”.
            For example, P3D expresses a duration of 3 days.

Email addresses

"email": Internet email address, see RFC 5321, section 4.1.2.
"idn-email": The internationalized form of an Internet email
             address, see RFC 6531.

Hostnames

"hostname": Internet host name, see RFC 1123, section 2.1.
"idn-hostname": An internationalized Internet host name,
                see RFC5890, section 2.3.2.3.

IP Addresses

"ipv4": IPv4 address, according to dotted-quad ABNF syntax
        as defined in RFC 2673, section 3.2.
"ipv6": IPv6 address, as defined in RFC 2373, section 2.2.

Resource identifiers

"uuid": A Universally Unique Identifier as defined by RFC 4122.
        Example: 3e4666bf-d5e5-4aa7-b8ce-cefe41c7568a
"uri": A universal resource identifier (URI), according to RFC3986.
"uri-reference": A URI Reference (either a URI or a relative-reference),
                 according to RFC3986, section 4.1.
"iri": The internationalized equivalent of a “uri”, according to RFC3987.
"iri-reference": The internationalized equivalent of a “uri-reference”,
                 according to RFC3987

If the values in the schema have the ability to be relative
to a particular source path (such as a link from a webpage),
it is generally better practice to use "uri-reference" (or
"iri-reference") rather than "uri" (or "iri"). "uri" should
only be used when the path must be absolute.

*/
