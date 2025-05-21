# git-cli-go
So, usually, to commit to a repo, I need to:

```bash
git add *
git commit -m "commit message"
git push
```

So I abstracted this into:

```bash
git-go -m="commit message"
```

Which runs the 3 commands above. 

And, lately, for some very specific reason, I needed to be posting the same commit history to two git repositories, which would be something like:

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

As this was starting to getting annoying, I abstracted into:

```bash
git-go -c=true -m="commit message"
```

Yes, so now I don't need to write 8 commands whenever I want to make a deploy, so yaay! 

## Installation guide

Welp, it is pretty easy. First, clone this repo. Then, build the project:

```go
go build -o git-go
```

And move this to your path, this is how I do it at MacOs:

```bash
sudo mv git-go /usr/local/bin/
```

If you are using another OS, probably you will need to find another way to move to path, but it shouldn't be too hard.

After that, you can just invoke the commands as mentioned before. 

