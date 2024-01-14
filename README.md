# tsc-node
### A simpler version of `ts-node` written in golang

`ts-node` is well known **TypeScript execution and REPL for node.js**. Essentially, it allows the user to directly run a `.ts` file and more.
But I just wanted the ability to run `.ts` files without any extra features. So, I implemented a basic golang package which exactly does that.
Here is how you install the library:

```bash
go install github.com/ritchielrez/tsc-node@latest
```

Then, just run `tsc-node`. Here is how to use it:

```bash
tsc-node <filename>
```

Replace `<filename>` with the file that you want to run. Feel free to open pull request or issues to help improve this project.
