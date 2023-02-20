## go.sum files

A module may have a text file named go.sum in its root directory, alongside its go.mod file. The go.sum file contains cryptographic hashes of the module’s direct and indirect dependencies. When the go command downloads a module .mod or .zip file into the module cache, it computes a hash and checks that the hash matches the corresponding hash in the main module’s go.sum file. go.sum may be empty or absent if the module has no dependencies or if all dependencies are replaced with local directories using replace directives.

Each line in go.sum has three fields separated by spaces: a module path, a version (possibly ending with /go.mod), and a hash.

1. The module path is the name of the module the hash belongs to.

2. The version is the version of the module the hash belongs to. If the version ends with /go.mod, the hash is for the module’s go.mod file only; otherwise, the hash is for the files within the module’s .zip file.
The hash column consists of an algorithm name (like h1) and a base64-encoded cryptographic hash, separated by a colon (:). Currently, SHA-256 (h1) is the only supported hash algorithm. If a vulnerability in SHA-256 is discovered in the future, support will be added for another algorithm (named h2 and so on).

3. The go.sum file may contain hashes for multiple versions of a module. The go command may need to load go.mod files from multiple versions of a dependency in order to perform minimal version selection. go.sum may also contain hashes for module versions that aren’t needed anymore (for example, after an upgrade). go mod tidy will add missing hashes and will remove unnecessary hashes from go.sum.