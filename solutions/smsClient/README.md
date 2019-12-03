# Integrate with Twilio

     curl -X POST 'https://api.twilio.com/2010-04-01/Accounts/<AccountSid>/Messages.json' \
             --data-urlencode 'To=<ToNumber>' \
             --data-urlencode 'From=<FromNumber>' \
             --data-urlencode 'Body=<BodyText>' \
             -u <AccountSid>:<Password>
             
             