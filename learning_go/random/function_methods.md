In Go, methods are functions that are associated with a particular type, and they can be called on instances of that type.
- Receiver Parameter: Function methods have a special parameter called a receiver, which is a reference to an instance of a specific type. This receiver parameter is placed before the function name in the method declaration. Normal functions do not have this receiver parameter.
- Associativity: Function methods are associated with a specific type. They are called on instances of that type and can access and modify the state of those instances. Normal functions are not associated with any specific type and are standalone entities.
- Method Call: To call a function method, you call it on an instance of the associated type. In contrast, you call a normal function by its name directly, without any receiver.

```
func (b Contents) Push(uploadRef regname.Tag, labels map[string]string, registry ImagesMetadataWriter, logger Logger) (string, error)
```
- When you call a method on a value of a user-defined type, it typically works on a copy of that value. This is a key behavior of methods in Go, known as pass-by-value semantics
- When you call a method on a value, Go makes a copy of that value and passes it to the method. This copy is what the method operates on. This behavior is in contrast to passing a pointer to the value.
 
If u want to work on directly :
```
func (b *Contents) Push(uploadRef regname.Tag, labels map[string]string, registry ImagesMetadataWriter, logger Logger) (string, error)
```