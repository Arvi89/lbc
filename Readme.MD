Test Le Bon Coin

Pour tester, cloner le repo puis lancer :<br />
`docker-compose up`<br />
Il se peut que l'API soit lancée avant la DB, juste relancer pour s'assurer que tous les services soient up

Aller ensuite sur l'adminer: http://localhost:8080 <br />
Se connecter avec les logins lbc/lbc et importer le fichier db.sql du repo

Ensuite on peut commencer à utiliser l'API.

Pour créer une annonce :<br />
`POST http://localhost:4000/ads` <br />
avec comme body en json :

````
{
    "title": "test title",
    "content": "test content",
    "car_model": "ds3"
}
````
La réponse sera sous cette forme :
````
{
    "id": 1,
    "title": "test title",
    "content": "test content",
    "category": "auto",
    "car": {
        "brand": "Citroen",
        "model": "Ds3"
    }
}
````
Pour récupérer une annonce :<br />
`POST http://localhost:4000/ads/{ID}` <br />
La réponse sera de la même forme qu'au dessus, sauf si l'id n'est pas bon, dans ce cas ce sera soit une 400 soit une 404

Pour éditer une annonce :<br />
`PUT http://localhost:4000/ads/{ID}` <br />
Il faut envoyer en body la même chose que pour un post (même si c'est pour ré-écrire la même chose, car c'est un PUT)

Et pour supprimer une annonce :<br />
`DELETE http://localhost:4000/ads/{ID}` <br />
Une 204 sera renvoyée en cas de succès, sinon une 400 ou 404.

Notes pour amélioration : les requêtes peuvent sembler moyennement performantes, mais avec la structure des repository, il est facile d'ajouter une couche de cache et ainsi non seulement améliorer les performances, mais aussi moins surcharger la db.

J'ajouterais aussi le verbe PATCH pour un edit plus facile, une meilleure gestion des erreurs par endroits et des tests plus poussés pour chercher à cerner plus de corner cases.