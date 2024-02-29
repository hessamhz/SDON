from django.http import JsonResponse
from influxdb_client import InfluxDBClient
from django.conf import settings

def get_latest_data_json(request):
    # InfluxDB connection details

    INFLUX_TOKEN = settings.INFLUX_TOKEN
    INFLUX_URL = settings.INFLUX_URL
    INFLUX_ORG = settings.INFLUX_ORG
    INFLUX_BUCKET = settings.INFLUX_BUCKET

    try:
        client = InfluxDBClient(url=INFLUX_URL, token=INFLUX_TOKEN, org=INFLUX_ORG)
        query_api = client.query_api()

        query = f'''
        from(bucket: "{INFLUX_BUCKET}")
        |> range(start: -40s)
        |> sort(columns: ["_time"], desc: true)
        |> limit(n:1)
        '''

        result = query_api.query(query)

        structured_data = []
        for table in result:
            for record in table.records:
                # Create a structured representation of the data
                data_point = {
                    "measurement": record.get_measurement(),
                    "field": record.get_field(),
                    "value": record.get_value(),
                    "time": record.get_time().isoformat()
                }
                # Optionally, you could filter here based on specific measurements or fields
                structured_data.append(data_point)

        client.close()

        return JsonResponse({"data_points": structured_data})
    except Exception as e:
        # Log the error or handle it as preferred
        return JsonResponse({'error': str(e)}, status=500)