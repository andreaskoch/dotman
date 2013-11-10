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

	dotman [-whatif] <command> [<filter>]

**The -whatif flag**

If you want to **test** what a certain command you can precede it with the `-whatif` flag.
This will ensure that no files are modified or copied.

Example:

```bash
dotman -whatif deploy
```

**Commands**

These are the available commands:

- **clone**: Clone a dotfile repository.
- **list**: Get a list of all modules in the current repository.
- **import**: Import files based on your current dotman configurations.
- **backup**: Backup your target files.
- **deploy**: Deploy your modules.
- **changes**: Show changed files.
- **commit**: Commit all changes.

**Filter**

If you want to restrict the scope of the "import", "list", "changes" or "deploy" command to a specific module or a set of modules you can follow the command with a **module-filter**.

```bash
dotman import <filter>
```

The filter can be just the name of the module or a full-blown [(RE2 compliant) regular expression](https://code.google.com/p/re2/wiki/Syntax).

### Getting help

If supply the `help` command to dotman (or any unknown command for that matter) it will print out the help dialog:

```bash
dotman help
```

	v0.1.0 - Backup and bootstrap your dotfiles and system configuration.

	usage: dotman [-whatif] <command> [<filter>]

	Available commands are:
	    clone     Clone a dotfile repository.
	    import    Import files based on your current dotman configurations.
	    list      Get a list of all modules in the current repository.
	    backup    Backup your target files.
	    changes   Show changed files.
	    deploy    Deploy your modules.
	    commit    Commit all changes.

	Options:
	    whatif    Enable the dry-run mode. Only print out what would happen.

	Arguments:
	    filter    You can add a module filter expression to the import, list, changes and deploy commands.

	Contribute: https://github.com/andreaskoch/dotman

### Cloning a dotfile repository

To clone an existing dotfile repository to your current working directory use the `clone` command.

```bash
dotman clone <repository-url>
```

This command will execute a `git clone --recursive` for the supplied repository url.

### Creating a dotfile-repository with "import"

If you want to start a new dotfile-repository from scratch for example for your vim files you can follow these steps:

#### 1. create a repository folder (e.g. "dotfiles")

```bash
mkdir -p ~/src/dotfiles/vim
```

#### 2. create a dotman file with mappings for your vim configuration

```bash
cat << EOF > ~/src/dotfiles/dotman
# Your .vimrc file
vimrc        		~/.vimrc

# Your .vim folder
vim/autoload        ~/.vim/autoload
vim/bundle          ~/.vim/bundle
EOF
```

#### 3. import the file into your repository

```bash
cd ~/src/dotfiles
dotman import
```

This will copy your ".vimrc", and the ".vim/autoload" and ".vim/bundle" folder into your new dotfile-repository - which gives you a good starting point for refining your personal dotfile repository.

### Getting a list of all modules in your current dotfile-repository

To get a list of all dotman-modules in the current directory use the `list` command.

```bash
dotman list
```

### Backup your dotfiles

To backup all files files that are mapped in your current dotfile-repository you can use the `backup` command.

```bash
dotman backup
```

This command will create a *.tar archive in the ".backup" folder of your dotfile-repository which contains all mapped target files. This an easy way to backup your system configuration.

### Showing changed files

To see which files have changed between your dotfile-repository and the target you can use the `changes` command.

```bash
dotman changes
```

This command will print out a list of all files that have changed, grouped by module.

### Deploy your dotfile-repository

The `deploy` comamnd will copy all mapped files from your dotfile-repository to the defined target locations.

```bash
dotman deploy
```

**Note**: If you are afraid what might happen when you execute this command you can add the `-whatif` flag. This way dotman will not copy any files but will show you what it would do:

```bash
dotman -whatif deploy
```

### Commit all changes to your dotfile-repository

To commit all changes to your dotfile-repository you can use the `commit` command followed by a commit message.

```bash
dotman commit "<your commit message>"
```

This will perform a `git add -A .` followed by a `git commit -m "<your commit message>"` on each module of your dotfile-repository and then on your dotfile-repository itself.

## Contribute

If you have an idea how to make this tool better please send me a message or a pull request.
All contributions are welcome.
