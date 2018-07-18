# arabicToRoman
#อธิบายโจทย์
```
สร้าง API ชื่อ convertToRomanNumber 

URL : localhost:3000/convertToRomanNumber?number=

ส่ง paramiter ชื่อ number เป็นเลข arabic

response json {

    romanNumber

}
```
#โครงสร้าง
```
calculateromannumber เป็นงานที่ พี่หมี พี่เจ ทำ

golf-lek เป็นงานที่ พี่กอล์ฟ เล็ก ทำ
```
## HOW TO RUN PROGRAM
### Set Go Path
```
set GOPATH=%CD%
```
## RUN Unit Test
```
go run romannumber
```
## RUN Acceptance Test
```
newman run src/atdd/postman/ ชื่อไฟล์
```