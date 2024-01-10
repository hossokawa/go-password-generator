# PassGen
PassGen is a terminal-based password generator written in Go. The generated passwords can be of any length you want and you can choose to include numbers and/or symbols in them. Your generated passwords are not stored in any form whatsoever.

## Usage
Download the repo, open the folder in your terminal and use the Go commands `go build` and `go install` and you're ready to go.

To start generating passwords just type `passgen generate` for a password using the default settings (8 characters long, no number or symbols) or specify the flags you want to change:
| Flag         | Use     |
|--------------|-----------|
| -l <number> | Changes the length of the password      |
| -n      | Includes numbers
| -s      | Includes symbols

These are the symbols that could be included in your password: **!@#$%^&*()_-+={}[]?<>;:**
