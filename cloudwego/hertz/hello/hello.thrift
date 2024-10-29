namespace go hello.world

service HelloService {
    string SayHello(1:string name) (api.get="/hello")
    string Bye(1:string name) (api.get="/bye")
}