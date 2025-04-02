package main

import (
  "time"
  "log"
  _ "bufio"
  "net/http"
  "regexp"
  "strings"
  "encoding/csv"
  "github.com/gomodule/redigo/redis"
  "flag"
  _ "fmt"
)

//const DEFAULT_URL="https://standards-oui.ieee.org/oui/oui.txt"

// Standard 24-bit OUI
const DEFAULT_URL="http://standards-oui.ieee.org/oui/oui.csv"
// MA-M, 28-bit OUI
const DEFAULT_URL28="http://standards-oui.ieee.org/oui28/mam.csv"
// MA-S, 36-bit OUI
const DEFAULT_URL36="http://standards-oui.ieee.org/oui36/oui36.csv"

const DEFAULT_REDIS_SOCKET="/tmp/redis.sock"

func main() {

_ = DEFAULT_URL28
_ = DEFAULT_URL36

  oui24 := regexp.MustCompile(`^([0-9a-fA-F]{6})$`)
  oui28 := regexp.MustCompile(`^([0-9a-fA-F]{7})$`)
  oui36 := regexp.MustCompile(`^([0-9a-fA-F]{9})$`)

  var opt_u string
  var opt_s string
  var opt_M string
  var opt_S string

  var err error

  flag.StringVar(&opt_s, "s", DEFAULT_REDIS_SOCKET, "redis unix socket")
  flag.StringVar(&opt_u, "u", DEFAULT_URL, "URL of 24-bit oui.csv")
  flag.StringVar(&opt_M, "M", DEFAULT_URL28, "URL of 28-bit mem.csv")
  flag.StringVar(&opt_S, "S", DEFAULT_URL36, "URL of 36-bit oui36.csv")
  flag.Parse()

  // Fetch 24-bit OUI
  client := http.Client{Timeout: time.Second*60}

  var response *http.Response
  response, err = client.Get(opt_u)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  data := make(map[string]string)

  reader := csv.NewReader(response.Body)
  var records [][]string

  if records, err = reader.ReadAll(); err != nil {
    log.Fatal(err)
  }

  data["time"] = time.Now().String()

  for _, record := range records {
    if len(record) == 4 && oui24.MatchString(record[1]) {
      data[ strings.ToLower(record[1]) ] = strings.TrimSpace(record[2])
    }
  }

  // Fetch 28-bit OUI
  client28 := http.Client{Timeout: time.Second*60}

  var response28 *http.Response
  response28, err = client28.Get(opt_M)
  if err != nil {
    log.Fatal(err)
  }
  defer response28.Body.Close()

  data28 := make(map[string]string)

  reader28 := csv.NewReader(response28.Body)
  var records28 [][]string

  if records28, err = reader28.ReadAll(); err != nil {
    log.Fatal(err)
  }

  data28["time"] = time.Now().String()

  for _, record := range records28 {
    if len(record) == 4 && oui28.MatchString(record[1]) {
      data28[ strings.ToLower(record[1]) ] = strings.TrimSpace(record[2])
    }
  }

  // Fetch 36-bit OUI
  client36 := http.Client{Timeout: time.Second*60}

  var response36 *http.Response
  response36, err = client36.Get(opt_S)
  if err != nil {
    log.Fatal(err)
  }
  defer response36.Body.Close()

  data36 := make(map[string]string)

  reader36 := csv.NewReader(response36.Body)
  var records36 [][]string

  if records36, err = reader36.ReadAll(); err != nil {
    log.Fatal(err)
  }

  data36["time"] = time.Now().String()

  for _, record := range records36 {
    if len(record) == 4 && oui36.MatchString(record[1]) {
      data36[ strings.ToLower(record[1]) ] = strings.TrimSpace(record[2])
    }
  }




  // save to Redis
  red, err := redis.Dial("unix", opt_s, redis.DialConnectTimeout(time.Second*10),
                         redis.DialReadTimeout(time.Second*10),
                         redis.DialWriteTimeout(time.Second*10),
  )

  if err != nil {
    log.Fatal(err)
  }

  defer red.Close()

  for oui, corp := range data {
    _, err = red.Do("HSET", "oui", oui, corp)
    if err != nil {
      log.Fatal(err)
    }
  }

  for oui, corp := range data28 {
    _, err = red.Do("HSET", "oui28", oui, corp)
    if err != nil {
      log.Fatal(err)
    }
  }

  for oui, corp := range data36 {
    _, err = red.Do("HSET", "oui36", oui, corp)
    if err != nil {
      log.Fatal(err)
    }
  }
}
