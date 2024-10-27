Clone The Repo


I have used redis as NoSQL for data storage

Created APIS To Enter the Particular game mode , and to leave the game mode

I have created a webhook to get the real-time updates of the number of users in that mode based on the area code 

Once subscribed , Ig there is any change in that area , you will get Real-time updates 

Repository Link to simple Subscriber servern : https://github.com/The-Hawkeye/Game_Mode_Usage_Web_service-WebHook_Subscriber

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



