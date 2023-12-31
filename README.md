# sedplus

a sed-like tool designed to replace sed with easier to read syntax

## Usage

```
cat config | sedplus --find "192.168.1.1" --replace "192.168.3.2" --error-if-not-found
```

## Parameters

`--lowercase` converts text to lower ( SED -> sed )

`--uppercase` convert text to upper ( sed -> SED )

`--find 'string-to-find'`

`--find-line 'ip:'` finds a line containing the string and replace the whole line

`--replace 'new-string'`

`--error-if-not-found` - if the find string is not found, exit with an error code otherwise it will return exit 0.

`--case-insensitive` - case insensitive find

`--trim` removes whitespace from the start and end of each line


`--numeric` remove all characters not a digit 0-9 ( asd1234 -> 1234 )

`--alpha`

`--alphanumeric`

`--compact-whitespace` reduce all spacing to a single space (processed last)

## Helpers
Removing common tasks from the input stream

`--remove-quotes`

`--remove-time`

`--remove-date`

`--remove-ip`

`--remove-uuid`


## To Install

**Mac**

`brew install szazeski/tap/sedplus`

**Linux (and mac)**

`wget https://github.com/szazeski/sedplus/releases/download/v0.1.0/sedplus_0.1.0_$(uname -s)_$(uname -m).tar.gz -O sedplus.tar.gz && tar -xf sedplus.tar.gz && chmod +x sedplus && sudo mv sedplus /usr/bin/`

**Windows**

`Invoke-WebRequest https://github.com/szazeski/sedplus/releases/download/v0.1.0/sedplus_0.1.0_Windows_x86_64.zip -outfile sedplus.zip; tar -xzf sedplus.zip; echo "if you want, move the file to a PATH directory like WINDOWS folder"
`
