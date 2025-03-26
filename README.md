
POST /subscribe
//Input Json
  {
  "url": "http://example.com/webhook",
  "area_code": "123"
  }

//Output
{
  "message": "Successfully subscribed to notifications."
}


POST /unsubscribe
    //Input Json

    {
  "url": "http://example.com/webhook",
  "area_code": "123"
}

//Output Json

{
  "message": "Successfully unsubscribed from notifications."
}


//Consider 3 MODES

    - tdm
    - solo
    - multiplayer

POST mode/join

  {
  "area_code": "123",
  "mode": "solo"
}


//Or

{
  "area_code": "123",
  "mode": "tdm"
}



# Go
