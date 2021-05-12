# Go Micro-Service Template

This is a Go micro-service written from scratch.

It shows how to use [net/http](https://godoc.org/net/http), and how to structure a Go project.

Dependency injection is used to insert a logger instance into the handler.

You can also notice how the test is constructed in order to provide testing for the handler.

A Docker container is available, thanks to the Dockerfile. It shows how to construct such containers.

## References

### Structuring Go applications

In order to learn how to approach package design in Go, you can read the following resources:

- [Style guideline for Go packages - JBD](https://rakyll.org/style-packages/)
- [Standard Package Layout - Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp)
- [Go best practices, six years in - Peter Bourgon](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)

Once done, this article will help you understand the [Design Philosophy On Packaging by William Kennedy](https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html).

### Exposing Go applications to the Internet

[This article](https://blog.cloudflare.com/exposing-go-on-the-internet/) describes how you can start approaching

## License

This project is under the MIT license. Please see the [LICENSE](LICENSE) file for more details.
