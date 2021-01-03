# flex-git-config
This enables to set git config every domain ( for now, email and username )

Binary is [here](https://github.com/ebihara99999/flex-git-config/releases).

## Usage

```
# email, username and domain are required
# Firstly use the noop option and check the lists of the directories to edit configuration.
$ flex-git-config -u YOUR_USERNAME -e YOUR_EMAIL -d "github.com" -n true

# Then start applying.
$ flex-git-config -u YOUR_USERNAME -e YOUR_EMAIL -d "github.com"
```

## Prerequisite
- Install and setup [ghq](https://github.com/motemen/ghq) in advance
