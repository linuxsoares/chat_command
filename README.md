# chat_command

**chat_command** is a `CLI tool` that generates and executes shell commands using OpenAI `(GPT-4 Turbo)` based on natural language instructions provided by the user—safely and interactively.

---

## Features

- **Natural Language Input:** Enter a phrase describing what you want to do.
- **AI-Powered Command Generation:** Uses OpenAI to convert your request into an actionable shell command.
- **User Confirmation:** See the generated command and approve or reject it before execution.
- **Interactive Shell Execution:** Executes only upon explicit user approval.

---

## Demo

```sh
$ go run main.go
Enter a phrase: list all files larger than 10MB in the current directory
Generated command line: find . -type f -size +10M
Can ? (yes/no)
yes
# Command output shows here
```

## License
[MIT License](https://rem.mit-license.org/)
