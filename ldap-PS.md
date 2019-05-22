## Comandos Utiles para interactuar con LDAP desde Powershell 

# Consultar CN
(New-Object adsisearcher((New-Object adsi("LDAP://dc.fulcrum.local","fulcrum\ldap","PasswordForSearching123!")),"(objectCategory=Computer)")).FindAll() | %{ $_.Properties.name }

# Comando Util
(New-Object adsisearcher((New-Object adsi("LDAP://dc.fulcrum.local","fulcrum\ldap","PasswordForSearching123!")),"(info=*)")).FindAll() | %{ $_.Properties }
