open System
open Microsoft.Data.Sqlite

type Student = {
    Id: int
    Name: string
    Marks: int
}

let connection = new SqliteConnection("Data Source=students.db")

[<EntryPoint>]
let main argv =
    connection.Open()

    use cmd = new SqliteCommand(
        "CREATE TABLE IF NOT EXISTS Students (Id INTEGER PRIMARY KEY, Name TEXT, Marks INTEGER)",
        connection
    )

    cmd.ExecuteNonQuery() |> ignore

    printfn "Student Management System"
    printfn "Database connected successfully!"

    connection.Close()
    0