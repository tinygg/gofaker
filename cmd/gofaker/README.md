# Gofaker Command Line Tool
Run gofaker via that command line. 
All functions are available to run in lowercase and if they require additional parameters you may pass by space seperation.

### Installation
```go
go get github.com/tinygg/faker
```

### Example
```bash
gofaker firstname // billy
gofaker sentence 5 // Exercitationem occaecati sed dignissimos inventore.
gofaker shufflestrings hello,world,whats,up // up,hello,whats,world
```

### List of available functions
```bash
gofaker list
```

![](https://raw.githubusercontent.com/tinygg/faker/master/cmd/gofaker/cmd.gif)