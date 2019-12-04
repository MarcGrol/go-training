# Integrate with Twilio

    curl -X POST \
         --data-urlencode 'To=<ToNumber>' \
         --data-urlencode 'From=<FromNumber>' \
         --data-urlencode 'Body=<BodyText>' \
         -u <AccountSid>:<Password>
         'https://api.twilio.com/2010-04-01/Accounts/<AccountSid>/Messages.json'
    
 or:
             
    curl -X POST \
      -H 'Accept: application/x-www-form-urlencoded' \
      -H 'Accept: application/json' \
      --data 'To=+3156666&Body=my body' \
      -u <AccountSid>:<Password>
       'https://api.twilio.com/2010-04-01/Accounts/<AccountSid>/Messages.json'
                          
                 
                     
   Use https://mholt.github.io/curl-to-go/
                     