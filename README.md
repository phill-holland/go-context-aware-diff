# Context Aware Code Diff 

A simple proof-of-concept diff tool implementation rewritten from C++ to Go, explicitly for code files, to help simplify and improve the output from traditional diff algorithms.

Designed to work on programming languages which using curly brackets to denote code blocks and semi-colons to end statements (for example languages like C++, C#, Java, etc..)

The code is loaded with a parser and into a tree like structure, with spaces and tabs striped.

The original idea was the need for a diff tool for files which may be reorganised but functionally the program would remain the same, but the procedure definition blocks could be implemented in different positions (however this feature still needs to be implemented here, the method however should make this trival).

Work in progress!

See C++ version here;

https://www.github.com/phill-holland/context-aware-diff

And the rust version;

https://www.github.com/phill-holland/rust-context-aware-diff

# Usage

diff filename1 filename2

# Output

A combined output, using the simple notion of + to highlight when a line is added, or - to highlight when a line is removed.

Note! Current implementation output stripes spaces and tabs from original source files(due to the algorithm implemented to determine the differences)

# Running

- Ensure project is open within the VSCode development container
- Hit F5

# Requirements

The VSCode development container plugin is installed;

https://code.visualstudio.com/docs/remote/containers

Docker must also be installed;

https://docs.docker.com/get-docker/

This application, however is configured with linux based containers, and will not work correctly on Windows without modification.
