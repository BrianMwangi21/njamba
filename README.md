# njamba

njamba is your OpenAI assistant on the terminal. It is a command line tool that uses the OpenAI API for natural language processing. It can be used to generate text, search for information, and answer questions. It is a work in progress and is currently in alpha stage. It is written in Go and uses Cobra for command line interface.

# Get started

## Prerequisite

- Golang - [https://go.dev/](https://go.dev/)
- OpenAI API keys - [https://beta.openai.com/](https://beta.openai.com/)

## How to run

1. Clone the project
   ```
   $ git clone https://github.com/BrianMwangi21/njamba
   ```
2. Navigate into the project
   ```
   $ cd njamba
   ```
3. Copy `.env.example` file to `.env` file and set your OpenAI API keys in the .env file using your editor of choice. [Neovim](https://neovim.io/) is highly recommended
   ```
   $ cp .env.example .env
   ```
4. Install the program
   ```
   $ go install
   ```
5. Run the program
   ```
   $ njamba
   ```
