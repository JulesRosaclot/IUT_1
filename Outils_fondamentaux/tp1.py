import numpy as np
import matplotlib.pyplot as plt


"""print(np.array([3,7,-1,2]))
#ici la fonction créée un tableau.
print(np.ones(4))
#créée un tableau de 1 suivi de points, avec en indice la longeur de tableau q'on veut.
print(np.array([3,7,-1,2])+np.ones(4))
#ajoute le tableau de 1 au tableau de base.
print(np.array([3,7,-1,2])+1)
#ajoute 1 à chaque valeurs du tableau.
print(np.array([[3,7],[-1,2]]))
#créée un tableau de tableau.
print(np.array([1,2,3,4])**2)
#met au carré chaque valeurs du tableau.
print(np.arange(10,30,5))
#créée un tableau avec les valeurs allant de 10 à 30 exclus avec un step de 5.
print(np.linspace(0,2,9)) 
#ici le dernier chiffre indique le nombre de chiffre intermédiaire avec toujours le même écart qu'on veut voir jusqu'au deuxieme.
print(np.sin(np.linspace(0,2*np.pi,20)))
#ici même chose que celui au dessus mais en plus de ça la fonction calcule le sinus de chaque valeurs.
x=np.array([1,2,5,7,10,13])
y=np.array([3,2,0,1,-4,6])
#ici on donne les valeurs des coordonnées x,y de point. 
print(plt.plot(x,y, color='g', linestyle='dashed', marker='o'))
print(plt.show())
#et la on modélise les points sur un graphique tous relier par une courbe."""
x = np.linspace(0, 3*np.pi,100)
f = x**2*np.sin(x)+4
g = 30/(x**2+1)

plt.plot(x,f)
plt.plot(x,g)
plt.show()
#représentation de la courbe de la formule 
plt.clf()
plt.scatter(x,f)
plt.show()
#représentation de la fonction avec les points