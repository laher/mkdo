mkdo
====

[mkdo](http://www.laher.net.nz/mkdo) - create appropriate directories then run a given command.

Synopsis
--------

Much like 'sudo', mkdo is prepended to an existing command to make it work better.

mkdo creates any directories mentioned in the command (if they don't already exist).

*Note that mkdo uses trailing slashes (or backslashes in Windows) to identify folder names.*

e.g.

         mkdo mv error.log old/2013-01-01/

This would create the directories "old/2013-01-01/" before running the 'mv' command.


e.g. 2

         mkdo notepad newdir/newfile.txt

This would create the directory 'newdir' before starting up notepad. In Linux this even works for console-based editors like vim, but note that in Windows the piping doesn't quite work properly.


Downloads
---------
If you're already running go, then please just go get:

      go get github.com/laher/goxc

Otherwise, download [latest binaries](http://laher.github.com/mkdo/dl/latest/) for Linux, Mac, Windows.

Lame joke
---------
With apologies, mkdo comes with a lame rip-off of the [xkcd sudo joke](http://xkcd.com/149/):

 - child: put that cake in my lunchbox/
 - parent: What?! You haven't even got a lunchbox/
 - child: *mkdo* put that cake in my lunchbox/
 - parent: okay

License
-------

   Copyright 2013 Am Laher

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

To do
-----

 - get a better joke
 - Insertion of current date/time using special notation, e.g. mv x.log old-[yyyymmdd]/
 - Add options to ignore certain arguments.
 - Attempt to fix piping for Windows
 - Testing on OSX
 - more extensive docs
