let list1 = [1..6]  
                |> List.map(fun x -> x * x) 
                |> List.filter (fun x -> x % 2 = 0)
                |> List.fold (fun acc x -> acc + x) 0

printfn "%d" list1

let list2 = [1..100]  
                |> List.map(fun x -> x * x) 
                |> List.filter (fun x -> x % 2 = 0)
                |> List.fold (fun acc x -> acc + x) 0

printfn "%d" list2
