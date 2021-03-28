## ArgoCD CLI
An experimental refactoring of ArgoCD to separate the CLI code from the original codebase.

### argocd
A CLI for ArgoCD users.

```
go build -o bin/argocd cmd/main.go
```

### argocd-util
A CLI for ArgoCD operators.

```
go build -o bin/argocd-util cmd/main.go
```


