rclone-unobscure
================

For some unknown reason, [rclone](https://rclone.org/) obfuscates all passwords
stored in its configuration file. To do so, it just encrypts the password with
AES in counter mode using the static key
`9c935b48730a554d6bfd7c63c886a92bd390198eb8128afbf4de162b8b95f638`
(implemented in
[`fs/config/obscure/obscure.go`](https://github.com/rclone/rclone/blob/v1.50.1/fs/config/obscure/obscure.go)).
This obviously does not add any security at all and is just annoying if you
want to retrieve a password from the configuration for whatever reason.

This little 5 line Go program will look for a rclone configuration in the
standard locations (typically `~/.config/rclone/rclone.conf`), decode the
passwords and print them. To use it, just clone the repository, run `go build`
and execute the resulting `./rclone-unobscure` binary.
