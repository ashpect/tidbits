```go
type Layer interface {
	// Digest returns the Hash of the compressed layer.
	Digest() (Hash, error)

	// DiffID returns the Hash of the uncompressed layer.
	DiffID() (Hash, error)

	// Compressed returns an io.ReadCloser for the compressed layer contents.
	Compressed() (io.ReadCloser, error)

	// Uncompressed returns an io.ReadCloser for the uncompressed layer contents.
	Uncompressed() (io.ReadCloser, error)

	// Size returns the compressed size of the Layer.
	Size() (int64, error)

	// MediaType returns the media type of the Layer.
	MediaType() (types.MediaType, error)
}
```
Defines an interface called Layer.
Interfaces are used to specify a set of method signatures that a type must implement to be considered an instance of that interface.
Method signatures are the method name, the parameters, and the return types.
<br><br>

#### Difference between function and methods

- A function is a piece of code that is called by name. It can be passed data to operate on (i.e. the parameters) and can skksoptionally return data (the return value). All data that is passed to a function is explicitly passed.

- A method is a piece of code that is called by name that is associated with an object. In most respects it is identical to a function except for two key differences:
  - A method is implicitly passed the object on which it was called.
  - A method is able to operate on data that is contained within the class (remembering that an object is an instance of a class - the class is the definition, the object is an instance of that data).

