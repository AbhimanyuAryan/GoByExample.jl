include("./libjsonpretty.jl")
import libjsonpretty

src = """{"name":"fred","age":25,"height":75,"units":"inch","weight":"239"}"""
println("Our origin JSON src", src)
value = libjsonpretty.jsonpretty(src)
println("And out pretty version\n", value)