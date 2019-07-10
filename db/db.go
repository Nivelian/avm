package db

import (
  "database/sql"
  _ "github.com/denisenkom/go-mssqldb"
  . "github.com/Nivelian/avm/utils"
)

type Connection = *sql.DB
type Transaction = *sql.Tx
type Config struct {
	Host     string
	User     string
	Password string
	Db       string
}

func Start (config Config) (Connection, Transaction) { 
  conn, err := sql.Open("mssql", FmtNamed("sqlserver://{User}:{Password}@{Host}?database={Db}", config))
  PanicIf(err, "Failed to oped to open db connection")
  tx, err := conn.Begin()
  PanicIf(err, "Failed to begin transaction")
  return conn, tx
}

func Finish (conn Connection, tx Transaction, err error) {
  if err != nil {
    tx.Rollback()
    panic(err)
  } else {
    PanicIf(tx.Commit(), "Failed to commit changes")
  }
  conn.Close()
}

func Check (conn Connection, tx Transaction) func (error, string) {
  return func (err error, succ string) {
           if err != nil { Finish(conn, tx, err) } else { Log(succ) }}
}

func delete (tx Transaction, table string, instance int) error {
  _, err := tx.Exec( Fmt("delete from {0} where InstanceId = {1}", table, instance) )
  return err
}
func ClearAll (tx Transaction, instance int) (err error) {
  for _, s := range []string{"Schedule", "RouteTracks", "StopPoints", "Routes"} {
    err = delete(tx, Fmt("[AvmGeneral].[dbo].[trby.{0}]", s), instance)
    if (err != nil) {
      Error( Fmt("Failed to clear {0} table", s) )
      break;
    }
  }
  return
}

const (
  ROUTE_ID       = "[dbo].[avm.Encode.Id](:Instance, :RouteId)"
  STATION_ID     = "case when :StationId = 0 then null else [dbo].[avm.MarshutStopPoints.Encode.Id](:Instance, :StationId) end"
  POINT_ID       = "[dbo].[avm.Encode.RouteTracks.Id](:Instance, :PointId, 1, :PointId)"
  EXTRA_POINT_ID = "[dbo].[avm.Encode.RouteTracks.ReferenceId](:PointId, 1, :PointId)"
  SCHEDULE_ID    = "[dbo].[avm.Schedules.Encode.Id](:Instance, :ScheduleId)"
  GEO_POINT      = "[geography]::[Point](:Lat, :Lng, 4326)"
)

func prepare (table, tpl string) string {return Fmt("insert into [AvmGeneral].[dbo].[trby.{0}] values ({1})", table, tpl)}
var TEMPLATES = map[string]string {
  "route":    prepare("Routes",      Fmt("{0}, :Instance, :RouteId, :RouteId, :Type", ROUTE_ID)),
  "station":  prepare("StopPoints",  Fmt("{0}, :Instance, :StationId, :Type, :Direction, {1}, :TitleRu, :TitleBy, :TitleEn",
                                         STATION_ID, GEO_POINT)),
  "point":    prepare("RouteTracks", Fmt("{0}, {1}, :Instance, {2}, :PointId, {3}, :Direction, :Distance, :Lat, :Lng",
                                         POINT_ID, EXTRA_POINT_ID, STATION_ID, ROUTE_ID)),
  "schedule": prepare("Schedule",    Fmt("{0}, :ScheduleId, {1}, :Instance, {2}, :CardNum, :Time, :Weekdays, :FileId, :TripId",
                                         SCHEDULE_ID, ROUTE_ID, STATION_ID)),
}

func namedArgs (o map[string]interface{}) (res []interface{}) {
  for k, v := range o {
    res = append(res, sql.Named(k, v))
  }
  return
}

func insert (tx Transaction, key string, data []interface{}) (err error) {
  tpl, err := tx.Prepare(TEMPLATES[key])
  defer tpl.Close()
  for _, x := range data {
    _, err = tpl.Exec(namedArgs( StructToMap(x) )...)
    if (err != nil) {
      Error( Fmt("Failed to insert {0} data to table", key) )
      break;
    }
  }
  return
 }

func InsertRoutes   (tx Transaction, data []interface{}) (err error) {return insert(tx, "route",    data)}
func InsertStations (tx Transaction, data []interface{}) (err error) {return insert(tx, "station",  data)}
func InsertPoints   (tx Transaction, data []interface{}) (err error) {return insert(tx, "point",    data)}
func InsertSchedule (tx Transaction, data []interface{}) (err error) {return insert(tx, "schedule", data)}
