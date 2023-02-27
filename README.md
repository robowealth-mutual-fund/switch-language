## Switch Language
The modules is contains the golang utilities for internal services.

## Installation
1. Use the below Go command to install Switch Language Utility
 ```bash
 go get github.com/robowealth-mutual-fund/switch-utility
 ```
2. Import it in your code
 ```
 import "github.com/robowealth-mutual-fund/switch-utility"
 ```

## Set up your project to support private modules
Mostly setup command in this instructions base on `git.robodev.co (GitLab)`

### Go
Go version >= 1.13 (RECOMMENDED)
```bash
go version # To know your go version

```

### Git
Under the `go get` command is using Git to pull the specified versions of your dependencies. So, the git configuration for wherever Go is running ***has to have*** the appropriate credentials to access any private repositories (In this case is GitLab).

>>>
How to get personal access token on GitLab [Here](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)
>>>


>>>
This is great for local development, but what about my CI/CD pipeline?
>>>