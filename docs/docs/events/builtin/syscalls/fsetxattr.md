
# fsetxattr

## Intro
fsetxattr - Set an extended attribute for a file

## Description
The `fsetxattr()` system call sets an extended attribute for the given file system object relative to the open file descriptor `fd`. The open file descriptor `fd` is used to reference the object which should have the attribute set.

The `flags` argument determines how the attribute is set. Attributes are creation, security, scalability and installation related informations.

Advantages of using `fsetxattr()` is that it provides more control than `setxattr()`, which is an analogous system call, that is used to affect the objects inside a file system.

## Arguments
* `fd`:`int`[K] - File descriptor. The file descriptor should reference a regular file.
* `name`:`const char*`[K] - Name of the attribute.
* `value`:`const void*`[K U] - Pointer to the supplied value. Its format is determined by the implementation.
* `size`:`size_t`[K] - Size of the value referenced in the `value` argument.
* `flags`:`int`[K] - Flags designating how the attribute should be set.

### Available Tags
* K - Originated from kernel-space.
* U - Originated from user space (for example, pointer to user space memory used to get it)
* TOCTOU - Vulnerable to TOCTOU (time of check, time of use)
* OPT - Optional argument - might not always be available (passed with null value)

## Hooks
### do_fsetxattr
#### Type
Kprobes
#### Purpose
To monitor the execution of `fsetxattr()` system call.

## Example Use Case
This event could be used to monitor the setting of custom attributes on files through the `fsetxattr()` system call. This is useful for implementing metadata management, security policies or system customization features.

## Issues
If the `flags` argument is not being checked, an attacker could set an extended attribute for a file that it shouldn't have access to.

## Related Events
* `fgetxattr` - Get an extended attribute from a file
* `setxattr` - Set an extended attribute to a file
* `removexattr` - Remove an extended attribute from a file

> This document was automatically generated by OpenAI and needs review. It might
> not be accurate and might contain errors. The authors of Tracee recommend that
> the user reads the "events.go" source file to understand the events and their
> arguments better.