# Protobuf Basics

Fundamentals and basic examples of Google's Protobuf in Go and Python

## Setting up the Protoc compiler

Find the correct protocol buffers version based on your Linux distro below [here](https://github.com/google/protobuf/releases).

### Installation on MacOSX

    brew install protobuf

## Basics

An example of a proto message is:

    syntax = "proto3";

    message Account {
        uint32 id = 1;
        stirng name = 2;
        bool is_active = 3;
    }

### Defaults

- When not set:
  - A field will not be serialized.
  - Populated with default value.

### Scalar Types

- Numbers
  - int32, int64, sint32, sint64
  - uint32, uint64
  - fixed32, fixed64, sfixed32, sfixed64
  - float, double
  - All above types will have a default value of 0.

- Strings
  - Value: Arbitrary length
  - Default: Empty string
  - Encoding: UTF-8
  - Keyword: string

- Bytes:
  - Value: Arbitrary length byte sequence
  - Default: Empty bytes
  - Encoding: Up to you to interpret in code
  - Keyword: bytes

### Tags

- Field names are not important for serialization, but `Tags` are. Tags impact your payload.
- Rules:
  - Smallest: 1
  - Largest: 536,870,911
- Reserved Tags
  - From 19000 to 19999
- How `Tags` are serialized?

### Repeated Fields

- Repeated fields can be repeated any number of times, including zero.
- Repeated:
  - Keyword: repeated

        repeated <type> <name> = <tag>;

  - Value: Any number of elements (0 or more)
  - Default: Empty list
  - Example:

        syntax = "proto3";

        message Account {
            uint32 id = 1;
            string name = 2;
            bytes thumbnail = 3;
            bool is_active = 4;
            float height = 5;
            repeated string phones = 6;
        }

### Enumerations (Enums)

- Enum:
  - Keyword: enum
  - Value:
  - Default: First value in enum
  - Example:

        enum EyeColor {
            UNSPECIFIED = 0;
            GEEN = 1;
            BLUE = 2;
            BROWN = 3;
        }
        message Account {
            uint32 id = 1;
            string name = 2;
            bytes thumbnail = 3;
            bool is_active = 4;
            float height = 5;
            repeated string phones = 6;
            EyeColor eye_color = 7;
        }
- First tag in enum is always 0

### Comments (Documentation)

- Single line comments
- Multiline comments
- Examples:

        /*
        Account used to identify someone in our system
        */
        message Account {
            uint32 id = 1;
            string name = 2;
            bytes thumbnail = 3; // picture in jpeg format
            bool is_active = 4;
            // height in centimeters
            float height = 5;
            repeated string phones = 6;
            EyeColor eye_color = 7;
        }

### Generating code in any language
