# SPDS
*spds* is both:
    - A library which implments different probabilistic data structures
    - A cli for the most common operations on streams (cardinality, frequent elements and sampling) 



spds (Small Probabilistic Data Structures) is a library 

## Installing

Using **spds** is easy. First, use `go get` to install the latest version of the library. This command will install the `spds` executable cmd along with the library and its dependencies:

```bash
    go get -u github.com/jomsdev/spds/spds
```

Next, include **spds** in your application:

```go
    import "github.com/jomsdev/spds"
```

## Examples


- **CLI**:

    The code for the command line programs can be found in the cmd folder. They are built on top of the spds library and cobra and expose 
    in a convenient way three of the basic features of spds: Cardinality, Frequent items and Sampling. The spds command has three subcomands 
    `spds cardinality`, `spds frequency` and `spds sampling`.

    ```go
        Usage:
        spds [command]

        Available Commands:
        cardinality Estimates the cardinality (number of distinct elements) of a stream
        help        Help about any command
        version     Print the version number of spds

        Flags:
        -a, --algorithm string   Algorithm (Default: KMV)
        -f, --file string        Input file (required)
        -h, --help               help for spds
        -s, --size int           Algorithm (default 256)

        Use "spds [command] --help" for more information about a command.

    ```

    Run `cmd --help` for more information.

## References

Problems like cardinality estimation or heavy hitters have been arround for decades in both *academia* and *industry*. 
With the raise of the internet, big data and IoT this is a hot topic in research right now. Relevant papers, videos and posts
can be found in the [REFERENCE](https://github.com/jomsdev/spds/blob/master/LICENSE) file.

## Contributing

First off, thanks for taking the time to contribute! 

Now, take a moment to be sure your contributions make sense 
to everyone else and please make sure to read the [Contributing Guide](https://github.com/jomsdev/spds/blob/master/CONTRIBUTING.md)
before making a pull request.

## Issue tracker

Found a problem? Want a new feature? First of all see if your issue or idea has [already been reported](../../issues).
If it hasn't, just open a [new clear and descriptive issue](../../issues/new).

## License

 LibDori is under the MIT license. See the [LICENSE](https://github.com/jomsdev/spds/blob/master/LICENSE) file for details.
