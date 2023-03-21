# njamba - now using gpt-3-turbo!

njamba is your OpenAI assistant on the terminal. It is a command line tool that uses the OpenAI API for natural language processing. It can be used to generate text, search for information, and answer questions. It is a work in progress and is currently in alpha stage. It is written in Go and uses Cobra for command line interface.

> Remember, AI will not replace you. A person using AI will.

# Get started

## Prerequisites

- Golang - [https://go.dev/](https://go.dev/)
- OpenAI API keys - [https://beta.openai.com/](https://beta.openai.com/)

## How to run

1. Clone the project
   ```
   $ git clone https://github.com/BrianMwangi21/njamba.git
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

## Commands

This CLI has the following commands:

### prompt

This command is used to request a completion response from the OpenAI API.
How to run the `prompt` command :

```
$ njamba prompt
```

On running this prompt, you will be guided through :

1. Prompt input
2. GPT Model selection - Text completion models are shown by default.
3. Temperature input
4. Max tokens input

The request will then be sent and the response shown.

This prompt command also accomodates for the flag `--code`. Using this flag i.e `$ njamba prompt --code` will then allow the user to select the code completion models instead of the default text completion models.

## Why though ?

**You** : "Why is this necessary when there is the sweet sweet [ChatGPT](https://chat.openai.com/) available for everyone and I don't even need the API keys to use?" <br />
**Me**: "True. That is a valid point. I often ask myself the same. Here's the thing : ChatGPT is amazing but the fact that it is open to everyone, the downtimes are becoming more frequent. With this tool and using your own API keys, you can use it as much as you want - or until you reach your quota. Also, and hear me out, it's just awesome on the terminal." <br /><br />

**You**: "Ok, but do you use the sweet sweet gpt-3-turbo?" <br />
**Me**: "Why, yes!" <br /><br />

## Tests

![Surely...](https://media1.giphy.com/media/jOpLbiGmHR9S0/giphy.gif)
