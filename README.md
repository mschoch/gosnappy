## gosnappy

A thin, not very well tested go wrapper around the C implementation of snappy.

## why?

There is a fine native port of snappy to go - [snappy-go](https://code.google.com/p/snappy-go/)

Yes, but byte streams produced by the snappy-go encoder are not the same as those produced by the C implementation.  For another project, I want to directly compare bytes produced by the C implementation.  This wrapper seemed the quickest way to accomplish that.