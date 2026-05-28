# git-branch-depth

Display commit history depth for each git branch.

## Features

- Show commit count per branch
- Sorted by depth (most commits first)
- Marks current HEAD branch with `<-`
- Works with any git repository

## Install

```bash
go install github.com/TataneSan/git-branch-depth@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-branch-depth.git
cd git-branch-depth
go build -o git-branch-depth .
```

## Usage

```bash
git-branch-depth
```

### Example Output

```
BRANCH                                    COMMITS  
-------------------------------------------------------
main                                      1234  <-
feature/auth                                567  
develop                                     890  
hotfix/login                                 23  

Total: 4 branch(es)
```

## License

MIT
