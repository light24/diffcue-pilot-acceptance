event = name [ "{" label "}" ] ":" value
label = name "=" name
name = letter { letter }
value = [ "-" ] digit { digit } [ duration_unit ]
duration_unit = "ms" | "s" | "m"
