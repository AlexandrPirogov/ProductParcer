# ProductParser

# Как использовать

```
docker build . --tag gopars:v1
```

```
docker run --rm -it gopars:v1 gopars -f files/<yourfile.ext>
```

## P.S
Размер imag'a Dockerfile будет до 300MB, хотя имел возможность сохранить image размером 12MB (вместе с файлами-примерами). Оставил бОльшую версию исходя из того, что
компилятор Go на хосте не установлен
