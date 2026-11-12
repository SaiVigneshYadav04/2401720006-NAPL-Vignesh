
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

    printfn "Student Records:"
    while reader.Read() do
        printfn "  ID: %d, Name: %s, Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

execute "CREATE TABLE IF NOT EXISTS Students (Id INT PRIMARY KEY, Name VARCHAR(50), Age INT)"

execute "DELETE FROM Students"

printfn "===== MySQL (Relational Database) ====="

printfn "\nCREATE Operation:"
execute "INSERT INTO Students VALUES (1, 'Rahul', 20)"
readStudents()

printfn "\nREAD Operation:"
readStudents()

printfn "\nUPDATE Operation:"
execute "UPDATE Students SET Name = 'Aman', Age = 21 WHERE Id = 1"
readStudents()

printfn "\nDELETE Operation:"
execute "DELETE FROM Students WHERE Id = 1"
readStudents()

connection.Close()
