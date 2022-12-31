open Fake.Core
open Fake.IO  

let relativePathFromExe  = 
    let regex = System.Text.RegularExpressions.Regex(@"\\")  
    fun absPath -> 
        System.IO.Path.GetRelativePath(
            System.Environment.CurrentDirectory, absPath
        )  
        |> fun x -> regex.Replace(x, "/")

let initTargets () =  
    Target.create "Default" (fun _ -> 
        printfn "OpenApi型定義を生成"   
        let args = 
            let proto = relativePathFromExe <| Path.combine __SOURCE_DIRECTORY__ @"./../../openapi/openapi.json"  
            let auto  = relativePathFromExe <| Path.combine __SOURCE_DIRECTORY__ @"./../core/openapi/openapi.ts"  
            $@"openapi-typescript {proto} " + $@"--output {auto} " 
        // https://fake.build/reference/fake-javascript-npm.html#install 
        // https://github.com/fsprojects/FAKE/blob/master/src/app/Fake.JavaScript.Npm/Npm.fs#L28-28 
        ProcessUtils.tryFindFileOnPath "npx"
        |> function 
          | Some npx when System.IO.File.Exists npx -> Shell.Exec(npx, args) |> ignore 
          | _ -> printfn $"npx {args}" 
        |> ignore 
    )
 
[<EntryPoint>]
let main args =
    args
    |> Array.toList
    |> Context.FakeExecutionContext.Create false "build.fsx"
    |> Context.RuntimeContext.Fake
    |> Context.setExecutionContext 
    initTargets ()  
    Target.runOrDefault "Default"
    0