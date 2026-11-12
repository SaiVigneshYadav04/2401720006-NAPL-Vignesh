open System

let workerTask taskid waitTime = async {
    printfn "Starting task %d" taskid
    do! Async.Sleep(2000 * waitTime)
    printfn "Completed task %d" taskid
    return taskid
}

let works =[workerTask 1 1;workerTask 2 1;workerTask 3 1]
        |> Async.Parallel

Async.RunSynchronously works