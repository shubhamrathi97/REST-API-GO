1. ADD FRIEND API 
curl -X POST \
  http://localhost:9090/addfriend \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: ae649983-1acc-423a-8f40-6ccdf98e4334' \
  -H 'cache-control: no-cache' \
  -d '{"UserID": "2", "FriendUserID":"1"}
'

2. UPDATE USER API
curl -X PUT \
  http://localhost:9090/updateuser \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: a9b53a68-254d-44a5-accf-588666327eaa' \
  -H 'cache-control: no-cache' \
  -d '{"ID": "2", "Name": "shubham3", "DOB":"10-03-1997"}
'

3. CREATE USER API
curl -X POST \
  http://localhost:9090/createuser \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 63d11967-0a83-4bfa-82c7-cfd47a307de6' \
  -H 'cache-control: no-cache' \
  -d '{"ID": "3", "Name": "shubham3", "DOB":"10-03-1997"}
'

4. GET USERS API 
curl http://localhost:9090/users

5. GET USER API 
http://localhost:9090/user?id=3