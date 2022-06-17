# Les tests unitaires en Go

Go intègre nativement une librairie de tests unitaires. Pour s'en servir,
rien de plus simple, il suffit de créer un fichier terminant par "_test.go".

À partir de ce moment, Go comprendra que c'est un fichier de test et il sera
exécuté avec la commande

```shell
cd ./maths # on doit l'exécuter depuis le bon dossier
go test -v # Pour activer le mode verbose et avoir tous les logs
```

Dans ce fichier de test, on va simplement créer des fonctions qui prennent comme
argument un élément `*testing.T`

Par exemple `func TestAdd(t *testing.T) {...}`