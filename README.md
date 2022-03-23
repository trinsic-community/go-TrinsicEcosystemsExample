# Go samples for Trinsic SDK

### Notes on pulling latest packages

Replit creates a cache of the Nix channel and it may install previous version.
In order to use a package latest store exclude the package installation from `.replit` file and run the following steps

```
nix-channel --update
nix-env -iA nixpkgs.okapi
nix-shell -p okapi
```
