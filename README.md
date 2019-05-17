# makehugefile

> Copyright (C) 2019 David Capello
>
> This file is released under the terms of the MIT license.
> Read [LICENSE.txt](LICENSE.txt) for more information.

This little utility creates an huge file with pseudo-random data
(using the golang [cryptographically secure random number generator](https://golang.org/pkg/crypto/rand/)).

### Usage

    makehugefile [-g=N] [-f] filename...

Generate one file for each given `filename`, where:
* `N` is the number of GB to write in each given file (by default 1 GB), and
* `-f` enables a "fast mode" (it generates 1 MB of random data and
  then repeat it in the whole file).
