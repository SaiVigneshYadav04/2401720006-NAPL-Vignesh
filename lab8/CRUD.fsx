#r "nuget: MySqlConnector"

open MySqlConnector

let connectionString =
    "Server=localhost;Port=3306;Database=college_db_napl;User ID=root;Password="

let connection = new MySqlConnection(connectionString)
connection.Open()

let execute sql =
    use command = new MySqlCommand(sql, connection)
    command.ExecuteNonQuery() |> ignore

let readStudents () =
    use command = new MySqlCommand("SELECT * FROM Students", connection)
    use reader = command.ExecuteReader()

    while reader.Read() do
        printfn "ID: %d, Name: %s, Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

printfn "CREATE:"
execute "INSERT INTO Students VALUES (1, 'Sid', 20)"
readStudents()

printfn "\nREAD:"
readStudents()

printfn "\nUPDATE:"
execute "UPDATE Students SET Age = 21 WHERE Id = 1"
readStudents()

printfn "\nDELETE:"
execute "DELETE FROM Students WHERE Id = 1"
readStudents()

printfn "CREATE:"
execute "INSERT INTO Students VALUES (2, 'Tom', 20)"
readStudents()

connection.Close()