## Design
The design of the sdk is to loosely wrap the data coming The One API. A user has access to the following core funcitons
```
movie.GetMovies() -> return list of Movie Objects
movie.GetMovieByID(movieID) -> returns Movie Object matching ID
quote.getQuotes() -> returns all Quote Objects
quote.GetQuoteByID(quoteID) -> reurn Quote Object matching ID 
quote.GetQuotesFromMovie(movieID) ->  returns all Quote objects from a movie
```
Each of these functions uses the Request package to call the appropriate endpoint and return a slice
of bytes which the function can then unpack into a Quote or Movie object, matching the structure of the 
json data returned. This is then returned to the user for them to manipulate as they need. There are other
functions I exposed for these two objects but they are not necessary and were merely for testing/playing
with the data.


## Future Improvements
- Better error handling, create custom errors for common situations, potentially look into retrying
when error occurs as opposed to returning early.

- Extend Usability, currently I am just returning the data from the endpoints, making this a pretty light wrapper.
If given more time I would potentially look into common use cases, making it easier to interact with the data or providing
common functions that users are likly to use.

- Add Support for Pagination, Sorting and Filtering. The benifits of adding these primarily seem like they would help
any use cases that interact with character information, i.e. filtering all quotes from a certain character. As this sdk
does not support character data yet I have chosen to not include this. Pagination might be useful when working with quotes,
but I made the descition to delay adding it.

- I made the decision to expose the underlying data structure for Movies and Quotes, which in this case I believe to be valid as
it is data returned from the API.

- The API does rate limit users to 100 request every 10 minutes so it might be looking into ways to cache data between requests,
this does introduce complexity and ideally the user would handle this, but it is worth keeping in mind.

- Currently the API base url is hardcoded in the request.go file. This is an issue if the url ever changes, and or if a user wants to
use a different version of the API. A smarter way to do this would be to allow the user to set the base url, maybe in a config file.
This API is seemingly pretty stable and I don't see any reason why it would be updated so I am fine with hardcoding it.

- Add support for more ways to store the API token. For the current use case and scale of the project I think locally exporting
the api key is fine, but supporting other options down the road would be helpful.
