
# ASN.1

BER - Basic Encoding Rules
DER - Distinguished Encoding Rules
PER - Packed Encoding Rules
XER - XML Encoding Rules

## Specifications

* Abstract Syntax Notation One (ASN.1):
  Specification of basic notation

  https://www.itu.int/ITU-T/studygroups/com17/languages/X.680-0207.pdf

* ASN.1 encoding rules:
  Specification of Basic Encoding Rules (BER),
  Canonical Encoding Rules (CER) and
  Distinguished Encoding Rules (DER)

  https://www.itu.int/ITU-T/studygroups/com17/languages/X.690-0207.pdf

* ASN.1 encoding rules:
  Specification of Packed Encoding Rules (PER)

  https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf
  - Basic PER, aligned        **
  - Basic PER, unaligned
  - Canonical PER, aligned
  - Canonical PER, unaligned

* ASN.1 encoding rules:
  XML Encoding Rules (XER)

  https://www.itu.int/ITU-T/studygroups/com17/languages/X.693-0112.pdf


## Readables

https://www.w3.org/Protocols/HTTP-NG/asn1.html


## Summary/Hints

* BER
  - TLV encoded
  - Tag: 2 bits: class
         1 bit : simple type?
         5 bits: tag number, or all set to indicate more tag-byte
  - Length: Definite: <128 single byte, high bit indicates length of length (following bytes)
            Indefinite: 0x10, object ends with two zero bytes
* PER
  - Only generates tags in unions (CHOICE)
  - Only generates length when size can vary, but as compact as possible
  - Optional: single bit in the beginning to show if optional item is present
  - Aligned PER: strings are aligned

## Examples / Libs

* encoding/asn1 (Go)

  https://golang.org/pkg/encoding/asn1/

  - DER-encoded ASN.1 data structures, as defined in ITU-T Rec X.690.


* BER and DER encoding (Go)

  https://github.com/Logicalis/asn1/

  - DER
  - BER encoding and decoding
  - ASN.1 CHOICE types.


* pyasn1 (Python)

  https://github.com/etingof/pyasn1

  - Generic ASN.1 types X.208 (superseded with X.680-683)
  - BER/CER/DER codecs

* PER encoder/decoder (Python)

  https://github.com/cartermc24/PyASN1-PER-Encoder

  - unaligned PER


* ANS.1 Compiler

  https://github.com/vlm/asn1c

  See:
  * Support parsing Information Object and Information Object Set:
    https://github.com/vlm/asn1c/pull/154
  * Patches supporting S1AP (not delivered):
    https://github.com/vlm/asn1c/issues/111
    https://github.com/vlm/asn1c/pull/115
    https://github.com/vlm/asn1c/pull/226
    https://github.com/vlm/asn1c/pull/234
    https://github.com/vlm/asn1c/pull/238

  * Forks handling APER used in S1AP:
    https://github.com/mouse07410/asn1c
    https://github.com/AuthenticEshkinKot/asn1c

# NGAP 38.413

https://portal.3gpp.org/desktopmodules/Specifications/SpecificationDetails.aspx?specificationId=3223

- X.691
- X.680
- X.681

- IEs shall be ordered as in object set definitions.

9.5	Message Transfer Syntax
  - Not defined, but 3GPP TS 36.413:

    "S1AP shall use the ASN.1 Basic Packed Encoding Rules (BASIC-PER)
    Aligned Variant as transfer syntax as specified in ITU-T Rec. X.691 [4]"

  ** Basic PER, aligned **
