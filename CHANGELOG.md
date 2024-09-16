# Changelog

## 1.1.0

gotag:

- rename `NewValueWriters()` to `NewFirstSupportedValueWriter()`
- rename `NewValueReaders()` to `NewFirstSupportedValueReader()`
- add `NewFirstSucceedValueWriter()` to support error handling
- add `NewFirstSucceedValueReader()` to support error handling

gotagio:

- add `NewBoolWriter()` to parse `bool` from custom strings
- add `NewTimeWriter()` to parse `time.Time`
- add `NewTimeReader()` to format `time.Time`
- add `WriteDuration()` to parse `time.Duration`
- add `ReadDuration()` to format `time.Duration`

gotaghttp:

- rename tag `requestBody` to `body`
- rename tag `pathParam` to `pathValue`
- rename `TagPathParam` to `TagPathValue`
- rename `NewPathParamWriter()` to `NewPathValueWriter()`
- rename `PathParamWriter` to `PathValueWriter`
- rename tag `queryParam` to `query`
- rename `TagQueryParam` to `TagQuery`
- rename `NewQueryParamWriter()` to `NewQueryWriter()`
- rename `QueryParamWriter` to `QueryWriter`
- add `IsExecutionError()`

## 1.0.0

Initial release
