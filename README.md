# github 

a cli tool to work with github.

## features 

**`auth`**

- [ ] `auth` github.com
- [ ] `auth` github enterprise

**`repositories`**

- [ ] list owned `repositories`
- [ ] list organization `repositories`
- [ ] list watched `repositories`
- [ ] list starred `repositories`
- [ ] list contributed `repositories`
- [ ] checkout a `repository` (git must be installed)
    - force ssh option
    - force http option
- [ ] create a `repository` (interactive)
    - select login (if more than 1)
    - select owner / organization
    - set a name
    - optional select from list of templates
    - optional create a readme
    - optional select and create a license
    - optional select and create a gitignore file
    - optional create dot-github files
    - optional skip checkout
    - all options can also set with flags (like non interactive)
- [ ] create a `repository` (non-interactive)
    - all known interactive options as flags
    - check for a .githubrc file (current dir / home dir)
    - check for a GITHUB_RC environment variable
    - option priority: `flags > env var > file`
    - merge options
- [ ] modify the `repository` description
- [ ] modify the `repository` topics
- [ ] modify the `repository` access (public / private)
- [ ] modify the `repository` contributors
- [ ] delete a `repository`
- [ ] checkout the `repository` wiki
- [ ] print `repository` details
- [ ] watch `repository`
- [ ] star `repository`
- [ ] fork `repository` (interactive)
- [ ] fork `repository` (non-interactive)

**`issues`**

- [ ] list the `issues` of a repository
- [ ] create an `issues` (interactive)
- [ ] create an `issues` (non-interactive)
- [ ] modify an `issues` (interactive)
- [ ] modify an `issues` (non-interactive)
- [ ] close an `issues`
- [ ] delete an `issues`

**`labels`**

- [ ] list the `labels` of a repository
- [ ] list the `labels` of an issue
- [ ] modify the `labels` of an issue

**`milestones`**

- [ ] list `milestones` of a repository
- [ ] show the milestone of an `issue`
- [ ] modify the milestone of an `issue`

**`notifications`**

- [ ] list `notifications`
- [ ] modify `notification`
- [ ] delete `notification` 

**next**

- [ ] `repository settings`
- [ ] `projects`
- [ ] `actions`
- [ ] `gist`
- [ ] github cli **repl** (https://github.com/c-bata/go-prompt)
- [ ] github cli **ui** (https://github.com/gizak/termui)

## extra todo

**actions**

- [ ] test & coverage
- [ ] release build + upload

**installer + description**

- [ ] homebrew
- [ ] chocolatey
- [ ] apt / yum
- [ ] releases download
- [ ] go get