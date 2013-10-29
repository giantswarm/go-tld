Go-tld
======

Package tld provides a way to verify that a given top level domain name is
valid. It comes with a predefined list however it can be updated at runtime by
running tld.Update(url) where url points to a text file containing a list of
acceptable TLDs. tld.IANA contains the URL to get the IANA database which should
be sufficient for most users.
