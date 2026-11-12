
open System
open System.Text.RegularExpressions

let reverseString (s: string) =
    new string(Array.rev (s.ToCharArray()))

let isPalindrome (s: string) =
    s = reverseString s

let transformString (s: string) =
    if Regex.IsMatch(s, @"^[a-zA-Z]+$") then
        s.ToUpper()
    elif Regex.IsMatch(s, @"^[0-9]+$") then
        s + s
    else
        s

let isValidEmail (email: string) =
    Regex.IsMatch(email, @"^[^@\s]+@[^@\s]+\.[^@\s]+$")

printfn "Reverse and Palindrome:"
["madam"; "hello"; "level"]
|> List.iter (fun s ->
    printfn "%s -> %s, Palindrome: %b" s (reverseString s) (isPalindrome s))

printfn "\nString Transformation:"
["hello"; "123"; "abc123"]
|> List.iter (fun s ->
    printfn "%s -> %s" s (transformString s))

printfn "\nEmail Validation:"
["user@gmail.com"; "abc@"; "hello.world@yahoo.com"]
|> List.iter (fun email ->
    printfn "%s -> %b" email (isValidEmail email))