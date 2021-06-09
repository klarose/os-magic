# os-magic

A simple golang project showing one way of hiding os-specific logic behind an interface, without having to do anything too crazy to bring into place the appropriate OS.

We rely on golang init functions to set everything up.

The general idea is that you have a package which exports an interface whose implementation depends on the operating system.
Packages underneath that package define the interface it is exporting, as well as per-os implementations of it. Each per-os package registers back with the main one. A per-os import
file in an 'import' package is the one spot where the different OS are linked together. Since golang will not compile files for other operating systems, the irrelevant imports will
be ignored, and the relevant ones will be brought in. This has the effect of ensuring that the implementation for the current OS is initialized.

## Motivation

We had some code which combined per-os files and functions within one package. This lead to a bunch of false positives and problems when running tools like staticcheck. A system structured
like this should behave better.

## Alternatives

One simple change to this approach could lead to a better interface, depending on whether you want an object to interact with, or a package. If you want, you could do something like
this to make it look like there's a per-os implementating in the package:

```
var theIf OsInterface

func init() {
    theIf = oslayer.Get(runtime.GOOS)
    if theIf == nil {
        panic(fmt.Sprintf("os %v not supported", runtime.GOOS)
    }
}

func DoThing() error {
    return theIf.DoThing()
}
```
