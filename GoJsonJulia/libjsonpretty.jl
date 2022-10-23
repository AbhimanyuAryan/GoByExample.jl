module libjsonpretty
# Now write our Julia idiomatic function
function jsonpretty(txt::AbstractString)
    value = ccall((:jsonpretty, "./libjsonpretty"), Cstring, (Cstring,), txt)
    convert(UTF8String, bytestring(value))
end

end