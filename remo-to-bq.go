package function

import (
    "context"
	"fmt"
    "os"
    "time"

    "cloud.google.com/go/bigquery"
    "github.com/tenntenn/natureremo"
)

type Data struct {
    Name         string    `bigquery:"NAME"`
    Datetime     time.Time `bigquery:"DATETIME"`
    Temperature  float64   `bigquery:"TEMPERATURE"`
    Humidity     float64   `bigquery:"HUMIDITY"`
    Illumination float64   `bigquery:"ILLUMINATION"`
}

func insertToBigQuery (ctx context.Context, items []Data) error {
    projectID := os.Getenv("PROJECT_ID")
    dataset := os.Getenv("DATASET")
    table := os.Getenv("TABLE")

    client, err := bigquery.NewClient(ctx, projectID)
    if err != nil {
        return err
    }
    defer client.Close()

    u := client.Dataset(dataset).Table(table).Uploader()

    if err = u.Put(ctx, items); err != nil {
        return err
    }

    return nil
}

func RemoToBq(ctx context.Context) error {
    cli := natureremo.NewClient(os.Getenv("REMO_ACCESS_TOKEN"))

    ds, err := cli.DeviceService.GetAll(ctx)
    if err != nil {
        fmt.Println(err)
        return err
    }

    items := []Data{}

    for _, d := range ds {
        data := Data{}
        data.Name = d.Name
        data.Datetime = time.Now()
        data.Temperature = d.NewestEvents[natureremo.SensorTypeTemperature].Value
        data.Humidity = d.NewestEvents[natureremo.SensorTypeHumidity].Value
        data.Illumination = d.NewestEvents[natureremo.SensortypeIllumination].Value

        fmt.Println(data)

        items = append(items, data)
    }

    if err := insertToBigQuery(ctx, items); err != nil {
		fmt.Println(err)
        return err
    }

   return nil
}
