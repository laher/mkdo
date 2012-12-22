mkdo
====
mkdo - create appropriate directories before running a command.

Synopsis
--------

Imagine this scenario, without mkdo:

	 $ mv error.log old/2012-01-01/
	 mv: cannot move `error.log' to `old/2012-01-01/': Not a directory
	 $ # D'Oh!
	 $ mkdir old/2012-01-01
	 mkdir: cannot create directory `old/2012-01-01': No such file or directory
	 $ # D'Oh!
	 $ mkdir -p old/2012-01-01
	 $ mv error.log old/2012-01-01/
	 $ # phew

Now, using mkdo:

 	 $ mkdo mv error.log old/2012-01-01/
	 $ # yay!


TODO
----

 - Parsing of date formats, e.g. mv x.log old-[yyyymmdd]/
 - Attempt to fix piping for Windows
 - Testing on OSX
 - man page, more extensive docs
