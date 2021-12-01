# Rewrite

Rewrite is a simple wrapper around `wget` or file download cli. The sole purpose is to
substitute the download url with an alternative url. 

## Usage

The basic usage command can be executed as  `rewrite <config json> <wget including args> <url>`. 

This makes the following assumptions about the passed parameters:
* The first parameter is the fully qualified path to the config file location.
* The last parameter is the url

## Sample config file

```json
{
  "urls": {
    "https://ftp.pcre.org/pub/pcre/pcre-8.44.tar.bz2": "https://sourceforge.net/projects/pcre/files/pcre/8.44/pcre-8.44.tar.bz2"
  }
}
```

