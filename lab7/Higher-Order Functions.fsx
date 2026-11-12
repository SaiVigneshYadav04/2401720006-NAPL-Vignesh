let list1 = [1; 2; 3; 4; 5]
let list2 = [10; 15; 20; 25; 30]

printfn "Original List 1: %A" list1
printfn "Original List 2: %A" list2

let square1 = List.map (fun x -> x * x) list1
let square2 = List.map (fun x -> x * x) list2

printfn "\nList.map - Squares:"
printfn "List 1: %A" square1
printfn "List 2: %A" square2

let even1 = List.filter (fun x -> x % 2 = 0) list1
let even2 = List.filter (fun x -> x % 2 = 0) list2

printfn "\nList.filter - Even Numbers:"
printfn "List 1: %A" even1
printfn "List 2: %A" even2

let sum1 = List.fold (fun acc x -> acc + x) 0 list1
let sum2 = List.fold (fun acc x -> acc + x) 0 list2

printfn "\nList.fold - Sum:"
printfn "List 1: %d" sum1
printfn "List 2: %d" sum2
