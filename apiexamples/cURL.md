**GET Request:**
```
curl http://bibletag.xyz:8080/tag/Gospel
```
**GET Response:**
```
[
  {
    "book_id": "John",
    "book_name": "John",
    "book_order": "58",
    "chapter_id": "3",
    "chapter_title": "Chapter 3",
    "paragraph_number": "68",
    "verse_id": "16",
    "verse_text": "\u201cFor God so loved the world, that he gave his only Son, that whoever believes in him should not perish but have eternal life."
  },
  {
    "book_id": "John",
    "book_name": "John",
    "book_order": "58",
    "chapter_id": "3",
    "chapter_title": "Chapter 3",
    "paragraph_number": "68",
    "verse_id": "17",
    "verse_text": "For God did not send his Son into the world to condemn the world, but in order that the world might be saved through him."
  }
]
```

**POST Request:**
```
curl -X POST \
     -H "Content-Type: application/json" \
     -H "Cache-Control: no-cache" \
     -H "Postman-Token: aced34fd-2249-bff2-af6a-905ad2eeddfa" \
     -d '{
         "tag": "worry",
         "book": "1 peter",
         "chapter": 5,
         "startVerse": 7,
         "endVerse": 7
       }' "http://bibletag.xyz:8080/tag"
```
**POST Response:**
```
{
  "code": 200,
  "text": "Tagged Passage"
}
```