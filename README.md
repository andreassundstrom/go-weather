# Go weather

This cli calls the SMHI weather API:

```ps
go-weather forecast

========================================================
Date            MaxTemp Precipitation (mm/h)    Symbol
========================================================
2024-01-27      3°C     0                       ❄️☔
2024-01-28      4°C     0                       ☁️
2024-01-29      4°C     0                       🌤️
2024-01-30      3°C     0                       ☁️
2024-01-31      3°C     0                       ☔☔
2024-02-01      6°C     0                       ☁️
2024-02-02      4°C     0                       ☀️
2024-02-03      1°C     0                       🌤️
2024-02-04      -1°C    0                       ⛅
2024-02-05      -1°C    0                       🌤️
```
## Build

```ps
go build .\go-weather\
```

At the moment it is hard coded to return current weather in Stockholm.
