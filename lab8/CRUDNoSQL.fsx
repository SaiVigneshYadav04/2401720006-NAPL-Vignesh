#r "nuget: MongoDB.Driver"

open MongoDB.Bson
open MongoDB.Driver

let client = MongoClient("mongodb://localhost:27017")
let database = client.GetDatabase("CollegeDB")
let collection = database.GetCollection<BsonDocument>("Students")


let student1 = BsonDocument()
student1.Add("Id", 1)
student1.Add("Name", "Sidharth")
student1.Add("Age", 20)

let student2 = BsonDocument()
student2.Add("Id", 2)
student2.Add("Name", "TOM")
student2.Add("Age", 21)

let student3 = BsonDocument()
student3.Add("Id", 3)
student3.Add("Name", "Elizabeth")
student3.Add("Age", 22)

printfn "CREATE:"
collection.InsertMany([student1; student2; student3])


printfn "\nREAD:"
let students =
    collection.Find(Builders<BsonDocument>.Filter.Empty).ToList()

for s in students do
    printfn "%O" s


printfn "\nUPDATE:"
let filter = Builders<BsonDocument>.Filter.Eq("Id", 1)
let update = Builders<BsonDocument>.Update.Set("Age", 21)

collection.UpdateOne(filter, update) |> ignore

let updatedStudent = collection.Find(filter).FirstOrDefault()
printfn "%O" updatedStudent


printfn "\nDELETE:"
collection.DeleteOne(filter) |> ignore

let remainingStudents =
    collection.Find(Builders<BsonDocument>.Filter.Empty).ToList()

for s in remainingStudents do
    printfn "%O" s

printfn "\nCRUD operations completed successfully!"