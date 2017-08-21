# wikimedia

This package let's us interact with the wikimedia API.

## Grabbing links 

There's two main functions you can use to grab links from wikipedia.

### `GetPagesLinks`
This lets you grab one link at a time. Simple and easy.

``` go
links, err := GetPagesLinks("Cat")
```

### `GetManyPagesLinks`
_Note: I originally wrote this, but then put it aside since it was easier to gain one link at a time. I will eventually come back and use this._

This lets you bulk up some links and ask for several at a time. This is faster, but more complicated.
Note that this request will also send several requests out to "continue" as is [documented here.](https://www.mediawiki.org/wiki/API:Query#Continuing_queries)

``` go 
pages, err := GetManyPagesLinks([]string{"Cat", "Dog", "Airplane"}, "")
pages["Dog"] // Links only for "Dog"
```
