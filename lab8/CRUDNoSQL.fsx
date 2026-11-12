
#r "nuget: MongoDB.Driver"

open MongoDB.Bson
open MongoDB.Driver

let client = MongoClient("mongodb://localhost:27017")
let database = client.GetDatabase("CollegeDB")
let collection = database.GetCollection<BsonDocument>("Students")

collection.DeleteMany(Builders<BsonDocument>.Filter.Empty) |> ignore

let student = BsonDocument()
student.Add("name", "Rahul")
student.Add("age", 20)

printfn "===== MongoDB (NoSQL Database) ====="

printfn "\nCREATE Operation:"
collection.InsertOne(student)

let allStudents =
    collection.Find(Builders<BsonDocument>.Filter.Empty).ToList()

printfn "Student Records:"
for s in allStudents do
    printfn "  %O" s

printfn "\nREAD Operation:"
let readStudents =
    collection.Find(Builders<BsonDocument>.Filter.Empty).ToList()

printfn "Student Records:"
for s in readStudents do
    printfn "  %O" s

printfn "\nUPDATE Operation:"
let filter = Builders<BsonDocument>.Filter.Eq("name", "Rahul")
let update = Builders<BsonDocument>.Update
                .Set("name", "Aman")
                .Set("age", 21)

collection.UpdateOne(filter, update) |> ignore

let updatedStudents =
    collection.Find(Builders<BsonDocument>.Filter.Empty).ToList()

printfn "Student Records:"
for s in updatedStudents do
    printfn "  %O" s

printfn "\nDELETE Operation:"
collection.DeleteOne(filter) |> ignore

let remainingStudents =
    collection.Find(Builders<BsonDocument>.Filter.Empty).ToList()

printfn "Student Records:"
for s in remainingStudents do
    printfn "  %O" s
