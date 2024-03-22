# Auto Reload Project Data

Create a microservice to update a json file every 15 seconds with random numbers between 1-100 for valuewater and wind. As follows:

```json
{
    "status": {
        "wind":78,
        "water":55
    }
}
```

Then display the data in the path `/` . Apart from that, you have to determine the water and wind status. Under the condition:
- if the water is below 5 then the status is safe
- if the water is between 6 - 8 then the status is alert
- if the water is above 8 then the status is dangerous
- if the wind is below 6 then the status is safe
- if the wind is between 7 - 15 then the alert status
- if the wind is above 15 then the status is dangerous
- water value in meters; wind value in meters per second