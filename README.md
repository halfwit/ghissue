# ghissue - simple utility to create issues on Github

`go install https://github.com/halfwit/ghissue`

## Usage

> This uses Plan9's factotum to fetch the oauth2 key (you have to get and store this yourself). PRs to modify this behavior are welcome. For use on Linux/Unix, see the plan9port project for details on setting up your factotum.


```
<cmd> | ghissue [-b branch] -t <issue title> <repo name>
```

ghissue will read from stdin to create issues on Github.com. <repo name> is a shortened version of the github.com URL, using just `someuser/example` instead of `https://github.com/someuser/example`.  `-b` is used to optionally select the branch, which defaults to `master`.

The output from a successful command will be the URL to the issue.

