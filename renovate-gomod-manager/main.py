#!/usr/bin/env python3

"""Main module"""


import os

def main() -> None:
    """Main function"""
    if not os.path.isfile('/workspace/go.mod'):
        print("Error: go.mod file is missing")
        return
    print("Success: /workspace/go.mod file found")

if __name__ == "__main__":
    main()
