# nugget
Forensic Domain Specific Language

# TODO

- implement a grammar generator which will populate the .g4 file based off some JSON files, allowing for easy addition of new tools

- possible json:

```
    function fields: 
        name
        typesItOperatesOn
        actions

    action:
        name
        inputType
        returnType
        execution code???
        
    for example, sha1
        "name": "sha1", 
        
                

```