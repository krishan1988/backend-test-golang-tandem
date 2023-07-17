

### Generate token

- Command
```sh
cd tools
chmod +x generate_token.sh
sh generate_token.sh

```
- Output

```
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTIyMDQyMjh9.TC2fGAQGa8U7sVZXL33y83fnvIEb6wk6c68rSE14B90m0Pa0m-xsKzNRAbZ5T1luN2BEC5v3169yWt0yPsSlO5OgY4EfbsnFiB1qq3KoR8JNXmgAd1TEtFobcrQXH7SSpkCLRT5_QrCedd2SJX6dsCNQiV_Zp90_CHPg71XdZlc
```

### Create a request

- Create below cURL request with proper host name and
token (generated token)
```sh

curl --location 'http://localhost:8083/api/v1/facts?limit=10&page=4&type=trivia&found=true' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTIyMDQyMjh9.TC2fGAQGa8U7sVZXL33y83fnvIEb6wk6c68rSE14B90m0Pa0m-xsKzNRAbZ5T1luN2BEC5v3169yWt0yPsSlO5OgY4EfbsnFiB1qq3KoR8JNXmgAd1TEtFobcrQXH7SSpkCLRT5_QrCedd2SJX6dsCNQiV_Zp90_CHPg71XdZlc'

```

- Response

```json
{
    "facts": [
        {
            "id": "64b2dfa78fee82cc268ca5d2",
            "text": "490 is the number of Pokémon available as of the release of Pokémon Diamond and Pearl (excluding event Pokémon).",
            "number": 490,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5d3",
            "text": "179 is the rank of the Anthology 1961-1977 (1992) by Curtis Mayfield and The Impressions on Rolling Stone magazine's list of the 500 greatest albums of all time.",
            "number": 179,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5d4",
            "text": "198 is the rank of Hey Joe (1966) by Jimi Hendrix on Rolling Stone magazine's list of The 500 Greatest Songs of All Time.",
            "number": 198,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5d5",
            "text": "5000 is the number of base pairs in the DNA of the simplest viruses.",
            "number": 5000,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5d6",
            "text": "328 is the weight in pounds of an ovarian cyst removed from a woman in Galveston, Texas, in 1905, a world record.",
            "number": 328,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5d7",
            "text": "23 is the number of times Julius Caesar was stabbed.",
            "number": 23,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5d8",
            "text": "31 is the number of flavors of Baskin-Robbins ice cream.",
            "number": 31,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5d9",
            "text": "1362310155 is the total number of items of mail that went through the Canadian postal system in 1950.",
            "number": 1362310155,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5da",
            "text": "352 is the number of international appearances by Kristine Lilly for the USA women's national soccer team, an all-time record.",
            "number": 352,
            "found": true,
            "type": "trivia"
        },
        {
            "id": "64b2dfa78fee82cc268ca5db",
            "text": "4 is the number of movements in a symphony.",
            "number": 4,
            "found": true,
            "type": "trivia"
        }
    ]
}
```