# Integrate with Twilio

    curl -X POST \
         --data-urlencode 'To=+31648928856' \
         --data-urlencode 'From=+12016043948' \
         --data-urlencode 'Body=Dit su een test' \
         -u AC7d402060a59cc4ec6fd49f7f3f107cd3:e4ab04fde182fec5c4245040e5699bc4 \
         'https://api.twilio.com/2010-04-01/Accounts/AC7d402060a59cc4ec6fd49f7f3f107cd3/Messages.json'
    
 or:
             
    curl -X POST \
      -H 'Accept: application/x-www-form-urlencoded' \
      -H 'Accept: application/json' \
      --data 'To=+3156666&Body=my body' \
      -u <AccountSid>:<Password>
       'https://api.twilio.com/2010-04-01/Accounts/<AccountSid>/Messages.json'
                          
                 
                     
   Use https://mholt.github.io/curl-to-go/
                     