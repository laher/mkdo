mkdo
====
mkdo - create appropriate directories then run a given command.

Synopsis
--------

Much like 'sudo', mkdo is prepended to an existing command to make it work better.

mkdo can be prepended to an existing command to create directories mentioned in the command.

e.g.

         mkdo mv error.log logs/old/2013-01-01/

This would create the directory "logs/old/2013-01-01/" before running the 'mv' command.

Note that mkdo uses trailing slashes (or backslashes in Windows) to identify folder names.

Downloads
---------
[Latest binaries](http://laher.github.com/mkdo/dl/latest/) for Linux, Mac, Windows.

Lame joke
---------
With apologies, mkdo comes with a lame rip-off of the [xkcd sudo joke](http://xkcd.com/149/):

 - child: put that cake in my lunchbox/
 - parent: What?! You haven't even got a lunchbox/
 - child: *mkdo* put that cake in my lunchbox/
 - parent: okay

To do
-----

 - Parsing of date formats, e.g. mv x.log old-[yyyymmdd]/
 - Attempt to fix piping for Windows
 - Testing on OSX
 - more extensive docs
 - get a better joke
