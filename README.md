# eagle 🦅

Well, what is eagle? Eagle is a simple, fast, and fun CLI-based application which functions as a helper to find answers to your programming questions.

Eagle works by searching your questions in Stack Overflow, and allowing you a plethora of options to get your answer to you as quick as possible.

This project was created with the intention of learning the language in more depth.

Takeaways:
- Got to know 'Cobra' in more detail, and I wasn't too much of a fan due to the overwhelming out-of-the-box functionality. Might try other libraries in the future.
- Though the script was very limited in it's scope got to understand net/http as well as encoding/json better.
- Bits about httptest, testing an application in general.

## Installation

1. Clone the git repository.
2. Go inside the desired directory containing the main.go file.
3. Lastly run: $ go build

After that, you're ready to use eagle.

## Example(s)

To search for the query "How to install Go?", sorted by votes & a maximum of 5 results: (both work)
```
$ .\eagle.exe search "How to install Go?" votes 5
$ .\eagle.exe search -t: "How to install Go?" -s: votes -r: 5
```

or... How to center a div? Using a default sort & max results:
```
$ .\eagle.exe search "How to center a div?"
```

Example Display:

<p align="center"><img src="https://i.imgur.com/NIPwico.jpg"></p>

## Help

Any further information about this application is available under --help.

## License

Licensed under the MIT License - see the [LICENSE file](https://github.com/k9mil/eagle/blob/master/LICENSE) for more details.