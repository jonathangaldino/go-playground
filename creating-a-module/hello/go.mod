module creating-a-module/hello

go 1.16

replace creating-a-module/greetings => ../greetings

require creating-a-module/greetings v0.0.0-00010101000000-000000000000
