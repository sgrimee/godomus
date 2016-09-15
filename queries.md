# Useful queries for debugging

Replace with your server's internal IP

## text responses

```
SESSION=$(curl http://10.0.1.6:8080/DomoBox/rs/Mobile/Login?site_key=SITE_00000000000000000000000000000000003&user_key=USER_00000000000000000000000000000000005&password=abcdef)
echo $SESSION
```

## xml responses

```
brew install xmlstarlet
curl http://10.0.1.6:8080/DomoBox/rs\?_wadl | xml format
curl http://10.0.1.6:8080/DomoBox/rs/Mobile/GetCategoriesTranslations | xml format
curl http://10.0.1.6:8080/DomoBox/Mobile/GetSites | xml format
curl http://10.0.1.6:8080/DomoBox/Mobile/GetUsers?site_key=SITE_00000000000000000000000000000000003 | xml format
curl http://10.0.1.6:8080/DomoBox/Mobile/LoginInfos?site_key=SITE_00000000000000000000000000000000003&user_key=USER_00000000000000000000000000000000005&password=abcdef
```

## json responses

```
brew install jq
curl http://10.0.1.6:8080/DomoBox/rs/Mobile/GetSites | jq
curl http://10.0.1.6:8080/DomoBox/rs/Mobile/GetUsers?site_key=SITE_00000000000000000000000000000000003 | jq
curl http://10.0.1.6:8080/DomoBox/rs/Mobile/LoginInfos?site_key=SITE_00000000000000000000000000000000003&user_key=USER_00000000000000000000000000000000005&password=abcdef | jq
curl http://10.0.1.6:8080/DomoBox/rs/Mobile/ExecuteAction?session_key=$SESSION&target_key=DEVC_00000000000000000000000000000000264&prop_clsid=CLSID-DEVC-PROP-TOR-SW&action_clsid=CLSID-ACTION-ON

```
also: GetDevices, GetCategories, Get
