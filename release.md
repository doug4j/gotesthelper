# Release Instructions

## Audience

This section is intended for project maintainers, not application developers.

## Terms

<table>
  <tr>
    <td>NAME<td>
    <td>VALUE<td>
  </tr>
  <tr>
    <td>${ProjectRoot}<td>
    <td>The root folder where the source code lives<td>
  </tr>
</table>

## 1. Open a command line to the product directory

```zsh
cd ${ProjectRoot}
```

## 2. Process Commits to master and Push Changes

```zsh
## Local commit
git commit -a -m "Doc: Improved release instructions"
## Push to master
git push origin master
```

## 3. Create a Tag for the release

```zsh
git tag -a v0.1.0 -m "First release"
```

## 4. Push to Git tag to Origin

```zsh
git push origin v0.1.0
```

## 5. Publish with goreleaser

```zsh
goreleaser release --rm-dist
```

## 6. Verify build (using version)

```zsh
./dist/templ_darwin_amd64/templ version
```
