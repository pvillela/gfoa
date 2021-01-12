# gfoa -- Golang Function-Oriented Architecture

This repo is a Go module.  The directories are Go packages containing:
- A software architecture framework to support a function-oriented modular structure for applications
- A sample application to demonstrate the function-oriented modular structure and the use of some of the frameworks

The idea is to structure applications as the composition of functions.  Each function must belong to a category (denominated a *module stereotype*) 
that that has certain well-defined responsibilities.

The **gfoa** Go module is not to be confused with the application modules which, in this architecture style, are implemented as Go functions.
