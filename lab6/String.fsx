let reverse (s: string) =
    s.ToCharArray() |> Array.rev |> System.String

let palindrome (s: string) =
    s = reverse s

printfn "Reverse: %s" (reverse "madam")
printfn "Palindrome: %b" (palindrome "madam")
printfn "Palindrome: %b" (palindrome "hello")

let transform (s: string) =
    if s <> "" && (s |> Seq.forall System.Char.IsLetter) then
        s.ToUpper()
    elif s <> "" && (s |> Seq.forall System.Char.IsDigit) then
        s + s
    else
        s

printfn "Result: %s" (transform "hello")
printfn "Result: %s" (transform "123")
printfn "Result: %s" (transform "abc123")

let emailCheck (s: string) =
    let parts = s.Split('@')
    parts.Length = 2 && parts.[0] <> "" && parts.[1].Contains(".")

printfn "Email: %b" (emailCheck "a@gmail.com")
printfn "Email: %b" (emailCheck "hello@yahoo.com")
printfn "Email: %b" (emailCheck "invalid")