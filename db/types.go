package db

type Route struct {
  Instance int
  RouteId  int
  Type     int
}

type Station struct {
  StationId int
  Instance  int
  Type      int
  TitleRu   string
  TitleEn   string
  TitleBy   string
  Direction int
  Lat       float64
  Lng       float64
}

type Point struct {
  PointId   int
  Instance  int
  StationId int
  RouteId   int
  Direction int
  Distance  int
  Lat       float64
  Lng       float64
}

type Schedule struct {
  ScheduleId int
  Instance   int
  RouteId    int
  StationId  int
  TripId     int
  Time       int
  Weekdays   int
  CardNum    int
  FileId     int
}

