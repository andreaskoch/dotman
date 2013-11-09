# dotman

dotman is a tool to bootstrap your system configuration files.

## Build Status

[![Build Status](https://travis-ci.org/andreaskoch/dotman.png)](https://travis-ci.org/andreaskoch/dotman)

## Motivation

Get up and running within seconds on a new system with all your favorite vim plugins, bash-scripts and tweaks that make you at home and productive.

dotman is inspired by [modman](https://github.com/colinmollenhour/modman) and allows you to easily manage your dotfile repositories.

Create a git repository which contains all your dotfiles such as your bash and vim configuration (e.g. [dotfiles-public](https://bitbucket.org/andreaskoch/dotfiles-public)):

	dotfiles-public/
	├── bash
	│   ├── bash-git-prompt
	│   │   ├── gitprompt.fish
	│   │   ├── gitprompt.png
	│   │   ├── gitprompt.sh
	│   │   ├── gitstatus.py
	│   │   └── README.md
	│   ├── bashrc
	│   ├── dotman
	│   └── scripts
	│       ├── contribue-source-to-target.sh
	│       ├── public-ip.sh
	│       ├── synchronize-source-to-target.sh
	│       └── vim-setup.sh
	└── vim
	    ├── dotman
	    ├── fontconfig
	    │   └── 10-powerline-symbols.conf
	    ├── fonts
	    │   └── PowerlineSymbols.otf
	    ├── vim
	    │   ├── autoload
	    │   └── bundle
	    └── vimrc

Add submodules from awesome modules such as [bash-git-prompt](https://github.com/magicmonty/bash-git-prompt) and then create a **dotman** mapping file for each submodule which tells dotman where to copy the files.

Take the `dotman` mapping for the vim-files listed above as an example:

	vimrc                                   ~/.vimrc
	vim/autoload                            ~/.vim/autoload
	vim/bundle                              ~/.vim/bundle

	#powerline fonts
	fonts/PowerlineSymbols.otf              ~/.fonts/PowerlineSymbols.otf
	fontconfig/10-powerline-symbols.conf    ~/.config/fontconfig/conf.d/10-powerline-symbols.conf

The file maps files and directories from your vim settings-repository directly into your home-directory.

## Terminology

### "Module"

A dotman **module** is a folder which contains a plain-text file named `dotman`.

### "Repository"

A **repository** is a collection of one or more modules. You can have only one module which contains all your dotfiles, but sometimes things get a little less messy when you seperate your dotfiles into more separate modules. Examples:

- git-configs
- vim
- bash
- ...

## Usage

	dotman <command>

## Getting help

If supply the `help` command to dotman (or any unknown command for that matter) it will print out the help dialog:

```bash
dotman help
```

	v0.1.0 - Backup and bootstrap your dotfiles and system configuration.

	usage: [-whatif] dotman <command> [args]

	Available commands are:
	    clone     Clone a dotfile repository.
	    list      Get a list of all modules in the current repository.
	    import    Import files based on your current dotman configurations.
	    backup    Backup your target files.
	    deploy    Deploy your modules.
	    changes   Show changed files.

	Options:
	    whatif    Enable the dry-run mode. Only print out what would happen.

	Arguments:
	    filter    You can add a module filter expression to all commands.

	Contribute: https://github.com/andreaskoch/dotman

## List of all modules

To get a list of all dotman-modules in the current directory use the `list` command.

```bash
dotman list
```

