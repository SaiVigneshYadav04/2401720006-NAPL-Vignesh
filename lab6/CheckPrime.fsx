let isPrime n =
    if n < 2 then false
    else
        [2 .. n-1]
        |> List.filter(fun i -> n % i = 0)
        |> List.isEmpty
        

printfn "%b" (isPrime 7)
printfn "%b" (isPrime 5)
printfn "%b" (isPrime 10)
printfn "%b" (isPrime 13)

