# Instructions et Résultats

**1. Affichage de l'adresse IP locale :**

Adresse IPv4 : ```10.33.77.199```

Masque de sous-réseau : ```255.255.240.0```

Passerelle par défaut : ```10.33.79.254```

**2. Informations sur le réseau :**

Adresse réseau : ```10.33.64.0```

Adresse de diffusion : ```10.33.79.255```

Masque de sous-réseau : ```255.255.240.0```

Adresse loopback : ```127.0.0.1```

**3. Passerelle du réseau Ynov :**

Passerelle par défaut : ```10.33.79.254```

**4. Adresse IP obtenue :**

IPv4 : ```10.33.77.154```

**5. Interrogation des adresses IP :**

Résultat pour google.com : ```2a00:1450:4007:805::200e, 172.217.18.206```

Résultat pour 8.8.8.8 : ```Erreur - domaine inexistant```

Résultat pour 1.1.1.1 : ```1.1.1.1```

**6. Adresses IP de la machine et adresse loopback :**

*Interface Loopback (lo) :*

IPv4 : ```127.0.0.1```

IPv6 : ```::1```

*Interface Ethernet (enp0s3) :*

IPv4 : ```10.0.2.15/24```

Adresse de diffusion : ```10.0.2.255```

IPv6 : ```fe80::a00:27ff:fe1d:850d/64```

**7. Résultat de la commande cat /etc/hostname :**

tpr-client1 : ```tpr-client1```

tpr-client2 : ```tpr-client2```

**8. Ping entre VM (tpr-client1 vers tpr-client2) :**

Paquets transmis : 150

Paquets reçus : 150

Perte de paquets : 0%

RTT (min/avg/max/mdev) : ```0.043/0.087/0.308/0.046 ms```

**9. Ping des VM depuis l'hôte :**

Adresse IP : ```10.0.2.15```

Paquets transmis : 11

Paquets reçus : 11

Perte de paquets : 0%

RTT (min/avg/max/mdev) : ```0.096/0.126/0.162/0.023 ms```

**10. Utilisation de nslookup et ping de www.ynov.com :**

Paquets transmis : 7

Paquets reçus : 7

Perte de paquets : 0%

RTT (min/avg/max/mdev) : ```16.523/20.606/29.618/4.530 ms```