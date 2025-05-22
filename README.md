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

Lately, I've also been running into a very specific issue, in which I need to commit the same git historic to two different repositories. Which would look like this:

```bash
git add *
git commit -m "commit message"
git push

git remote rm origin
git remote add origin github.com/2
git push

git remote rm origin
git remote add origin github.com/1
```

Nena abstracts this weird necessity into a single command:

```bash
nena -c=true -m="commit message"
```

Yes, so now I don't need to write 8 commands whenever I want to make a deploy, so yaay!

## Installation guide

It's pretty easy to install Nena, actually. You just need to have Golang installed on your machine and run the following commmand:

```go
go install github.com/nandowastaken/nena
```

After this, you are good to go, just use Nena as much as you want to!
