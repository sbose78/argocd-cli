## ArgoCD CLI
An experimental refactoring of ArgoCD to separate the CLI code from the original codebase.

### argocd
A CLI for ArgoCD users.

```
go build -o bin/argocd cmd/main.go
```

### argocd-util
A CLI for ArgoCD operators.

#### Build
```
go build -o bin/argocd-util cmd/main.go
```

#### Usage

```
argocd-util has internal utility tools used by Argo CD

Usage:
  argocd-util [flags]
  argocd-util [command]

Available Commands:
  app         Manage applications configuration
  cluster     Manage clusters configuration
  export      Export all Argo CD data to stdout (default) or a file
  help        Help about any command
  import      Import Argo CD data from stdin (specify `-') or a file
  proj        Manage projects configuration
  repo        Manage repositories configuration
  settings    Provides set of commands for settings validation and troubleshooting
  version     Print version information
```
