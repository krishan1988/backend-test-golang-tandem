FROM mongo

ADD db.json /db.json
CMD mongoimport --host mongodb --db btg-db --collection facts --type json --file /db.json --jsonArray