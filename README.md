# Pagination

This repo is a library to facilitate pagination alongside [Pop](https://github.com/gobuffalo/pop) and
[protocol buffers](https://developers.google.com/protocol-buffers/). It contains models, protobufs definitions,
and optional data accessor interfaces for interacting with a design of pagination that is endorsed by
[Google Cloud APIs](https://cloud.google.com/apis/design/design_patterns#list_pagination).

## What
There are three central concerns in this repo. The first is the actual Pop-based model struct itself for
`PageToken`s. This can be referenced like any other Pop model in your codebase (assuming a migration has
been applied to your database to facilitate the storage of the model). For example, you might retrieve a cursor
and pass it into a list function in order to know what `created_at` point the last pagination request
ended at:

```go
	cursor, err := pagination.GetCursor(ctx, dataAccessor, token)
	if err != nil {
		return nil, "", err
	}
	teams, more, err := ListData(cursor, pageSize)
	if err != nil {
		return nil, "", err
	}
```

Full examples may be found in the [beacon-people](https://github.com/solarwinds/beacon-people/blob/94838a9185f196d93dec7d0f225727d03c6484a6/peopleservice/user_service.go#L28)
repo.

## How
When using the pagination library it is necessary to do a few things up front:
1. Utilize the sample migration to create the backing database table necessary for page tokens to be stored
and recalled within your application.
2. Integrate the protobuf message definitions for pagination requests.
3. Use the Pop-based models to maintain state within your application for pagination requests.