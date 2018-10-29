# Goa bug
This is a simple, demo repo to reproduce a bug with `goa gen`

```bash
$ goa gen github.com/adam-hanna/goa-bug/design
exit status 2
panic: reflect: call of reflect.Value.Set on zero Value

goroutine 1 [running]:
reflect.flag.mustBeExported(0x0)
  /home/adam/.gvm/gos/go1.10.3/src/reflect/value.go:215 +0xae
reflect.Value.Set(0x849f80, 0xc4203953e0, 0x195, 0x0, 0x0, 0x0)
  /home/adam/.gvm/gos/go1.10.3/src/reflect/value.go:1368 +0x40
reflect.Append(0x81de80, 0xc4203953c0, 0x97, 0xc42028a640, 0x1, 0x1, 0x97, 0xc420330800, 0xc4203437a0)
  /home/adam/.gvm/gos/go1.10.3/src/reflect/value.go:1864 +0xe3
goa.design/goa/design.(*Array).MakeSlice(0xc420270310, 0xc42037d900, 0x4, 0x4, 0x4, 0x4)
  /home/adam/go/src/goa.design/goa/design/types.go:468 +0x125
goa.design/goa/design.(*Array).Example(0xc420270310, 0xc4203437a0, 0xc4202be6c0, 0x11)
  /home/adam/go/src/goa.design/goa/design/types.go:459 +0x142
goa.design/goa/design.(*UserTypeExpr).recExample(0xc4203307e0, 0xc4203437a0, 0x0)
  /home/adam/go/src/goa.design/goa/design/user_type.go:97 +0x125
goa.design/goa/design.(*UserTypeExpr).Example(0xc4203307e0, 0xc4203437a0, 0x0, 0x0)
  /home/adam/go/src/goa.design/goa/design/user_type.go:81 +0x35
goa.design/goa/design.(*AttributeExpr).Example(0xc420332320, 0xc4203437a0, 0xc42028aad8, 0x3)
  /home/adam/go/src/goa.design/goa/design/example.go:63 +0x1ca
goa.design/goa/http/codegen.buildResponseBodyType(0xc420378900, 0xc4202edf40, 0xc420332320, 0xc420361cc0, 0x0, 0xc42028aea0, 0xc420272910, 0xb, 0xc42039abe0)
  /home/adam/go/src/goa.design/goa/http/codegen/service_data.go:2055 +0x4cb
goa.design/goa/http/codegen.buildResponses(0xc4202edf40, 0xc420361cc0, 0x1, 0xc420378900, 0xc420272910, 0xb, 0x98, 0x98, 0xc420393040)
  /home/adam/go/src/goa.design/goa/http/codegen/service_data.go:1333 +0x3a7
goa.design/goa/http/codegen.buildResultData(0xc4202edf40, 0xc420378900, 0x11c)
  /home/adam/go/src/goa.design/goa/http/codegen/service_data.go:1255 +0x249
goa.design/goa/http/codegen.ServicesData.analyze(0xc4202f1170, 0xc420320000, 0x6)
  /home/adam/go/src/goa.design/goa/http/codegen/service_data.go:793 +0x2672
goa.design/goa/http/codegen.ServicesData.Get(0xc4202f1170, 0x8ca643, 0x6, 0xc4202bfd40)
  /home/adam/go/src/goa.design/goa/http/codegen/service_data.go:564 +0xa9
goa.design/goa/http/codegen.server(0xc4202d61d2, 0x21, 0xc420320000, 0xc4203679b0)
  /home/adam/go/src/goa.design/goa/http/codegen/server.go:30 +0xf8
goa.design/goa/http/codegen.ServerFiles(0xc4202d61d2, 0x21, 0xc04780, 0x0, 0x4, 0x4)
  /home/adam/go/src/goa.design/goa/http/codegen/server.go:19 +0xdc
goa.design/goa/codegen/generator.Transport(0xc4202d61d2, 0x21, 0xc4202779c0, 0x3, 0x4, 0x4, 0x4, 0x4, 0x0, 0x0)
  /home/adam/go/src/goa.design/goa/codegen/generator/transport.go:19 +0x94
goa.design/goa/codegen/generator.Generate(0x7fffff202606, 0x1, 0x8c2d04, 0x3, 0x0, 0x0, 0xc42026edc0, 0x0, 0xffffffffffffffff)
  /home/adam/go/src/goa.design/goa/codegen/generator/generate.go:56 +0x2fe
main.main()
  /home/adam/go/src/github.com/adam-hanna/goa-bug/goa704262151/main.go:54 +0x2a2
 ```
