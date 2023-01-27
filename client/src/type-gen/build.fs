open Fake.Core
open Fake.IO  
open System.Text.RegularExpressions
let abs2relative from  = 
    let regex = Regex(@"\\")  
    fun toPath ->
        regex.Replace(
            System.IO.Path.GetRelativePath(from, toPath)
            , "/"
        )

let initTargets () =  
    Target.create "Default" (fun _ -> 
        printfn "OpenApi型定義を生成"   
        let toRelative = abs2relative __SOURCE_DIRECTORY__ //System.Environment.CurrentDirectory
        let args = 
            let basePath = Path.combine __SOURCE_DIRECTORY__ @"./../../.." 
            let proto = toRelative @$"{basePath}/openapi/openapi.json"  
            let auto  = toRelative @$"{basePath}/client/src/core/openapi/openapi.ts"  
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