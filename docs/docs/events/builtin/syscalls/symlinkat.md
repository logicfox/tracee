
# symlinkat

## Intro
symlinkat - Creates a symbolic link named by linkpath to an object referenced by target. 

## Description
Symlinkat creates a symbolic link with the name specified in `linkpath` to the object referenced by `target`. It differs from `symlink` in that `linkpath` is relative to the directory file descriptor provided in `newdirfd`.

Normally, symbolic links can only point to other files located in the same filesystem.
However, when the `target` argument is prefixed with `/proc/self/fd/`, it can
reference a file descriptor opened by the same process.

There are a few possible edge-cases when using `symlinkat`. If `linkpath`
already exists, the existing link will be overwritten, and if the directory
referenced by `newdirfd` is not writable, a `EACCES` error will be returned.

## Arguments
* `target`:`const char*`[U] - The target to which the symbolic link points.
* `newdirfd`:`int`[U] - The file descriptor for the target directory.
* `linkpath`:`const char*`[U] - The name of the link to be created.

### Available Tags
* K - Originated from kernel-space.
* U - Originated from user space (for example, pointer to user space memory used to get it)
* TOCTOU - Vulnerable to TOCTOU (time of check, time of use)
* OPT - Optional argument - might not always be available (passed with null value)

## Hooks
### sys_symlinkat
#### Type
Kprobes
#### Purpose
Trace calls to symlinkat, including the arguments passed to it.

## Example Use Case
Tracing the origin of symbolic links created in the system. This can be used to monitor privilege escalations and other malicious actions.

## Issues
The `target` argument can reference parts of the filesystem. This means that if `target` contains a relative path, its interpretation will depend on the current working directory of the process.

## Related Events
The `lstat` event can be used to check if a file is a symbolic link, and if so, which file or directory it points to.

> This document was automatically generated by OpenAI and needs review. It might
> not be accurate and might contain errors. The authors of Tracee recommend that
> the user reads the "events.go" source file to understand the events and their
> arguments better.