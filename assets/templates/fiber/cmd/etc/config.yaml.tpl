Name: {{.}}
Host: 0.0.0.0
Port: 8888

Logger:
  Name: {{.}}
  IsDev: true
  Level: Debug

captcha:
  KeyLong: 6
  ImgWidth: 240
  ImgHeight: 80

Mysql:
  DataSource: root:root@tcp(127.0.0.1:3306)/ly_admin?charset=utf8&parseTime=True&loc=Local

Redis:
  Addr: localhost:6379

JWT:
  SigningKey: "@fineByMe"
  ExpiresTime: 604800
  BufferTime: 86400

Cypher:
  AesKey: RicoLyRyanSunny。