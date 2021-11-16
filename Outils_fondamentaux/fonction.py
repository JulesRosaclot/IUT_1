import numpy as np
import matplotlib.pyplot as plt
from numpy.core.function_base import linspace
a=1
b=0
c=-4
def resoudre(a,b,c) :
    D = b**2-4*a*c
    if a == 0 : 
        return []

    elif D > 0 :
        x1=(-b-np.sqrt(D))/(2*a)
        x2=(-b+np.sqrt(D))/(2*a)
        f1=a*x1**2+b*x1+c
        f2=a*x2**2+b*x2+c
        return [x1,x2]
        
    elif D == 0 :
        x=-b/(2*a)
        f=a*x**2+b*x+c
        return [x]
        
    else :
        return []

def parabole(a,b,c) :
    x=linspace((-b/2*a)-5,(-b/2*a)+5,100)
    y=a*x**2+b*x+c
    xsom=-b/2*4
    ysom=a*xsom**2+b*xsom+c
    plt.scatter(xsom, ysom)
    plt.plot(x, y)
    for i in resoudre(a,b,c) :
        plt.scatter(i, 0)
    plt.show()
parabole(a,b,c)
#ax.annotate("texte",xy=(np.pi,np.cos(np.pi)),xytext=(np.pi,np.cos(np.pi)+0.1))