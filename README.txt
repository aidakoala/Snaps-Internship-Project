Hello,
assertions folder contains assertions and signed assertions.
.snap folder contains a gpg key, created with 'snap create-key'. Password is 1234. Key was used to sign assertions.
snap-sources contains an opensshd folder which was modified to match the size/assertions.
ihpatch.txt is the patch against snapd, over the commit:
 commit a53f527f976ac539c3e018b16f4e5082eba80682
 Merge: b814491 5c25377
 Author: Pawel Stolowski <stolowski@gmail.com>
 Date:   Tue Sep 11 08:34:49 2018 +0200
stefi-dev-repo is the full folder with server code, it has hardcodings.

I used snap hidden commands like snap create-key, snap sign.

Happy coding!

