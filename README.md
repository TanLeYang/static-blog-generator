# Go-SSG
A simple static site/blog generator written in Go. I wrote this as a way to practice/learn using Go and also to create my personal site where I can write blog posts with markdown rather than HTML :)

## How to use
### Prerequisites
You must have Go installed.

### Installation
1. Clone this repository
2. Run `go build .` to build an executable for your OS
2. Copy the executable to a directory where you want to keep your blog posts and html templates

### Directory setup
The exectuable looks for folders named `data` and `template` in the same folder as where it is placed. 

Inside the `data` folder there should be a single `config.yml` file and a `posts` folder. 

Each individual post should be placed into its seperate folder inside the `posts` folder. There is a naming structure that must be followed for each post. Suppose the folder is named `new_post`, then the markdown file containing the post itself should be named `new_post.md`, images referenced by the markdown should be placed inside a folder `new_post_images` and there should be a `meta.yml` file. You may refer to the data folder in this repo as an example. 

The template folder should contain the same html files as those in this repo. Feel free to edit the contents of the file to your liking, however the variables, which are marked as {{.XXX}} inside the HTML files should still be present.

### Todos
Since this tool is supposed to be used to build my personal website that will serve as both a blog and a portfolio of sorts, perhaps I will implement adding projects, skills, etc.
