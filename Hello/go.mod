module example/hello

go 1.22.4

require example/greetings v0.0.0-00010101000000-000000000000

require golang.org/x/example/hello v0.0.0-20240205180059-32022caedd6a // indirect

replace example/greetings => ../greetings
