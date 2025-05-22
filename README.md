# Nena

Git commands can get quite repetitive when you do dozens a day. So I wrote Nena to run away a little bit from the verbose. For example, usually, when you are going to add something into a repository, you need to do these 3 commands:

```bash
git add *
git commit -m "commit message"
git push
```

Nena abstracts them into a single command, which internally runs the 3 git commands above:

```bash
nena -m="commit message"
```

So now I reduced the amount of git commands I do per day by 3 times, yipeee!

## Installation guide

It's pretty easy to install Nena, actually. You just need to have Golang installed on your machine and run the following commmand:

```go
go install github.com/nandowastaken/nena
```

After this, you are good to go, just use Nena as much as you want to!
