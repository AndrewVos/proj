# proj

A project management tool which is kind of like a note-taking tool except that
each project is it's own markdown file.

This means you can store anything you want in there really.

I've found that it's nice to have this bit of separation for each "project" that I take on.

## Usage

```
proj list
proj create "My New Project" # will automatically open the project in $EDITOR
proj edit 1 # Edit a project by it's identifier
proj complete 1 # Mark a project complete
proj list-all # List all projects, even completed ones
proj incomplete 1 # Mark a project incomplete
```
