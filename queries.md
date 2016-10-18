# Useful queries for debugging

## User settings

```
PASSWORD=abcd
USER=USER_00000000000000000000000000000000009
SITE=SITE_00000000000000000000000000000000003
DEVICE=DEVC_00000000000000000000000000000000264
ROOM=ROOM_00000000000000000000000000000000005
SCENARIO=SCNR_00000000000000000000000000000002251
SERVER=http://10.0.1.6:8080
SESSION=$(curl -k "$SERVER/DomoBox/rs/Mobile/Login?site_key=$SITE&user_key=$USER&password=$PASSWORD")
```

## xml responses

```
brew install xmlstarlet

curl -k "$SERVER/DomoBox/rs\?_wadl" | xml format

curl -k "$SERVER/DomoBox/rs/Mobile/GetCategoriesTranslations" | xml format

curl -k "$SERVER/DomoBox/Mobile/GetSites" | xml format

curl -k "$SERVER/DomoBox/Mobile/GetUsers?site_key=$SITE" | xml format

curl -k "$SERVER/DomoBox/Mobile/LoginInfos?site_key=$SITE&user_key=$USER&password=$PASSWORD"
```

## json responses

```
brew install jq
curl -k "$SERVER/DomoBox/rs/Mobile/GetSites" | jq

curl -k "$SERVER/DomoBox/rs/Mobile/GetUsers?site_key=$SITE" | jq

curl -k "$SERVER/DomoBox/rs/Mobile/LoginInfos?site_key=$SITE&user_key=$USER&password=$PASSWORD" | jq

curl -k "$SERVER/DomoBox/rs/Mobile/GetCategories?session_key=$SESSION&room_key=$ROOM" | jq

curl -k "$SERVER/DomoBox/rs/Mobile/ExecuteAction?session_key=$SESSION&target_key=$DEVICE&prop_clsid=CLSID-DEVC-PROP-TOR-SW&action_clsid=CLSID-ACTION-ON"

curl -k "$SERVER/DomoBox/rs/Mobile/ExecuteAction?session_key=$SESSION&target_key=$SCENARIO&prop_clsid=0&prop_numr=0&action_clsid=CLSID-ACTION-RUN"

curl -k "$SERVER/DomoBox/rs/Mobile/GetDeviceState?session_key=$SESSION&device_key=$DEVICE" | jq

curl -k "$SERVER/DomoBox/rs/Mobile/GetGroups?session_key=$SESSION" | jq

```

